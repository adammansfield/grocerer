#!/usr/bin/env python3

import pathlib
import shutil
import subprocess
import sys

def main():
    check_python3()
    check_go()
    check_golint()

    hooks_dir = get_hooks_dir()
    pre_commit_link = hooks_dir / 'pre-commit'
    pre_push_link = hooks_dir / 'pre-push'

    scripts_dir = pathlib.Path(__file__).resolve().parent
    pre_commit_script = scripts_dir / 'pre-commit.py'
    pre_push_script = scripts_dir / 'pre-push.py'

    symlink(pre_commit_link, pre_commit_script)
    symlink(pre_push_link, pre_push_script)

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
