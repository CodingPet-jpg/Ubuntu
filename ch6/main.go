package main

import (
	"fmt"

	"image/color"

	"jokoi.com/ch6/colorpoint"

	"jokoi.com/ch6/geometry"
	. "jokoi.com/ch6/net/url"
	"jokoi.com/ch6/package"
)

func main() {
	p0 := geometry.Point{X: 1, Y: 2}
	p1 := geometry.Point{X: 2, Y: 3}
	p2 := geometry.Point{X: 4, Y: 5}
	//fu := p0.Distance
	//fmt.Println(fu(p1))
	perim := geometry.Path{p0, p1, p2}
	fmt.Println(p0.Distance(geometry.Point{X: 3, Y: 4}))
	fmt.Println(perim.Distance())
	fmt.Printf("Before Factor:%v\n", p0)
	p0.ScaleBy(2)
	fmt.Printf("Before Factor:%v\n", p0)

	v := Value{"k1": {"v1", "v2"}, "k2": {"v3", "v4"}}
	fmt.Println(v)
	cp := colorpoint.ColorPoint{
		Point: geometry.Point{X: 1, Y: 2},
		Color: color.RGBA{R: 1, G: 1, B: 1, A: 1},
	}
	fmt.Println(cp)
	fmt.Println(cp.Distance(cp.Point))
	fmt.Println(pack.SSM)
}
