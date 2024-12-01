package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	similarities := 0
	left := []int{}
	right := make(map[int]int)

	// get all nums
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		leftItem, rightItem := parse(line)

		// just append the left
		left = append(left, leftItem)

		// for right, count the appearances
		if rightVal, exists := right[rightItem]; exists {
			right[rightItem] = rightVal + 1
		} else {
			right[rightItem] = 1
		}
	}

	// calcualte the similarities
	for i := 0; i < len(left); i++ {
		leftItem := left[i]
		if rightCounter, exists := right[leftItem]; exists {
			similarities += leftItem * rightCounter
		}
	}

	fmt.Println(similarities)
}

func parse(line string) (int, int) {
	parts := strings.Split(line, " ")
	left, _ := strconv.Atoi(parts[0])
	right, _ := strconv.Atoi(parts[len(parts)-1])
	return left, right
}
