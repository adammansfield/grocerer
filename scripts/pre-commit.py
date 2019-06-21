#!/usr/bin/env python3

import fnmatch
import pathlib
import subprocess
import sys
from typing import List

def main():
    toplevel_dir = get_toplevel_dir()
    files = get_diff_files(toplevel_dir)

    go_files = [f for f in files if f.match('*.go')]
    if go_files:
        check_gofmt(go_files)
        check_golint(go_files)

    make('build', toplevel_dir)
    make('test', toplevel_dir)

def check_gofmt(go_files: List[pathlib.Path]):
    unformatted_files = gofmt(go_files)
    if unformatted_files:
        print("Go files must be formatted with gofmt. Please run:")
        for unformatted_file in unformatted_files:
            print("  gofmt -s -w {}".format(unformatted_file))
        sys.exit(1)

def check_golint(go_files: List[pathlib.Path]):
    warnings = golint(go_files)
    if warnings:
        print("Go files must pass golint. Please fix:")
        for warning in warnings:
            print("  {}".format(warning))
        sys.exit(1)

def get_toplevel_dir() -> pathlib.Path:
    result = subprocess.run(
            ['git', 'rev-parse', '--show-toplevel'],
            check=True,
            stdout=subprocess.PIPE)
    return pathlib.Path(result.stdout.decode('utf-8').rstrip('\n'))

def get_diff_files(toplevel_dir: pathlib.Path) -> List[pathlib.Path]:
    result = subprocess.run(
            ['git', 'diff', '--cached', '--name-only', '--diff-filter=ACM'],
            check=True,
            stdout=subprocess.PIPE)
    files = result.stdout.decode('utf-8').rstrip('\n').splitlines()
    return [toplevel_dir / f for f in files]

def gofmt(files: List[pathlib.Path]) -> List[str]:
    files_arg = ' '.join([str(f) for f in files])
    result = subprocess.run(
            ['gofmt', '-l', '-s', files_arg],
            check=True,
            stdout=subprocess.PIPE)
    return result.stdout.decode('utf-8').rstrip('\n').splitlines()

def golint(files: List[pathlib.Path]) -> List[str]:
    files_arg = ' '.join([str(f) for f in files])
    result = subprocess.run(
            ['golint', files_arg],
            check=True,
            stdout=subprocess.PIPE)
    return result.stdout.decode('utf-8').rstrip('\n').splitlines()

def make(command: str, toplevel_dir: pathlib.Path):
    print("make {}".format(command))
    result = subprocess.run(
            ['make', '-C', str(toplevel_dir), command],
            stderr=subprocess.STDOUT,
            stdout=subprocess.PIPE)
    if result.returncode != 0:
        print("make {} failed:".format(command))
        print(result.stdout.decode('utf-8').rstrip('\n'))
        sys.exit(1)

main()
