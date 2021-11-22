package main

import (
	"fmt"
)

func main() {
	s := "プログラム"
	fmt.Printf("% 3x\n", s) //在%和x间插入空格则输出每两个数位插入空格
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string('\u3451'))
}
