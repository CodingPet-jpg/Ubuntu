package main

import (
	"fmt"
	"math"
)

const Pi64 float64 = math.Pi
const StrNotype = "Hello"
const Str string = "Hello"

func main() {
	var x float32 = math.Pi
	var y float32 = float32(math.Pi)
	fmt.Println(x, y)
	fmt.Printf("%T\n", '\000')
}
