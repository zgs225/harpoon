package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Build docker image

func hBuild(args []string) {
	c := loadConfig()
	c.check()
	fmt.Printf("[i] Building docker image %v\n", c.Image)

	cmdArgs := []string{"build", "-t", c.Image}
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, ".")

	cmd := exec.Command("docker", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("[i] Error.")
	} else {
		fmt.Println("[i] Done.")
	}
}
