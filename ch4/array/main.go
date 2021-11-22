package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var arr []byte = []byte("x")
	b := sha256.Sum256(arr)
	fmt.Println(b)
}
