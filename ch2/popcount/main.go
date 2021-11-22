package main

import (
	"fmt"
	"unsafe"
)

var pc [256]byte

func main() {
	fmt.Println(PopCountMinus(4095))
}

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PcString(pc [256]byte) {
	for n, v := range pc {
		fmt.Printf("index:%v\tvalue:%v\n", n, v)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(8*0))] +
			pc[byte(x>>(8*1))] +
			pc[byte(x>>(8*2))] +
			pc[byte(x>>(8*3))] +
			pc[byte(x>>(8*4))] +
			pc[byte(x>>(8*5))] +
			pc[byte(x>>(8*6))] +
			pc[byte(x>>(8*7))])
}

func PopCountLoop(x uint64) int {
	var b byte
	var count int = int(unsafe.Sizeof(x) / unsafe.Sizeof(b))
	sum := 0
	for i := 0; i < count; i++ {
		sum += int(pc[byte(x>>(8*i))])
	}
	return count
}

func PopCountObO(x uint64) int {
	count := 0
	for i := 0; i < int(unsafe.Sizeof(x)*8); i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

func PopCountMinus(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((x - (x & (x - 1))))
		x = x >> 1
	}
	return count
}
