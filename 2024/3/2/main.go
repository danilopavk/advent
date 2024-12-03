package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	do := true
	for scanner.Scan() {
		input := scanner.Text()
		res, stillDo := parse(input, do)
		do = stillDo
		score += res
	}
	fmt.Println(score)
}

func parse(input string, do bool) (int, bool) {
	if len(input) == 0 {
		return 0, do
	}

	if strings.HasPrefix(input, "don't()") {
		return parse(input[1:], false)
	}

	if strings.HasPrefix(input, "do()") {
		return parse(input[1:], true)
	}

	if !do {
		return parse(input[1:], false)
	}

	if !strings.HasPrefix(input, "mul(") {
		return parse(input[1:], true)
	}

	next := input[4:]
	val := test(next)

	res, do := parse(next, true)
	return val + res, do
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
