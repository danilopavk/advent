package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Where I learn how to parse json file, and how to work with objects without type using interface[]
// and type casting
func main() {
	content, err := os.ReadFile("../../../input")
	if err != nil {
		fmt.Println(err)
	}

	result := 0

	var data []interface{}

	json.Unmarshal(content, &data)

	for _, item := range data {
		result += eval(item)
	}
	fmt.Println(result)
}

func eval(item interface{}) int {
	if num, ok := item.(float64); ok {
		return int(num)
	} else if arr, ok := item.([]interface{}); ok {
		counter := 0
		for _, item := range arr {
			counter += eval(item)
		}
		return counter
	} else if obj, ok := item.(map[string]interface{}); ok {
		counter := 0
		for _, item := range obj {
			counter += eval(item)
		}
		return counter
	} else {
		return 0
	}
}
