package main

import (
	"fmt"
	"strings"
)

func main() {
	target := "Hello World"
	target = strings.Map(replaceSpace, target)
	fmt.Println(target)
}

func replaceSpace(r rune) rune {
	if r == ' ' {
		r = rune('/')
	}
	return r
}
