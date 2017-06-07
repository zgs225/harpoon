package main

import (
	"fmt"
)

// Build docker image

func hBuild() {
	c := loadConfig()
	fmt.Println(*c)
}
