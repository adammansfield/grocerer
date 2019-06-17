#!/usr/bin/env python3

"""Cross-platform script for searching for files in a directory hierarchy"""

import pathlib
import sys

def main():
    if len(sys.argv) != 3:
        script_filename = pathlib.Path(__file__).resolve().name
        print("USAGE: {} <path> <pattern>".format(script_filename))
        sys.exit()

    search_path = sys.argv[1]
    pattern = sys.argv[2]
    for path in pathlib.Path(search_path).rglob(pattern):
        print(path, end=' ')

main()
