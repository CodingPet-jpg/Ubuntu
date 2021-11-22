package colorpoint

import (
	"image/color"

	//. "jokoi.com/ch6/fakepoint"
	. "jokoi.com/ch6/geometry"
)

type ColorPoint struct {
	Point
	Color color.RGBA
	//Fakepoint
}

/*
func (cp ColorPoint) Distance(des ColorPoint) float64 {
	return cp.Point.Distance(des.Point)
}
*/
