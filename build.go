package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// Build docker image

func hBuild(args []string) {
	f := flag.NewFlagSet("harpoon build", flag.PanicOnError)
	noCache := f.Bool("no-cache", false, "是否使用缓存")
	erro := f.Parse(args)
	if erro != nil {
		panic(erro)
	}

	c := loadConfig()
	c.check()
	fmt.Printf("[i] Building docker image %v\n", c.Image)
	cmdArgs := []string{"build", "-t", c.Image}
	if *noCache {
		cmdArgs = append(cmdArgs, "--no-cache")
	}
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
