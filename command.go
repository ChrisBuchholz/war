package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
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
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if c.isRunning {
		log.Println("restarting process...")
		if err := c.process.Process.Kill(); err != nil {
			log.Println("failed to stop process.")
		}
	}

	c.isRunning = true

	c.process = exec.Command(getBash(), "-c", c.Cmd)

	c.process.Stdout = &stdout
	c.process.Stderr = &stderr
	c.process.Stdin = os.Stdin

	err := c.process.Run()

	clock := time.Now().Format("15:04:05")

	if err != nil {
		fmt.Printf("Command failed... [%s]\n", clock)
		fmt.Printf("----------------------------\n")

		if strings.TrimSpace(err.Error()) != "" {
			fmt.Printf("%s\n", err)
		}
	} else {
		fmt.Printf("Running command... [%s]\n", clock)
		fmt.Printf("-----------------------------\n")
	}

	if strings.TrimSpace(stdout.String()) != "" {
		fmt.Printf("%s\n", stdout.String())
	}

	if strings.TrimSpace(stderr.String()) != "" {
		fmt.Printf("%s\n", stderr.String())
	}

	c.isRunning = false
}
