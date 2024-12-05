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

			if lineNum > 2 && allLetters[lineNum-1][col] == rune('M') &&
				allLetters[lineNum-2][col] == rune('A') && allLetters[lineNum-3][col] == rune('S') {
				score++
			}

			if lineNum > 2 && col < columns-3 && allLetters[lineNum-1][col+1] == rune('M') &&
				allLetters[lineNum-2][col+2] == rune('A') && allLetters[lineNum-3][col+3] == rune('S') {
				score++
			}

			if col < columns-3 && allLetters[lineNum][col+1] == rune('M') &&
				allLetters[lineNum][col+2] == rune('A') && allLetters[lineNum][col+3] == rune('S') {
				score++
			}

			if lineNum < lines-3 && col < columns-3 && allLetters[lineNum+1][col+1] == rune('M') &&
				allLetters[lineNum+2][col+2] == rune('A') && allLetters[lineNum+3][col+3] == rune('S') {
				score++
			}

			if lineNum < lines-3 && allLetters[lineNum+1][col] == rune('M') &&
				allLetters[lineNum+2][col] == rune('A') && allLetters[lineNum+3][col] == rune('S') {
				score++
			}

			if lineNum < lines-3 && col > 2 && allLetters[lineNum+1][col-1] == rune('M') &&
				allLetters[lineNum+2][col-2] == rune('A') && allLetters[lineNum+3][col-3] == rune('S') {
				score++
			}

			if col > 2 && allLetters[lineNum][col-1] == rune('M') &&
				allLetters[lineNum][col-2] == rune('A') && allLetters[lineNum][col-3] == rune('S') {
				score++
			}

			if lineNum > 2 && col > 2 && allLetters[lineNum-1][col-1] == rune('M') &&
				allLetters[lineNum-2][col-2] == rune('A') && allLetters[lineNum-3][col-3] == rune('S') {
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
		if c == rune('X') {
			xs[lineNum] = append(xs[lineNum], i)
		}
		allLetters[lineNum][i] = c
	}
}
