import contextlib
from typing import Any
import subprocess

from hatchling.builders.hooks.plugin.interface import BuildHookInterface


@contextlib.contextmanager
def CD(d: str):
    import os

    old = os.getcwd()
    os.chdir(d)
    yield
    os.chdir(old)


class CustomBuilder(BuildHookInterface):
    def initialize(
        self,
        version: str,  # noqa: ARG002
        build_data: dict[str, Any],
    ) -> None:
        if self.target_name == "sdist":
            return

        with CD("diskstat"):
            subprocess.run("go build diskstat-api.go".split(), check=True)
            subprocess.run("pyside6-rcc res.qrc -o res.py".split(), check=True)
