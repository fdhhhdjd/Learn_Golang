package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//	type Person struct {
//		Fullname string
//		Age      int
//	}
type Person struct {
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
}

func main() {
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])

	// fmt.Println("len:", len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// //? 2X3 (2 row and 3 column)
	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	people := []Person{
		{
			Fullname: "Tai",
			Age:      24,
		},
		{
			Fullname: "Hao",
			Age:      24,
		},
	}
	jsonData, err := json.Marshal(people)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

//Todo 1: emp: [0 0 0 0 0]

//Todo 2: set: [0 0 0 0 100]

//Todo 3: get: 100

//Todo 4: len: 5

//Todo 5: dcl: [1 2 3 4 5]

//Todo 6: 2d:  [[0 1 2] [1 2 3]]
