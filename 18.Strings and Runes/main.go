package main

import "fmt"

func main() {
	str := "Nguyen Tien Tai!"

	for _, char := range str {
		fmt.Printf("%c ", char)
	}

	runes := []rune(str)
	fmt.Printf("\nChuỗi dưới dạng runes: %v\n", runes)
}
