#!/usr/bin/env python3

import pathlib
import shutil
import subprocess
import sys

def main():
    check_python3()
    check_go()
    check_golint()

    link = get_hooks_dir() / 'pre-push'
    script = pathlib.Path(__file__).resolve().parent / 'pre-push.py'
    symlink(link, script)

def check_go():
    if shutil.which("go") == None:
        print("ERROR: go is not installed")
        print("To install: https://golang.org/doc/install")
        sys.exit(1)

def check_golint():
    if shutil.which("golint") == None:
        print("ERROR: golint is not installed")
        print("To install: https://github.com/golang/lint")
        sys.exit(1)

def check_python3():
    if shutil.which("python3") == None:
        print("ERROR: python3 is not in PATH")
        print("Create a symlink 'python3' in PATH")
        sys.exit(1)

def get_hooks_dir() -> pathlib.Path:
    result = subprocess.run(
            ['git', 'rev-parse', '--git-dir'],
            check=True,
            stdout=subprocess.PIPE)
    git_dir = result.stdout.decode('utf-8').rstrip('\n')
    return pathlib.Path(git_dir) / 'hooks'

def symlink(link: pathlib.Path, target: pathlib.Path):
    if link.exists() or link.is_symlink():
        print("ERROR: {} exists".format(link))
        sys.exit(1)
    link.symlink_to(target)

main()
