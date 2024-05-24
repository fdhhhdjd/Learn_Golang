package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	c := Circle{
		Radius: 5.0,
	}
	fmt.Println("Diện tích của hình tròn:", c.Area())

}
