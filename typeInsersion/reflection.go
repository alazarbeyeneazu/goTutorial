package main

import (
	"fmt"
	"math"
	"reflect"
)

type empytInterface interface{}
type Rectangle struct {
	Width, Height int
}
type Circle struct {
	r float64
}
type Shape interface {
	Area()
	Value()
}
type Polygoan interface {
	Sides()
}

func (r Rectangle) Area() {
	fmt.Println(r.Height * r.Width)
}
func (r Rectangle) Value() {
	fmt.Println(r.Height * r.Width * 4)
}
func (c Circle) Area() {
	fmt.Println(math.Pi * c.r)
}
func (c Circle) Value() {
	fmt.Println(c.r * 3 * math.Pi)
}
func (r Rectangle) Sides() {
	fmt.Println("sides ", r.Height*r.Width)
}
func Calculate(inf empytInterface) {
	whatis := reflect.TypeOf(inf)

	fmt.Println(whatis.Name())
	fmt.Println(whatis.NumField())
	fmt.Println(whatis.Kind()) // return the name of type
}
func main() {
	v := Rectangle{Width: 3, Height: 34}
	Calculate(v)
	v2 := Circle{2}
	Calculate(v2)
}
