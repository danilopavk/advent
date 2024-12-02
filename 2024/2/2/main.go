package main

import (
	"bufio"
	"fmt"
	"math"
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
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, parse(line))
	}

	safeCounter := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCounter++
		} else {
			for i := 0; i < len(report); i++ {
				copy := append([]int{}, report...)
				copy = remove(copy, i)
				if isSafe(copy) {
					safeCounter++
					break
				}
			}
		}
	}

	fmt.Println(safeCounter)
}

func isSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	}
	isAscending := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		if isAscending && report[i] <= report[i-1] {
			return false
		}
		if !isAscending && report[i] >= report[i-1] {
			return false
		}
		diff := int(math.Abs(float64(report[i] - report[i-1])))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func remove(slice []int, s int) []int {
	copy(slice[s:], slice[s+1:])
	return slice[:len(slice)-1]
}

func parse(line string) []int {
	parts := strings.Split(line, " ")
	var levels []int
	for i := 0; i < len(parts); i++ {
		level, _ := strconv.Atoi(parts[i])
		levels = append(levels, level)
	}
	return levels
}
