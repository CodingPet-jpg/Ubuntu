package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	var format string = "%08b\n"
	fmt.Printf(format, x)    //00100010
	fmt.Printf(format, y)    //00000110
	fmt.Printf(format, x&y)  //00000010
	fmt.Printf(format, x|y)  //00100110
	fmt.Printf(format, x^y)  //00100100
	fmt.Printf(format, x&^y) //00100000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // 1,5
		}
	}
	fmt.Printf(format, x<<1) //01000100
	fmt.Printf(format, x>>1) //00010001
}
