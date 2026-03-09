from __future__ import annotations

import os
import tempfile
import unittest
from pathlib import Path

from sikuligo.sikulix import _packaged_runtime_platform_dirs, _resolve_packaged_runtime_binary


class PackagedRuntimeResolutionTest(unittest.TestCase):
    def test_resolve_packaged_runtime_binary_uses_current_platform_dir(self) -> None:
        platform_dirs = _packaged_runtime_platform_dirs()
        self.assertTrue(platform_dirs, "expected at least one packaged runtime dir for this platform")

        with tempfile.TemporaryDirectory(prefix="sikuli-go-python-runtime-") as tmpdir:
            package_root = Path(tmpdir)
            runtime_dir = package_root / "runtime" / platform_dirs[0] / "bin"
            runtime_dir.mkdir(parents=True, exist_ok=True)
            binary_name = "sikuli-go.exe" if os.name == "nt" else "sikuli-go"
            binary_path = runtime_dir / binary_name
            binary_path.write_bytes(b"MZ" if os.name == "nt" else b"\xcf\xfa\xed\xfe")
            if os.name != "nt":
                binary_path.chmod(0o755)

            resolved = _resolve_packaged_runtime_binary((binary_name,), package_root=package_root)
            self.assertEqual(resolved, str(binary_path.resolve()))


if __name__ == "__main__":
    unittest.main()
