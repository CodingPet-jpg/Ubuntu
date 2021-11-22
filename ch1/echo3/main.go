package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Println(i)
		fmt.Println(arg)
	}
}
