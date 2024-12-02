package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Learned bitwise operations and switch/case!
// This solves task 2, since only input changes
func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wires := map[string]string{}
	values := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		addToWires(wires, line)
	}

	for _, exists := values["a"]; !exists; _, exists = values["a"] {
		for wire, input := range wires {
			handleWire(wire, input, values)
		}
	}

	fmt.Println(values["a"])
}

func handleWire(wire string, input string, values map[string]int) {
	if _, exists := values[wire]; exists {
		return
	}

	parts := strings.Split(input, " ")
	switch len(parts) {
	case 1:
		num, exists := resolve(input, values)
		if exists {
			values[wire] = num
		}
	case 2:
		num, exists := resolve(parts[1], values)
		if exists {
			values[wire] = ^num
		}
	case 3:
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[2])
		leftVal, leftExists := resolve(left, values)
		rightVal, rightExists := resolve(right, values)
		if leftExists && rightExists {
			if num, exists := add(leftVal, rightVal, parts[1]); exists {
				values[wire] = num
			}
		}
	}
}

func add(left int, right int, operation string) (int, bool) {
	switch operation {
	case "RSHIFT":
		return left >> right, true
	case "OR":
		return left | right, true
	case "AND":
		return left & right, true
	case "LSHIFT":
		return left << right, true
	default:
		fmt.Printf("Unknown operation %v", operation)
		return -1, false
	}
}

func resolve(s string, values map[string]int) (int, bool) {
	if isOnlyDigits(s) {
		result, _ := strconv.Atoi(s)
		return result, true
	}

	val, exists := values[s]
	return val, exists
}

func isOnlyDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func addToWires(wires map[string]string, line string) {
	parts := strings.Split(line, "->")
	wires[strings.TrimSpace(parts[1])] = strings.TrimSpace(parts[0])
}
