package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	xs := make(map[int][]int)
	allLetters := make(map[int]map[int]rune)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	lines := 0
	columns := 0

	for scanner.Scan() {
		input := scanner.Text()
		columns = len(input)
		parse(input, lineNum, xs, allLetters)
		lineNum++
		lines++
	}
	score := 0
	for lineNum, cols := range xs {
		for _, col := range cols {
			if col == 0 || col == columns-1 {
				continue
			}
			if lineNum == 0 || lineNum == lines-1 {
				continue
			}

			if allLetters[lineNum-1][col-1] == rune('M') &&
				allLetters[lineNum+1][col+1] == rune('S') &&
				allLetters[lineNum-1][col+1] == rune('S') &&
				allLetters[lineNum+1][col-1] == rune('M') {
				score++
			}

			if allLetters[lineNum-1][col-1] == rune('S') &&
				allLetters[lineNum+1][col+1] == rune('M') &&
				allLetters[lineNum-1][col+1] == rune('S') &&
				allLetters[lineNum+1][col-1] == rune('M') {
				score++
			}

			if allLetters[lineNum-1][col-1] == rune('M') &&
				allLetters[lineNum+1][col+1] == rune('S') &&
				allLetters[lineNum-1][col+1] == rune('M') &&
				allLetters[lineNum+1][col-1] == rune('S') {
				score++
			}

			if allLetters[lineNum-1][col-1] == rune('S') &&
				allLetters[lineNum+1][col+1] == rune('M') &&
				allLetters[lineNum-1][col+1] == rune('M') &&
				allLetters[lineNum+1][col-1] == rune('S') {
				score++
			}
		}
	}
	fmt.Println(score)
}

func parse(line string, lineNum int, xs map[int][]int, allLetters map[int]map[int]rune) {
	allLetters[lineNum] = make(map[int]rune)
	xs[lineNum] = []int{}
	for i, c := range line {
		if c == rune('A') {
			xs[lineNum] = append(xs[lineNum], i)
		}
		allLetters[lineNum][i] = c
	}
}
