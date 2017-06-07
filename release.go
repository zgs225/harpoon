package main

import (
	"fmt"
	"os"
	"os/exec"
)

func hRelease(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage of harpoon release:\n")
		fmt.Fprintln(os.Stderr, "\tharpoon release <VERSION>")
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}

	c := loadConfig()
	c.check()

	name := fmt.Sprintf("%s/%s:%s", c.Repo, c.Image, args[0])
	fmt.Printf("[i] Releasing to %s, Version %s\n", c.Repo, args[0])
	err := tag(c.Image, name)
	fmt.Println("[i] Tagging...")
	if err != nil {
		fmt.Println("[i] Tag error.")
		os.Exit(1)
	}
	fmt.Println("[i] Pushing...")
	err = push(name)
	if err != nil {
		fmt.Println("[i] Push error.")
		os.Exit(1)
	}

	fmt.Println("[i] Writing version...")
	c.Version = args[0]
	err = c.writeToDisk()
	if err != nil {
		fmt.Println("[i] Config write to dist error")
		os.Exit(1)
	}

	fmt.Println("[i] Done.")
}

func tag(o, d string) error {
	c := newCommand("docker", "tag", o, d)
	return c.Run()
}

func push(d string) error {
	c := newCommand("docker", "push", d)
	return c.Run()
}

func newCommand(name string, arg ...string) *exec.Cmd {
	c := exec.Command(name, arg...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c
}
