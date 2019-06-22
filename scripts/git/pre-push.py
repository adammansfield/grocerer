#!/usr/bin/env python3

import pathlib
import subprocess
import sys

def main():
    toplevel_dir = get_toplevel_dir()
    make('test-large', toplevel_dir)

def get_toplevel_dir() -> pathlib.Path:
    result = subprocess.run(
            ['git', 'rev-parse', '--show-toplevel'],
            check=True,
            stdout=subprocess.PIPE)
    return pathlib.Path(result.stdout.decode('utf-8').rstrip('\n'))

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
