import logging
import os
import shutil
from functools import cached_property
from pathlib import Path
from subprocess import run
from typing import Any

import pathspec

from hatchling.builders.hooks.plugin.interface import BuildHookInterface

log = logging.getLogger(__name__)
log_level = logging.getLevelName(os.getenv("HATCH_BUILD_SCRIPTS_LOG_LEVEL", "INFO"))
log.setLevel(log_level)


class ScriptConfig:
    def __init__(
        self,
        commands=[],
        artifacts=[],
        out_dir=".",
        work_dir=".",
        clean_artifacts=True,
        clean_out_dir=False,
    ) -> None:
        self.out_dir = out_dir
        self.work_dir = work_dir
        self.commands = commands
        self.artifacts = artifacts
        self.out_dir = out_dir
        self.work_dir = work_dir
        self.clean_artifacts = clean_artifacts
        self.clean_out_dir = clean_out_dir

    def work_files(self, root: str | Path, *, relative: bool = False) -> list[Path]:
        """Get files in the work directory that match the artifacts spec."""
        abs_dir = Path(root, self.work_dir)
        if not abs_dir.exists():
            return []
        return [
            Path(f) if relative else abs_dir / f
            for f in self.artifacts_spec.match_tree(abs_dir)
        ]

    def out_files(self, root: str | Path, *, relative: bool = False) -> list[Path]:
        """Get files in the output directory that match the artifacts spec."""
        abs_dir = Path(root, self.out_dir)
        if not abs_dir.exists():
            return []
        return [
            Path(f) if relative else abs_dir / f
            for f in self.artifacts_spec.match_tree(abs_dir)
        ]

    @cached_property
    def artifacts_spec(self) -> pathspec.PathSpec:
        """A pathspec for the artifacts."""
        return pathspec.PathSpec.from_lines(
            pathspec.patterns.GitWildMatchPattern, self.artifacts
        )


def load_scripts(config: dict[str, Any]) -> list[ScriptConfig]:
    cs = []
    for script_config in config.get("scripts", []):
        c = ScriptConfig(**script_config)
        cs.append(c)
    return cs


class CustomBuilder(BuildHookInterface):
    PLUGIN_NAME = "build-scripts"

    def initialize(
        self,
        version: str,  # noqa: ARG002
        build_data: dict[str, Any],
    ) -> None:
        created: set[Path] = set()

        all_scripts = load_scripts(self.config)

        for script in all_scripts:
            if script.clean_out_dir:
                out_dir = Path(self.root, script.out_dir)
                log.debug(f"Cleaning {out_dir}")
                shutil.rmtree(out_dir, ignore_errors=True)
            elif script.clean_artifacts:
                for out_file in script.out_files(self.root):
                    log.debug(f"Cleaning {out_file}")
                    out_file.unlink(missing_ok=True)

        for script in all_scripts:
            # breakpoint()
            # log.debug(f"Script config: {asdict(script)}")
            work_dir = Path(self.root, script.work_dir)
            out_dir = Path(self.root, script.out_dir)
            out_dir.mkdir(parents=True, exist_ok=True)

            for cmd in script.commands:
                log.info(f"Running command: {cmd}")
                run(cmd, cwd=str(work_dir), check=True, shell=True)  # noqa: S602

            log.info(f"Copying artifacts to {out_dir}")
            for work_file in script.work_files(self.root, relative=True):
                src_file = work_dir / work_file
                out_file = out_dir / work_file
                log.debug(f"Copying {src_file} to {out_file}")
                if src_file not in created:
                    out_file.parent.mkdir(parents=True, exist_ok=True)
                    shutil.copyfile(src_file, out_file)
                    created.add(out_file)
                else:
                    log.debug(f"Skipping {src_file} - already exists")

            build_data["artifacts"].append(str(out_dir.relative_to(self.root)))
