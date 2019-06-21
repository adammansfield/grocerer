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
BASENAME := $(PYTHON3) scripts/make/basename.py
CLEAN := $(PYTHON3) scripts/make/clean.py
EXTRACT := $(PYTHON3) scripts/make/extract.py
FIND := $(PYTHON3) scripts/make/find.py
HELP := $(PYTHON3) scripts/make/help.py
TOUCH := $(PYTHON3) scripts/make/touch.py
VERSION := $(PYTHON3) scripts/make/version.py
