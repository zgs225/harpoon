package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Harpoon is a tool for simplify building and releasing docker image.")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "\tharpoon command <arguments>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "\tinit   \t进行初始化，生成工具脚本")
	fmt.Fprintln(os.Stderr, "\tbuild  \t编译Docker镜像")
	fmt.Fprintln(os.Stderr, "\trelease\t将镜像打包版本并推送到仓库中")
	fmt.Fprintln(os.Stderr, "")
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
	case "build":
		hBuild()
	case "release":
		hRelease(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}
