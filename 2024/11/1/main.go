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
	scanner := bufio.NewScanner(file)

	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}
	stones := parse(input)
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	fmt.Println(len(stones))
}

func parse(input string) []int {
	stones := []int{}
	for _, part := range strings.Split(input, " ") {
		stone, _ := strconv.Atoi(part)
		stones = append(stones, stone)
	}
	return stones

}

func blink(stones []int) []int {
	next := []int{}
	for _, stone := range stones {
		if stone == 0 {
			next = append(next, 1)
			continue
		}

		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%2 == 1 {
			next = append(next, stone*2024)
			continue
		}

		firstStr := stoneStr[:len(stoneStr)/2]
		first, _ := strconv.Atoi(firstStr)
		next = append(next, first)

		secondStr := stoneStr[len(stoneStr)/2:]
		second, _ := strconv.Atoi(secondStr)
		next = append(next, second)
	}
	return next
}
