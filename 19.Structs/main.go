package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	emp := Employee{
		Name:   "John",
		Age:    30,
		Salary: 50000.0,
	}

	fmt.Println("Thông tin nhân viên:")
	fmt.Println("Tên:", emp.Name)
	fmt.Println("Tuổi:", emp.Age)
	fmt.Println("Lương:", emp.Salary)
}
