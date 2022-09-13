package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type Rectangle struct {
	heigth float64
	width float64
}

func (r Rectangle) area() float64 {
	return r.heigth * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func calArea(f form) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {

	r := Rectangle{10, 15}
	calArea(r)

	c := Circle{10}
	calArea(c)

}