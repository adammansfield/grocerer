#!/usr/bin/env python3

"""Cross-platform script for updating a file's timestamp"""

import pathlib
import sys

pathlib.Path(sys.argv[1]).touch()
