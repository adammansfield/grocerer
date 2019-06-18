#!/usr/bin/env python3

"""Cross-platform script for this project's `make clean`"""

import pathlib
import shutil

def main():
    project_directory = pathlib.Path(__file__).resolve().parent.parent

    cleaner = ProjectCleaner(project_directory)
    cleaner.remove_file("internal/go/logger.go")
    cleaner.remove_file("internal/go/model_version.go")
    cleaner.remove_file("internal/go/routers.go")
    cleaner.remove_file("internal/go/version.go")
    cleaner.remove_file("internal/Dockerfile")
    cleaner.remove_file("internal/main.go")
    cleaner.remove_directory("bin")
    cleaner.remove_directory("gen")

    # TODO: remove below when openapi-generator is removed
    cleaner.remove_file("internal/go/api_default.go")
    cleaner.remove_file("internal/.openapi-generator-ignore")
    cleaner.remove_file("internal/README.md")
    cleaner.remove_directory("internal/.openapi-generator")
    cleaner.remove_directory("internal/api")
    cleaner.remove_empty_directory("internal/go")
    cleaner.remove_empty_directory("internal")

class ProjectCleaner:
    def __init__(self, project_directory):
        self.project_directory = project_directory

    def remove_directory(self, directory):
        path = self.project_directory / directory
        if path.exists():
            shutil.rmtree(str(path))

    def remove_empty_directory(self, directory):
        path = self.project_directory / directory
        if not path.exists():
            return
        is_empty = not bool(sorted(path.rglob('*')))
        if is_empty:
            pathlib.Path.rmdir(path)

    def remove_file(self, filename):
        path = self.project_directory / filename
        if path.exists():
            pathlib.Path.unlink(path)

main()
