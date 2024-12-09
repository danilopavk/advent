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
	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}
	numbers := parse(input)

	blocks := blocks(numbers)

	shifted := shift(blocks)

	checksum := checksum(shifted)

	fmt.Println(checksum)
}

func checksum(shifted []int) int {
	checksum := 0
	for i, num := range shifted {
		checksum += i * num
	}
	return checksum
}

func shift(numbers []int) []int {
	i := 0
	j := len(numbers) - 1
	shifted := []int{}
	for i <= j {
		if numbers[i] >= 0 {
			shifted = append(shifted, numbers[i])
			i++
		} else if numbers[j] >= 0 {
			shifted = append(shifted, numbers[j])
			i++
			j--
		} else {
			j--
		}
	}
	return shifted
}

func blocks(numbers []int) []int {
	is_file := true
	id := 0
	blocks := []int{}
	for _, number := range numbers {
		for i := 0; i < number; i++ {
			if is_file {
				blocks = append(blocks, id)
			} else {
				blocks = append(blocks, -1)
			}
		}
		if is_file {
			id++
		}
		is_file = !is_file
	}
	return blocks
}

func parse(input string) []int {
	numbers := []int{}
	for _, c := range input {
		number, _ := strconv.Atoi(string(c))
		numbers = append(numbers, number)
	}
	return numbers
}
