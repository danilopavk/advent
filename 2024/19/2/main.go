package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	towels, maxTowelLength := parseTowels(scanner.Text())

	// emptyLine
	scanner.Scan()
	scanner.Text()

	possible := 0
	for scanner.Scan() {
		possible += buildDesign(scanner.Text(), towels, maxTowelLength)
	}

	fmt.Println(possible)
}

func buildDesign(design string, patterns map[string]bool, maxPatternLength int) int {
	possibilities := make(map[int]int)
	possibilities[len(design)] = 1
	for start := len(design) - 1; start >= 0; start-- {
		for end := start + maxPatternLength; end > start; end-- {
			if end > len(design) {
				continue
			}
			pattern := design[start:end]
			if patterns[pattern] {
				possibilities[start] = possibilities[start] + possibilities[end]
			}
		}
	}
	return possibilities[0]
}

func parseTowels(input string) (map[string]bool, int) {
	towels := make(map[string]bool)
	maxTowelLength := 0

	parts := strings.Split(input, ",")
	for _, part := range parts {
		towel := strings.Trim(part, " ")

		towels[towel] = true

		towelLength := len(towel)
		if towelLength > maxTowelLength {
			maxTowelLength = towelLength
		}
	}

	return towels, maxTowelLength
}
