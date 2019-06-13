#!/usr/bin/env python3

"""Cross-platform script for updating a file's timestamp"""

import os
import subprocess
import sys

if os.name == 'nt':
    command = "(ls {0}).LastWriteTime = Get-Date".format(sys.argv[1])
    subprocess.run(['powershell', '-NoProfile', '–ExecutionPolicy', 'Bypass', command])
else:
    subprocess.run(['touch', sys.argv[1]])
