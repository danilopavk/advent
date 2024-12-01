package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	distance := 0
	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}
	// get all nums
	for scanner.Scan() {
		line := scanner.Text()
		leftItem, rightItem := parse(line)

		left = append(left, leftItem)
		right = append(right, rightItem)
	}

	//sort all
	sort.Ints(left)
	sort.Ints(right)

	// add them up
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(distance)
}

func parse(line string) (int, int) {
	parts := strings.Split(line, " ")
	left, _ := strconv.Atoi(parts[0])
	right, _ := strconv.Atoi(parts[len(parts)-1])
	return left, right
}
