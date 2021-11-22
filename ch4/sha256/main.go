package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	zero(&s2)
	zeronew(&s1)
	fmt.Printf("%x\n%x\n%t\n%T\n", s1, s2, s1 == s2, s1)
}
func zero(b *[32]byte) {
	for v := range *b {
		b[v] = 0
	}
}
func zeronew(b *[32]byte) {
	*b = [32]byte{}
}
