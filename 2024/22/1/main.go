package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		score += calculatePrice(num)
	}

	fmt.Println(score)
}

func calculatePrice(num int) int {
	for i := 0; i < 2000; i++ {
		num = generateNextSequence(num)
	}
	return num
}

func generateNextSequence(num int) int {
	num = pruneAndMix(num*64, num)
	num = pruneAndMix(num/32, num)
	return pruneAndMix(num*2048, num)
}

func pruneAndMix(num int, mixIn int) int {
	mixed := mixIn ^ num
	return mixed % 16777216
}
