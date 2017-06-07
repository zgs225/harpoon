package main

import (
	"flag"
	"fmt"
	"os"
)

func hInit(args []string) {
	f := flag.NewFlagSet("harpoon init", flag.PanicOnError)
	repo := f.String("repo", "", "Docker registry host.")
	imag := f.String("image", "", "The name of building docker image.")
	erro := f.Parse(args)

	if erro != nil {
		panic(erro)
	}

	if len(*repo) == 0 || len(*imag) == 0 {
		f.Usage()
		os.Exit(1)
	}

	writeConfig(*repo, *imag)
	fmt.Fprintln(os.Stderr, "[i] Done.\n")
}

func writeConfig(repo, image string) {
	s := fmt.Sprintf("repo=%s\nimage=%s\n", repo, image)
	c := &config{
		Repo:  repo,
		Image: image,
	}
	c.writeToDisk()
	fmt.Fprint(os.Stderr, "[i] Initilizing...\n")
	fmt.Fprint(os.Stderr, s)
}
