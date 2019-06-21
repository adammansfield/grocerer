#!/usr/bin/env python3

"""
Cross-platform script for this project's `make help`

Print the target and description that matches the form:
    <target>: [<prerequisites>] ## <description>

Example:
    Makefile:
        build: $(SRC) ## Build the container
        deploy: ## Deploy the container
    Output:
        build   Build the container
        deploy  Deploy the container
"""

import pathlib
import re

def main():
    regex = r"^([a-zA-Z_-]+):.*##\s*(.+)$"
    makefile = pathlib.Path(__file__).resolve().parent.parent.parent / "Makefile"
    with open(str(makefile)) as stream:
        matches = re.findall(regex, stream.read(), re.MULTILINE)

    width = max([len(match[0]) for match in matches]) + 2
    for match in matches:
        print("{command:{width}}{description}".format(
            command=match[0],
            width=width,
            description=match[1]))

main()
