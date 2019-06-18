# Provide common functions on both Linux and Windows by using Python

ifeq ($(OS),Windows_NT)
	PYTHON3 := python
	# TODO: Remove below when openapi-generator is removed
	CP := xcopy /s
	RM := del
	SEP := \\
else
	PYTHON3 := /usr/bin/env python3
	# TODO: Remove below when openapi-generator is removed
	CP := cp -r --no-target-directory
	RM := rm -f
	SEP := /
endif

# Use uppercase to avoid naming conflicts.
# e.g. $(shell $(basename) $(CURDIR)) would conflict with make's basename function.
BASENAME := $(PYTHON3) scripts/basename.py
CLEAN := $(PYTHON3) scripts/clean.py
EXTRACT := $(PYTHON3) scripts/extract.py
FIND := $(PYTHON3) scripts/find.py
HELP := $(PYTHON3) scripts/help.py
TOUCH := $(PYTHON3) scripts/touch.py
VERSION := $(PYTHON3) scripts/version.py
