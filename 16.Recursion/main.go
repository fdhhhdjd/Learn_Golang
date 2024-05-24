package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

func main() {
	fmt.Println("Tổng của các số từ 1 đến 5 là:", sum(5))
}
