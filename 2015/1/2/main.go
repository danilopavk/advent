package main

import (
	"fmt"
	"os"
)

// Well, this took just a small amount of modifications, but here we are.
// I learned that in golang, break is still break
func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
		
	result := 0
	for i, c := range content {
		if c == '(' {
			result++
		}
		if c == ')' {
			result--
		}
		if result < 0 {
			fmt.Println(i + 1)
			break
		}
	}
}