package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Không thể chia cho 0")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("Kết quả chia:", result)
	}

	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("Kết quả chia:", result)
	}
}
