package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}
type Line struct {
	Begin, End Point
}
type Path []Point

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}
func main() {
	side := Line{Point{2.34, 5.43}, Point{5.33, 9.99}}
	path := Path{Point{3.34, 434.323}, Point{23.34, 44.323}, Point{25.34, 434.323}, Point{23.34, 434.323}}
	fmt.Println(side.Distance())
	fmt.Println(path.Distance())
}
