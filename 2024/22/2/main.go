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

	maxSell := 0

	sequences := map[sequence]int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		newSequences := calculateSequences(num)
		for sequence, sell := range newSequences {
			addedSell := sequences[sequence] + sell
			sequences[sequence] = addedSell
			if addedSell > maxSell {
				maxSell = addedSell
			}
		}
	}

	fmt.Println(maxSell)
}

func calculateSequences(num int) map[sequence]int {
	sequenceChanges := map[sequence]int{}
	lastPrice := -1
	sequences := []int{}
	for i := 0; i < 2000; i++ {
		num = generateNextSequence(num)
		numString := strconv.Itoa(num)
		price, _ := strconv.Atoi(string(numString[len(numString)-1]))
		if lastPrice > -1 {
			sequences = append(sequences, price-lastPrice)
			sequenceLength := len(sequences)
			if sequenceLength >= 4 {
				newSequence := sequence{
					sequences[sequenceLength-4],
					sequences[sequenceLength-3],
					sequences[sequenceLength-2],
					sequences[sequenceLength-1],
				}
				if sequenceChanges[newSequence] == 0 {
					sequenceChanges[newSequence] = price
				}
			}
		}
		lastPrice = price

	}
	return sequenceChanges
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

type sequence struct {
	a, b, c, d int
}
