package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println("Hello, Let's Go!!!")

	cmdstr := "/bin/sh"
	cmd := exec.Command(cmdstr)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PS1=-[letsgo]- # "}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the %s command - %s\n", cmdstr, err)
		os.Exit(1)
	}
}
