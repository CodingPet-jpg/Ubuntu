package fakepoint

import "math"

//import . "jokoi.com"

type Fakepoint struct {
	X, Y float64
}

func (fp Fakepoint) Distance(des Fakepoint) float64 {
	return math.Hypot(fp.X-des.X, fp.Y-des.Y)
}
