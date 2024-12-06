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
				return 0
			}

		}
		previouses = append(previouses, update)
	}
	half := len(previouses) / 2
	return previouses[half]
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
