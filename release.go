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
	if err != nil {
		fmt.Println("[i] Tag error.")
		return
	}
	err = push(name)
	if err != nil {
		fmt.Println("[i] Push error.")
		return
	}

	err = writeVersion(args[0])
	if err != nil {
		fmt.Println("[i] Write version error.")
		return
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

func writeVersion(v string) error {
	vf := "VERSION"
	f, err := os.OpenFile(vf, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(v + "\n")
	if err != nil {
		return err
	}

	c := newCommand("git", "add", vf)
	err = c.Run()
	if err != nil {
		return err
	}

	c = newCommand("git", "commit", "--amend", "--no-edit")
	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}
