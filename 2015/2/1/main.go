package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// It's getting fun! This time I learned, in no particular order:
// 1) How to parse file line by line
// 2) Private functions in Golang
// 3) Converting char to int64
// 4) Splitting the string
// 5) Different native data types in golang
// 6) Structs
func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var score int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		dimension := parse(line)
		score += calculateScore(dimension)
	}

	fmt.Println(score)

}

func calculateScore(dimension dimension) int64 {
	x := dimension.w * dimension.h
	y := dimension.w * dimension.l
	z := dimension.l * dimension.h
	min := calculateMin(x, y, z)
	return 2*x + 2*y + 2*z + min

}

func calculateMin(x int64, y int64, z int64) int64 {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}

func parse(line string) dimension {
	parts := strings.Split(line, "x")
	if len(parts) != 3 {
		return dimension{0, 0, 0}
	}
	w, _ := strconv.ParseInt(parts[0], 10, 64)
	l, _ := strconv.ParseInt(parts[1], 10, 64)
	h, _ := strconv.ParseInt(parts[2], 10, 64)
	return dimension{w, l, h}
}

type dimension struct {
	w int64
	l int64
	h int64
}
