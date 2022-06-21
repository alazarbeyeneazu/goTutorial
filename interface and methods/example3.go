package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	prime() float64
}
type Rect struct {
	width, height float64
}
type Circle struct {
	Radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}
func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rect) prime() float64 {
	return 2*r.height + 2*r.width
}
func (c Circle) prime() float64 {
	return math.Pi*2 + c.Radius*2
}
func getGeometry(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}
func main() {
	var r = Rect{width: 2, height: 3}
	var c = Circle{Radius: 5}
	fmt.Println(c.area())

	fmt.Printf("%#v\n", r.area())
}
