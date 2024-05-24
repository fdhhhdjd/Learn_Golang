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
}
