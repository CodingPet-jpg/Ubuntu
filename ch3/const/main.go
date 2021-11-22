package main

import (
	"fmt"
	"time"
)

const pi = 3.1415926
const NoDelay time.Duration = 0
const timeout = 5 * time.Minute
const (
	a = iota + 1 // 常量生成器
	b
	c
	d = 2
	e
)

func main() {
	fmt.Printf("%T\t%[1]v\n", NoDelay)
	fmt.Printf("%T\t%[1]v\n", timeout)
	fmt.Printf("%T\t%[1]v\n", time.Minute)
	fmt.Printf("%d\t%d\t%d\t%d\t%d\n", a, b, c, d, e)
}
