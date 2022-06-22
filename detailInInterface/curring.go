package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.y-p.y)
}
func main() {
	p := Point{3, 3}
	q := Point{4, 6}
	distanceFrome := q.Distance
	fmt.Println(q.Distance(p))
	fmt.Println(distanceFrome(q))
}
