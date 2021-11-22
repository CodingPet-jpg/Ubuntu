package geometry

import (
	"math"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

type Point struct {
	X, Y float64
}

type Path []Point

func (p Path) Distance() float64 {
	var sum float64
	for i := 1; i < len(p); i++ {
		sum += p[i-1].Distance(p[i])
	}
	return sum
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// function has it's own namespace
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
