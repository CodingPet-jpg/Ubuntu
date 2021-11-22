package main

import (
	"fmt"
	"math"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o", o)
	fmt.Println()
	x := 0xdeadbeef
	fmt.Printf("%d %[1]x %#[1]x %#[1]X", x)
	fmt.Println()
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	s := "%d %[1]c %[1]q"
	fmt.Printf(s, ascii)
	fmt.Println()
	fmt.Printf(s, unicode)
	fmt.Println()
	fmt.Printf(s, newline)
	fmt.Println()

	var f float32 = 1 << 24
	var mf float32 = math.MaxFloat32
	fmt.Printf("%g %g\n", f, mf)

	for i := 0; i < 9; i++ {
		fmt.Printf("x = %d e^x = %8.3f\n", i, math.Exp(float64(i)))
	}

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan, math.IsNaN(nan))
}
func compute() (value float64, ok bool) {
	if false {
		return 0, false
	}
	return 1.9, true
}
