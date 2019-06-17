#!/usr/bin/env python3

"""Script for extracting a file or directory from a docker image"""

import pathlib
import subprocess
import sys

def main():
    docker_image, source, destination = parse_args()
    container = docker_create(docker_image)
    docker_cp(container, source, destination)
    docker_rm(container)

def docker_create(docker_image):
    result = subprocess.run(
            ['docker', 'create', docker_image],
            stdout=subprocess.PIPE)
    result.check_returncode()
    container = result.stdout.decode('utf-8')
    return container.rstrip('\n')

def docker_cp(container, source, destination):
    container_source = "{}:{}".format(container, source)
    result = subprocess.run(['docker', 'cp', container_source, destination])
    result.check_returncode()

def docker_rm(container):
    result = subprocess.run(
            ['docker', 'rm', container],
            stdout=subprocess.DEVNULL)
    result.check_returncode()

def parse_args():
    if len(sys.argv) != 4:
        script_filename = pathlib.Path(__file__).resolve().name
        print("USAGE: {} <docker_image> <source> <destination>".format(script_filename))
        sys.exit()

    docker_image = sys.argv[1]
    source = sys.argv[2]
    destination = sys.argv[3]
    return (docker_image, source, destination)

main()
