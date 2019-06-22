#!/usr/bin/env python3
 
import os
import pathlib
import subprocess
import sys

def main():
    os.chdir(str(get_toplevel_dir()))
    make('test-large')

def get_toplevel_dir() -> pathlib.Path:
    result = subprocess.run(
            ['git', 'rev-parse', '--show-toplevel'],
            check=True,
            stdout=subprocess.PIPE)
    return pathlib.Path(result.stdout.decode('utf-8').rstrip('\n'))

def make(command: str):
    print("make {}".format(command))
    result = subprocess.run(
            ['make', command],
            stderr=subprocess.STDOUT,
            stdout=subprocess.PIPE)
    if result.returncode != 0:
        print("make {} failed:".format(command))
        print(result.stdout.decode('utf-8').rstrip('\n'))
        sys.exit(1)

main()
