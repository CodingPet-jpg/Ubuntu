package main

import (
	"fmt"
	"os"
)

var a int = b + c
var b int = f()
var c int = 1

func main() {
	fmt.Fprintf(os.Stdout, "a:%v\tb:%v\tc:%v\n", a, b, c)
}

func f() int {
	return c + 1
}

// 一个包可以包含任意数量的init函数，不会触发重名
func init() {

}
func init() {

}
func init() {

}
