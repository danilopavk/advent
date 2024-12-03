package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// poor memory management because it copies the string a lot, but works for my
// input string
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		input := scanner.Text()
		score += parse(input)
	}
	fmt.Println(score)
}

func parse(input string) int {
	if len(input) == 0 {
		return 0
	}

	if !strings.HasPrefix(input, "mul(") {
		return parse(input[1:])
	}

	next := input[4:]
	val := test(next)

	return val + parse(next[1:])
}

func test(input string) int {
	a, chars, exists := getNum(input)
	if !exists {
		return 0
	}
	input = input[chars:]
	if !strings.HasPrefix(input, ",") {
		return 0
	}
	input = input[1:]
	b, chars, exists := getNum(input)
	if !exists {
		return 0
	}
	input = input[chars:]
	if !strings.HasPrefix(input, ")") {
		return 0
	}
	return a * b
}

func getNum(input string) (int, int, bool) {
	if !unicode.IsDigit(rune(input[0])) {
		return 0, 0, false
	}

	if !unicode.IsDigit(rune(input[1])) {
		num, _ := strconv.Atoi(input[:1])
		return num, 1, true
	}

	if !unicode.IsDigit(rune(input[2])) {
		num, _ := strconv.Atoi(input[:2])
		return num, 2, true
	}

	num, _ := strconv.Atoi(input[:3])
	return num, 3, true
}
