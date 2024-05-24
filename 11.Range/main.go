package main

import "fmt"

func main() {
	subjects := []string{"Math", "Science", "English", "History"}

	for index, subject := range subjects {
		fmt.Printf("Môn học %d: %s\n", index+1, subject)
	}
}
