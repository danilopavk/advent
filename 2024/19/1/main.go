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
		if buildDesign(scanner.Text(), towels, maxTowelLength) {
			possible++
		}
	}

	fmt.Println(possible)
}

func buildDesign(design string, patterns map[string]bool, maxPatternLength int) bool {
	pointers := []int{0}
	visited := make(map[int]bool)
	for len(pointers) > 0 {
		newPointers := []int{}
		for _, pointer := range pointers {
			if pointer == len(design) {
				return true
			}
			if visited[pointer] {
				continue
			}
			visited[pointer] = true

			for i := 1; i <= maxPatternLength; i++ {
				if pointer+i > len(design) {
					break
				}
				pattern := design[pointer : pointer+i]
				if patterns[pattern] {
					newPointers = append(newPointers, pointer+i)
				}
			}
		}
		pointers = newPointers
	}
	return false
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
