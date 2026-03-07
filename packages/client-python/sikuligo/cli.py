from __future__ import annotations

import argparse
import importlib.metadata
import os
import re
import shutil
import subprocess
import sys
import venv
from importlib import resources
from pathlib import Path
try:
    from .sikulix import _resolve_sikuli_binary
except ImportError:  # pragma: no cover - direct script execution fallback
    from sikulix import _resolve_sikuli_binary  # type: ignore


def _prompt_project_dir() -> Path:
    value = input("Project directory name: ").strip()
    if not value:
        raise SystemExit("Project directory name is required")
    return Path(value).expanduser().resolve()


def _resolve_project_dir(arg_dir: str | None) -> Path:
    if arg_dir:
        return Path(arg_dir).expanduser().resolve()
    return _prompt_project_dir()


def _copy_scaffold_examples(target_examples_dir: Path) -> None:
    try:
        source_root = resources.files("sikuligo").joinpath("scaffold/python/examples")
    except ModuleNotFoundError:
        source_root = Path(__file__).resolve().parent / "scaffold" / "python" / "examples"
    if not source_root.is_dir():
        raise SystemExit("packaged Python examples are missing from the sikuli-go distribution")

    if target_examples_dir.exists():
        shutil.rmtree(target_examples_dir)
    target_examples_dir.mkdir(parents=True, exist_ok=True)

    for item in source_root.iterdir():
        dst = target_examples_dir / item.name
        if item.is_dir():
            shutil.copytree(item, dst)
        else:
            shutil.copy2(item, dst)


def _write_requirements(project_dir: Path) -> Path:
    version = _package_version()
    requirements_path = project_dir / "requirements.txt"
    requirements_path.write_text(f"sikuli-go=={version}\n", encoding="utf-8")
    return requirements_path


def _package_version() -> str:
    try:
        return importlib.metadata.version("sikuli-go")
    except importlib.metadata.PackageNotFoundError:
        pyproject = Path(__file__).resolve().parents[1] / "pyproject.toml"
        if pyproject.exists():
            for line in pyproject.read_text(encoding="utf-8").splitlines():
                if line.strip().startswith("version ="):
                    return line.split("=", 1)[1].strip().strip('"')
        return "0.1.0"


def _venv_python(project_dir: Path) -> Path:
    if os.name == "nt":
        return project_dir / ".venv" / "Scripts" / "python.exe"
    return project_dir / ".venv" / "bin" / "python"


def _create_venv(project_dir: Path) -> Path:
    venv_dir = project_dir / ".venv"
    if not venv_dir.exists():
        builder = venv.EnvBuilder(with_pip=True)
        builder.create(venv_dir)
    python_bin = _venv_python(project_dir)
    if not python_bin.exists():
        raise SystemExit(f"venv python not found: {python_bin}")
    return python_bin


def _install_requirements(project_dir: Path, python_bin: Path, requirements_path: Path) -> None:
    out = subprocess.run(
        [str(python_bin), "-m", "pip", "install", "-r", str(requirements_path)],
        cwd=project_dir,
        check=False,
    )
    if out.returncode != 0:
        raise SystemExit("pip install -r requirements.txt failed")


def _run_init_py_examples(args: argparse.Namespace) -> int:
    project_dir = _resolve_project_dir(args.dir)
    project_dir.mkdir(parents=True, exist_ok=True)

    requirements_path = _write_requirements(project_dir)
    _copy_scaffold_examples(project_dir / "examples")
    python_bin = _create_venv(project_dir)
    if not args.skip_install:
        _install_requirements(project_dir, python_bin, requirements_path)

    print(f"Initialized SikuliGO Python project in: {project_dir}")
    print(f"Virtual environment: {project_dir / '.venv'}")
    print(f"Examples copied to: {project_dir / 'examples'}")
    return 0


def _detect_shell_profile() -> tuple[Path, str] | None:
    shell = os.environ.get("SHELL", "")
    home = Path.home()
    if "zsh" in shell:
        return (home / ".zshrc", "source ~/.zshrc")
    if "bash" in shell:
        return (home / ".bash_profile", "source ~/.bash_profile")
    return None


def _prompt_yes_no(question: str) -> bool:
    if not sys.stdin.isatty():
        return False
    answer = input(f"{question} [y/N]: ").strip().lower()
    return answer in ("y", "yes")


def _ensure_path_export(profile: Path, bin_dir: Path) -> bool:
    marker = "# Added by sikuli-go install-binary"
    export_line = f'export PATH="{bin_dir}:$PATH"'
    existing = profile.read_text(encoding="utf-8") if profile.exists() else ""
    if export_line in existing:
        return False
    snippet = f"{marker}\n{export_line}\n"
    prefix = "\n" if existing and not existing.endswith("\n") else ""
    profile.write_text(f"{existing}{prefix}{snippet}", encoding="utf-8")
    return True


def _runtime_canonical_name(name: str) -> str | None:
    ext = ".exe" if name.lower().endswith(".exe") else ""
    stem = name[: -len(ext)] if ext else name
    stem = stem.lower()
    if re.fullmatch(r"sikuli(?:-go|go|grpc)(?:-[0-9a-f]{8,})?", stem):
        return f"sikuli-go{ext}"
    if re.fullmatch(r"sikuli(?:-go|go|grpc)-monitor(?:-[0-9a-f]{8,})?", stem):
        return f"sikuli-go-monitor{ext}"
    return None


def _runtime_source_rank(name: str) -> int:
    ext = ".exe" if name.lower().endswith(".exe") else ""
    stem = name[: -len(ext)] if ext else name
    stem = stem.lower()
    if stem in ("sikuli-go", "sikuli-go-monitor"):
        return 0
    if stem in ("sikuligo", "sikuligo-monitor"):
        return 1
    if stem in ("sikuligrpc", "sikuligrpc-monitor"):
        return 2
    if stem.startswith("sikuli-go-monitor-") or stem.startswith("sikuli-go-"):
        return 3
    if stem.startswith("sikuligo-monitor-") or stem.startswith("sikuligo-"):
        return 4
    if stem.startswith("sikuligrpc-monitor-") or stem.startswith("sikuligrpc-"):
        return 5
    return 6


def _discover_runtime_sources(primary: Path) -> dict[str, Path]:
    runtimes: dict[str, tuple[int, Path]] = {}
    primary_canonical = _runtime_canonical_name(primary.name)
    if primary_canonical:
        runtimes[primary_canonical] = (_runtime_source_rank(primary.name), primary)
    for entry in primary.parent.iterdir():
        if not entry.is_file():
            continue
        if not re.match(r"^sikuli.*(\.exe)?$", entry.name, flags=re.IGNORECASE):
            continue
        canonical = _runtime_canonical_name(entry.name)
        if not canonical:
            continue
        rank = _runtime_source_rank(entry.name)
        current = runtimes.get(canonical)
        if current is None or rank < current[0] or (rank == current[0] and str(entry) < str(current[1])):
            runtimes[canonical] = (rank, entry)
    return {name: entry for name, (_, entry) in runtimes.items()}


def _cleanup_installed_runtime_aliases(target_dir: Path, exempt: set[Path]) -> None:
    for entry in target_dir.iterdir():
        if not entry.is_file():
            continue
        if entry.resolve() in exempt:
            continue
        canonical = _runtime_canonical_name(entry.name)
        if canonical and entry.name != canonical:
            entry.unlink()


def _run_install_binary(args: argparse.Namespace) -> int:
    target_dir = Path(args.dir).expanduser().resolve() if args.dir else (Path.home() / ".local" / "bin")
    target_dir.mkdir(parents=True, exist_ok=True)
    primary = Path(_resolve_sikuli_binary(None))
    runtimes = _discover_runtime_sources(primary)
    _cleanup_installed_runtime_aliases(target_dir, {runtime.resolve() for runtime in runtimes.values()})
    copied: list[Path] = []
    for target_name in sorted(runtimes):
        runtime = runtimes[target_name]
        target = target_dir / target_name
        shutil.copy2(runtime, target)
        if os.name != "nt":
            target.chmod(0o755)
        copied.append(target)
        print(target)

    if not args.no_shell_update:
        detected = _detect_shell_profile()
        if detected:
            profile, source_cmd = detected
            should_update = bool(args.yes) or _prompt_yes_no(f"Add {target_dir} to PATH in {profile}?")
            if should_update:
                _ensure_path_export(profile, target_dir)
                print(f"Run {source_cmd} to reload PATH in this shell.")
                return 0
    print(f"Ensure {target_dir} is on PATH for new shells.")
    return 0


def _build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(prog="sikuli-go", add_help=True)
    subparsers = parser.add_subparsers(dest="command")

    init_py = subparsers.add_parser(
        "init:py-examples",
        help="Scaffold a Python project, create .venv, install requirements, copy examples",
    )
    init_py.add_argument("--dir", default=None, help="Target project directory")
    init_py.add_argument("--skip-install", action="store_true", help="Skip pip install in .venv")

    init_py_alias = subparsers.add_parser(
        "init-py-examples",
        help="Alias for init:py-examples",
    )
    init_py_alias.add_argument("--dir", default=None, help="Target project directory")
    init_py_alias.add_argument("--skip-install", action="store_true", help="Skip pip install in .venv")

    install_binary = subparsers.add_parser(
        "install-binary",
        help="Copy sikuli-go runtimes to a PATH-ready directory",
    )
    install_binary.add_argument("--dir", default=None, help="Target directory (default: ~/.local/bin)")
    install_binary.add_argument("--yes", action="store_true", help="Automatically update shell profile PATH when detected")
    install_binary.add_argument("--no-shell-update", action="store_true", help="Do not modify shell profile")
    return parser


def main(argv: list[str] | None = None) -> int:
    parser = _build_parser()
    args = parser.parse_args(argv)

    if args.command in ("init:py-examples", "init-py-examples"):
        return _run_init_py_examples(args)
    if args.command == "install-binary":
        return _run_install_binary(args)

    parser.print_help()
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
