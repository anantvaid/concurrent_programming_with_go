package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

// Implementation
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Polymorphism
func printArea(s Shape) {
	fmt.Printf("The Area of %T is %.2f\n", s, s.Area())
}

func main() {
	r := Rectangle{Width: 10, Length: 20}
	c := Circle{Radius: 10}

	printArea(c)
	printArea(r)
}
