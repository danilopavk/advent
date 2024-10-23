package main

import (
	"fmt"
	"os"
)

// Since this is the warmup, I went with the simplest possible solution. Of course I could
// read lines separately and then use map reduce to calculate it more efficiently, but hey,
// I'm learning the language, so I'll stick to the basics for now.
func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
		
	result := 0
	for _, c := range content {
		if c == '(' {
			result++
		}
		if c == ')' {
			result--
		}
	}
	fmt.Println(result)
}