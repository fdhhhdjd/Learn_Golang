package main

import (
	"fmt"
)

const s string = "Tai Dep Trai"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
}

// Todo 1: Tai Dep Trai

// Todo 2: 6e+11

// Todo 3: 600000000000
