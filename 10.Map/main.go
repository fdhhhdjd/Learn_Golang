package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	people := map[string]map[string]interface{}{
		"John": {
			"age":    30,
			"gender": "male",
		},
		"Jane": {
			"age":    25,
			"gender": "female",
		},
		"Alice": {
			"age":    28,
			"gender": "female",
		},
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Lỗi khi chuyển đổi thành JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
