package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	A float64
	B float64
}

func (t Rectangle) Area() float64 {
	return t.A * t.B
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{3, 4}

	Shape.Area(circle)
	Shape.Area(rectangle)
}
