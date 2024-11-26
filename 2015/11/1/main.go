package main

import (
	"fmt"
)

// Rune manipulation. I learned that you can do rune++ to get the next char in line
func main() {
	line := "cqjxjnds"
	line = increment(line)
	for isInvalid(line) {
		line = increment(line)
	}
	fmt.Println(line)
}

func increment(line string) string {
	runes := []rune(line)
	for i := len(line) - 1; i >= 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i] = runes[i] + 1
			return string(runes)
		}
	}
	return string(runes)
}

func isInvalid(line string) bool {
	return !hasTwoPairs(line) || hasInvalidChar(line) || !hasTripleConsecutiveChar(line)
}

func hasTripleConsecutiveChar(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i]+1 == line[i+1] && line[i]+2 == line[i+2] {
			return true
		}
	}
	return false
}

func hasInvalidChar(line string) bool {
	for _, char := range line {
		if char == 'i' || char == 'o' || char == 'l' {
			return true
		}
	}
	return false
}

func hasTwoPairs(line string) bool {
	pairs := 0
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			pairs++
			i++
		}
		if pairs == 2 {
			return true
		}
	}
	return false
}
