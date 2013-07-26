War - Watch and repeat
======================

War will watch a file or folder (and all of its content) and run a command
every time a change is detected. This is quite handy if for example you want
to run a test whenever a change is detected in your source code.

War uses [GNU Bash](https://www.gnu.org/software/bash/) to run the command
and because of that, it is required that bash is installed on your machine.
This should be the case for any modern UNIX based operating system like
Mac OS X or Linux.

Besides that, the only other requirement is that [Go](http://golang.org) is
[installed](http://golang.org/doc/install) and set up correctly on the
host machine.

## Installation

    $ go get https://github.com/ChrisBuchholz/war

## Usage

    $ war /my/project/src "make test"

This will run `make test` every time a change is detected in /my/project/src.

The way I have been developing War, is by running the following command which
continuesly builds War as I work on it. It will output whatever `go build`
outputs, or 'No errors' if the code compiles.

    $ war . "go build && [ $? -eq 0 ] && echo 'No errors' && rm war"
