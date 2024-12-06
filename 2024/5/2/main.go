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

	orders := make(map[int]map[int]bool)
	score := 0

	scanner := bufio.NewScanner(file)

	orderingNow := true
	for scanner.Scan() {
		input := scanner.Text()
		if orderingNow && len(input) == 0 {
			orderingNow = false
			continue
		}

		if orderingNow {
			processOrder(input, orders)
		} else {
			score += processPage(input, orders)
		}

	}

	fmt.Println(score)
}

func processPage(input string, orders map[int]map[int]bool) int {
	parts := strings.Split(input, ",")
	previouses := []int{}
	for _, part := range parts {
		update, _ := strconv.Atoi(part)
		for _, previous := range previouses {
			if orders[previous][update] == true {
				return fix(parts, orders)
			}

		}
		previouses = append(previouses, update)
	}
	return 0
}

func fix(input []string, orders map[int]map[int]bool) int {
	updates := []int{}
	for _, updateStr := range input {
		update, _ := strconv.Atoi(updateStr)
		updates = append(updates, update)
	}

	half := len(updates) / 2
	var best int
	for i := 0; i <= half; i++ {
		best = findBest(updates, orders)
		updates = remove(updates, best)
	}
	return best
}

func findBest(updates []int, orders map[int]map[int]bool) int {
	for i, update := range updates {
		isBest := true
		for j := i + 1; j < len(updates); j++ {
			second := updates[j]
			if orders[update][second] {
				isBest = false
				break
			}
		}
		if isBest {
			return update
		}
	}
	return updates[0]
}

func remove(updates []int, val int) []int {
	newUpdates := []int{}
	for _, update := range updates {
		if update != val {
			newUpdates = append(newUpdates, update)
		}
	}
	return newUpdates
}

func processOrder(input string, orders map[int]map[int]bool) {
	parts := strings.Split(input, "|")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	if _, exists := orders[y]; !exists {
		orders[y] = make(map[int]bool)
	}

	orders[y][x] = true
}
