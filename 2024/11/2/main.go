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
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}
	stonesCount := 0
	for _, values := range stones {
		stonesCount += values
	}

	fmt.Println(stonesCount)
}

func parse(input string) map[int]int {
	stones := make(map[int]int)
	for _, part := range strings.Split(input, " ") {
		stone, _ := strconv.Atoi(part)
		stones[stone] = getOrZero(stones, stone) + 1
	}
	return stones

}

func getOrZero(hashMap map[int]int, key int) int {
	if value, exists := hashMap[key]; exists {
		return value
	}
	return 0
}

func blink(stones map[int]int) map[int]int {
	next := make(map[int]int)
	for stone, stoneCount := range stones {
		if stone == 0 {
			next[1] = getOrZero(next, 1) + stoneCount
			continue
		}

		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%2 == 1 {
			nextVal := stone * 2024
			next[nextVal] = getOrZero(next, nextVal) + stoneCount
			continue
		}

		firstStr := stoneStr[:len(stoneStr)/2]
		first, _ := strconv.Atoi(firstStr)
		next[first] = getOrZero(next, first) + stoneCount

		secondStr := stoneStr[len(stoneStr)/2:]
		second, _ := strconv.Atoi(secondStr)
		next[second] = getOrZero(next, second) + stoneCount
	}
	return next
}
