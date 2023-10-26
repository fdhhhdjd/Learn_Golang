package main

import "fmt"

func main() {
	const a int = 23

	if a == 23 {
		fmt.Println("This is Tai Dev")
	} else {
		fmt.Println("Not must Tai Dev")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is Hao")
	} else if num < 10 {
		fmt.Println(num, "is Tai")
	} else {
		fmt.Println(num, "is Linh")
	}
}

//Todo 1: This is Tai Dev

//Todo 2: 8 is divisible by 4

//Todo 3: is Tai
