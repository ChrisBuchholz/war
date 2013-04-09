package main

import (
	"log"
	"os"
	"os/exec"
)

// look up the absolute path to where bash is installed on the host
func getBash() string {
	bash, err := exec.LookPath("bash")
	if err != nil {
		log.Fatal("Looks like you don't have bash installed...")
	}
	return bash
}

type Command struct {
	Cmd       string
	process   *exec.Cmd
	isRunning bool
}

// execute string Command.Cmd as a bash command
//
// if a process is already running, it will kill that one
// and before spawning a new one
func (c *Command) Execute() {
	if c.isRunning {
		log.Println("restarting process...")
		if err := c.process.Process.Kill(); err != nil {
			log.Println("failed to stop process.")
		}
	}

	c.isRunning = true

	c.process = exec.Command(getBash(), "-c", c.Cmd)
	c.process.Stderr = os.Stderr
	c.process.Stdin = os.Stdin
	c.process.Stdout = os.Stdout
	if err := c.process.Run(); err != nil {
		log.Println("failed to run process.")
	}

	c.isRunning = false
}
