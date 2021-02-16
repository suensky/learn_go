package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (r Rect) perim() float64 {
	return 2.0 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2.0 * math.Pi * c.radius
}

func interfaces() {
	fmt.Println("--------- Interface -------------")
	shapes := []Shape{Rect{10, 10}, Circle{10}}
	for _, s := range shapes {
		fmt.Println(s.area())
		fmt.Println(s.perim())
	}
}
