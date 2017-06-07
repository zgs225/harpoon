package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Build docker image

func hBuild() {
	c := loadConfig()
	c.check()
	fmt.Printf("[i] Building docker image %v\n", c.Image)
	cmd := exec.Command("docker", "build", "-t", c.Image, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("[i] Error.")
	} else {
		fmt.Println("[i] Done.")
	}
}
