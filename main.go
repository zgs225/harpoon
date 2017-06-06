package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "\tharpoon command <arguments>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "\tinit\t进行初始化，生成工具脚本")
	fmt.Fprintln(os.Stderr, "")
}

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
	f, err := os.OpenFile(".harpoon", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprint(f, s)
	fmt.Fprint(os.Stderr, "[i] Initilizing...\n")
	fmt.Fprint(os.Stderr, s)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "init":
		hInit(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}
