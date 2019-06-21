#!/usr/bin/env python3

import pathlib
import shutil
import subprocess
import sys

def main():
    check_python3()
    get_pre_commit_symlink().symlink_to(get_pre_commit_file())

def check_python3():
    if shutil.which("python3") == None:
        print("ERROR: python3 is not in PATH")
        print("Create a symlink 'python3' in PATH")
        sys.exit(1)

def get_pre_commit_symlink():
    result = subprocess.run(
            ['git', 'rev-parse', '--git-dir'],
            check=True,
            stdout=subprocess.PIPE)
    git_directory = result.stdout.decode('utf-8').rstrip('\n')
    pre_commit_symlink = pathlib.Path(git_directory) / 'hooks' / 'pre-commit'
    if (pre_commit_symlink.exists()):
        print("ERROR: {} exists".format(pre_commit_symlink))
        sys.exit(1)
    return pre_commit_symlink

def get_pre_commit_file():
    scripts_directory = pathlib.Path(__file__).resolve().parent
    return scripts_directory / 'pre-commit.py'

main()
