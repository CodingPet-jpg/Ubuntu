package main

import (
	"fmt"
)

type point struct {
	X, Y int
}

type Circle struct {
	point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w0 := Wheel{Circle{point{1, 2}, 3}, 4}
	w1 := Wheel{
		Circle: Circle{
			point: point{
				X: 1,
				Y: 2,
			},
			Radius: 3,
		},
		Spokes: 4,
	}
	fmt.Printf("%#v\n", w0)
	fmt.Printf("%#v\n", w1)
}
