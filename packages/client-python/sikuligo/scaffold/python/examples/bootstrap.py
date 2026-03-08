from __future__ import annotations

import os
import hashlib
import shutil
from pathlib import Path

from sikuligo.sikulix import _resolve_sikuli_binary


def ensure_sikuli_go_on_path() -> str:
    existing = shutil.which("sikuli-go") or shutil.which("sikuligo")
    if existing:
        os.environ["SIKULI_GO_BINARY_PATH"] = existing
        return existing

    source = Path(_resolve_sikuli_binary(None))
    install_dir = Path.cwd() / ".sikuli-go" / "bin"
    install_dir.mkdir(parents=True, exist_ok=True)
    binary_name = "sikuli-go.exe" if os.name == "nt" else "sikuli-go"
    installed = install_dir / binary_name

    should_copy = True
    if installed.exists():
        should_copy = _file_digest(source) != _file_digest(installed)
    if should_copy:
        shutil.copy2(source, installed)
        if os.name != "nt":
            installed.chmod(0o755)

    current_path = os.environ.get("PATH", "")
    entries = [entry for entry in current_path.split(os.pathsep) if entry]
    if str(install_dir) not in entries:
        os.environ["PATH"] = f"{install_dir}{os.pathsep}{current_path}" if current_path else str(install_dir)

    os.environ["SIKULI_GO_BINARY_PATH"] = str(installed)
    return str(installed)


def _file_digest(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()
