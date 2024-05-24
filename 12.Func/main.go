package main

import "fmt"

func rectangleArea(length, width float64) float64 {
	return length * width
}

func main() {
	area := rectangleArea(5.5, 3.0)
	fmt.Println("Diện tích hình chữ nhật là:", area)
}
