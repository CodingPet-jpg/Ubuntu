package main

import (
	"fmt"

	"github.com/CcodingPet-jpg/Uubuntu/ch6/geometry"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter = 1
	fmt.Printf("before use:%d\n", c)
	fmt.Fprintf(&c, "%s", "sdaf")
	fmt.Printf("after use:%d\n", c)
	p := geometry.Point{X: 1, Y: 2}
	fmt.Println(p)
}
