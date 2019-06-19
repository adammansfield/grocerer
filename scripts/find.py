#!/usr/bin/env python3

"""Cross-platform script for searching for files in a directory hierarchy"""

import pathlib
import sys

if not 3 <= len(sys.argv) <= 4:
    script_filename = pathlib.Path(__file__).resolve().name
    print("USAGE: {} <path> <pattern> [<exclude>]".format(script_filename))
    sys.exit()

search_path = sys.argv[1]
pattern = sys.argv[2]
exclude = len(sys.argv) == 4
if exclude:
    exclude_pattern = sys.argv[3]

for path in pathlib.Path(search_path).rglob(pattern):
    if exclude and path.match(exclude_pattern):
        continue
    print(path, end=' ')
