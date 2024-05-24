package main

import (
	"fmt"
	"time"
)

func fetchData(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Data fetched successfully"
}

func main() {
	// TH1
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// TH2
	ch := make(chan string)
	go fetchData(ch)
	result := <-ch
	fmt.Println(result)

	// TH2
	resultChan := demo()

	resultTH2 := <-resultChan
	demo1(resultTH2)
}

func demo() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	return ch
}

func demo1(number int) {
	fmt.Println(number)
}
