# Run `make <target> verbose=1` to echo every command
ifndef verbose
MAKEFLAGS += --silent
endif
