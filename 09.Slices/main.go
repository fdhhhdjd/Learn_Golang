package main

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
)

type Person struct {
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
}

func main() {
	people := []Person{
		{
			Fullname: "Nguyen Van A",
			Age:      30,
		},
		{
			Fullname: "Tran Thi B",
			Age:      28,
		},
		{
			Fullname: "Le Van C",
			Age:      35,
		},
	}

	// Add a new person to the slices
	newPerson := Person{
		Fullname: "Le Van D",
		Age:      29,
	}

	people = slices.Insert(people, len(people), newPerson)

	people = slices.Delete(people, 1, 2)

	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

}
