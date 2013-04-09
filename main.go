package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	NAME        string = "War"
	DESCRIPTION string = "Watch and Repeat a command when a change is " +
		"detected in the directory or file that are being listened to."
	VERSION string = "0.1.0"
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s - %s\n\n", NAME, DESCRIPTION)
	fmt.Fprintf(os.Stderr, "Version %s\n\n", VERSION)
	fmt.Fprintf(os.Stderr, "usage: %s [path] [command]\n", NAME)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	// there must be given both a path and a command, if not,
	// show usage
	if len(args) != 2 {
		flag.Usage()
	}

	path := args[0]
	command_str := args[1]

	command := new(Command)
	command.Cmd = command_str
	err := Watch(path, *command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
