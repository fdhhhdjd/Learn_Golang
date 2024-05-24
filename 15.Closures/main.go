package main

import "fmt"

func main() {
	square := func(num int) int {
		return num * num
	}

	fmt.Println("Bình phương của 5:", square(5))
}
