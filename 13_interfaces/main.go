package main

import (
	"fmt"
	"math"
)

// Define interface

// Shape interface
type Shape interface {
	area() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rect := Rectangle{width: 10, height: 5}

	fmt.Println("Circle Area:", getArea(circle))
	fmt.Println("Rectangle Area:", getArea(rect))
}
