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

	flatten := flatten(shifted)

	checksum := checksum(flatten)

	fmt.Println(checksum)
}

func checksum(shifted []int) int {
	checksum := 0
	for i, num := range shifted {
		if num > 0 {
			checksum += i * num
		}
	}
	return checksum
}

func flatten(shifted [][]int) []int {
	flattened := []int{}
	for _, block := range shifted {
		for i := 0; i < block[0]; i++ {
			flattened = append(flattened, block[1])
		}
	}
	return flattened
}

func shift(numbers [][]int) [][]int {
	blockToMove := len(numbers) - 1
	for blockToMove >= 0 {
		if numbers[blockToMove][1] > -1 {
			if moveTo, exists := findMoveTo(numbers, blockToMove); exists {
				shifted := shiftBetween(numbers, moveTo, blockToMove)
				return shift(shifted)
			}
		}
		blockToMove--
	}
	return numbers
}

func shiftBetween(numbers [][]int, moveTo int, blockToMove int) [][]int {
	moveToBlock := numbers[moveTo]

	shifted := [][]int{}
	for i := 0; i < moveTo; i++ {
		shifted = append(shifted, numbers[i])
	}

	leftover := numbers[moveTo][0] - numbers[blockToMove][0]
	shifted = append(shifted, numbers[blockToMove])
	if leftover > 0 {
		shifted = append(shifted, []int{leftover, moveToBlock[1]})
	}

	for i := moveTo + 1; i < len(numbers); i++ {
		if i != blockToMove {
			shifted = append(shifted, numbers[i])
		} else {
			shiftedBlock := []int{numbers[blockToMove][0], numbers[moveTo][1]}
			shifted = append(shifted, shiftedBlock)
		}
	}

	return shifted
}

func findMoveTo(numbers [][]int, blockToMove int) (int, bool) {
	numOfBlocksToMove := numbers[blockToMove][0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i][1] > -1 {
			continue
		}
		if i >= blockToMove {
			return -1, false
		}

		block := numbers[i]
		if block[0] >= numOfBlocksToMove {
			return i, true
		}
	}

	return -1, false
}

func blocks(numbers []int) [][]int {
	is_file := true
	id := 0
	blocks := [][]int{}
	for _, number := range numbers {
		if is_file {
			blocks = append(blocks, []int{number, id})
		} else if number > 0 {
			blocks = append(blocks, []int{number, -1})
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
