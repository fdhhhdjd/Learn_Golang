package main

import (
	"fmt"
)

type Status int

const (
	Pending Status = iota
	Processing
	Completed
)

func main() {
	status := Processing

	switch status {
	case Pending:
		fmt.Println("Công việc đang chờ xử lý.")
	case Processing:
		fmt.Println("Công việc đang được xử lý.")
	case Completed:
		fmt.Println("Công việc đã hoàn thành.")
	}
}
