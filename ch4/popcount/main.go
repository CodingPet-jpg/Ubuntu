package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func main() {
	s := "Hello World"
	sha := sha256.Sum256([]byte(s))
	fin := count(sha)
	fmt.Println(fin)
	//final()int转byte时会取最后四位
}
func count(b [32]byte) int {
	sum := 0
	for ad := range b {
		sum += int(pc[int(ad)])
	}
	return sum
}
func init() {
	for n, _ := range pc {
		pc[n] = pc[n/2] + byte(n&1)
		fmt.Printf("num:%d\tcount:%d\n", n, pc[n])
	}
}

func final() {
	var i int64 = 0x1234567876543212
	fmt.Println(byte(i >> 8))
}
