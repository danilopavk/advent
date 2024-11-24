package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Transforming digits. Simple char to int and back manipulation
func main() {
	line := "1321131112"
	for i := 0; i < 50; i++ {
		line = transform(line)
	}
	fmt.Println(len(line))
}

func transform(input string) string {
	current, _ := strconv.Atoi(string(input[0]))
	var result []int
	counter := 1
	for i := 1; i < len(input); i++ {
		digit, _ := strconv.Atoi(string(input[i]))
		if digit == current && i < len(input)-1 {
			counter++
		} else {
			result = append(result, counter)
			result = append(result, current)
			current = digit
			counter = 1
		}
	}
	result = append(result, counter)
	result = append(result, current)
	var str strings.Builder
	for _, digit := range result {
		str.WriteString(strconv.Itoa(digit))
	}
	return str.String()
}
