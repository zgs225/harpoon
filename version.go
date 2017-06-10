package main

import (
	"fmt"
)

func hVersion() {
	c := loadConfig()
	c.check()

	if len(c.Version) > 0 {
		fmt.Printf("image %s current version is %s\n", c.Image, c.Version)
	} else {
		fmt.Printf("image %s no version\n", c.Image)
	}
}
