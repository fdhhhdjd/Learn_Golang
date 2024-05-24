package main

import "fmt"

func findMinMax(numbers []int) (min int, max int) {
	min, max = numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	numbers := []int{3, 1, 7, 5, 2}

	min, max := findMinMax(numbers)
	fmt.Printf("Số nhỏ nhất: %d, Số lớn nhất: %d\n", min, max)
}
