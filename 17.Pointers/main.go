package main

import "fmt"

func main() {
	num := 10

	ptr := &num

	fmt.Println("Giá trị của biến:", num)
	fmt.Println("Giá trị được trỏ đến:", *ptr)

	*ptr = 20
	fmt.Println("Giá trị của biến sau khi thay đổi:", num)
}
