package main

import (
	"bufio"
	"fmt"
	"os"
)

// Little bit more of shuffling chars to strings and similar, but not much new here
func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var score int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			score++
		}
	}

	fmt.Println(score)
}

func isNice(line string) bool {
	if !hasThreeVowels(line) {
		return false
	}
	if !hasConsecutive(line) {
		return false
	}
	if hasNaughty(line) {
		return false
	}

	return true
}

func hasNaughty(line string) bool {
	var previous *rune
	for _, char := range line {
		if previous != nil && isNaughty(*previous, char) {
			return true
		}
		previous = &char
	}
	return false
}

var naughty = map[string]struct{}{"ab": {}, "cd": {}, "pq": {}, "xy": {}}

func isNaughty(first rune, second rune) bool {
	together := string(first) + string(second)
	_, exists := naughty[together]
	return exists
}

func hasConsecutive(line string) bool {
	var previous *rune
	for _, char := range line {
		if previous != nil && *previous == char {
			return true
		}
		previous = &char
	}
	return false
}

var vowels = map[rune]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}

func hasThreeVowels(line string) bool {
	vowelsCount := 0
	for _, char := range line {
		if _, exists := vowels[char]; exists {
			vowelsCount++
		}
		if vowelsCount >= 3 {
			return true
		}
	}
	return false
}
