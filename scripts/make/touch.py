#!/usr/bin/env python3

"""Cross-platform script for updating a file's timestamp"""

import pathlib
import sys

if len(sys.argv) != 2:
    script_filename = pathlib.Path(__file__).resolve().name
    print("USAGE: {} <path>".format(script_filename))
    sys.exit()

pathlib.Path(sys.argv[1]).touch()
