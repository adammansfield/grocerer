#!/usr/bin/env python3

"""Cross-platform script to create this project's version.go"""

import datetime
import pathlib
import subprocess
import textwrap

def main():
    project_directory = pathlib.Path(__file__).resolve().parent.parent.parent
    path = project_directory / "internal/go/version.go"
    if path.exists():
        pathlib.Path.unlink(path)

    date = get_last_commit_date()
    commit = get_last_commit_hash()
    code = create_code(commit, date)
    with open(str(path), "w+") as handle:
        handle.write(code)

def create_code(commit, date):
    args = { 'commit': commit, 'date': date }
    return textwrap.dedent(
        """\
        package openapi

        // PackageVersion contains version info
        var PackageVersion = Version{{
        	Api:    1,
        	Commit: "{commit}",
        	Date:   "{date}"}}
        """.format(**args))

def get_last_commit_date():
    result = subprocess.run(
            ['git', 'log', '-1', '--format=%at'],
            stdout=subprocess.PIPE)
    result.check_returncode()
    timestamp = result.stdout.decode('utf-8').rstrip('\n')
    date = datetime.datetime.utcfromtimestamp(int(timestamp))
    return date.strftime('%Y-%m-%dT%H:%MZ')

def get_last_commit_hash():
    result = subprocess.run(
            ['git', 'log', '-1', '--format=%h'],
            stdout=subprocess.PIPE)
    result.check_returncode()
    date = result.stdout.decode('utf-8')
    return date.rstrip('\n')

main()
