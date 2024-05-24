package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Width: 3.0, Height: 4.0},
	}

	for _, shape := range shapes {
		fmt.Printf("Diện tích của hình là: %.2f\n", shape.Area())
	}
}
