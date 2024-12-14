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

	var machines []prize
	for true {
		scanner.Scan()
		a := parseButton(scanner.Text())
		scanner.Scan()
		b := parseButton(scanner.Text())
		scanner.Scan()
		goal := parseButton(scanner.Text())
		machines = append(machines, prize{a, b, goal})
		if !scanner.Scan() {
			break
		}
		scanner.Text()
	}

	var tokens int
	for _, machine := range machines {
		if minTokens, hasSolution := play(machine); hasSolution {
			tokens += minTokens
		}
	}

	fmt.Println(tokens)
}

func play(machine prize) (int, bool) {
	maxAforX := machine.goal.x / machine.a.x
	maxAforY := machine.goal.y / machine.a.y
	maxA := upTo100(min(maxAforX, maxAforY))

	tokens := math.MaxInt
	hasSolution := false
	for a := 0; a <= maxA; a++ {
		xDistanceForB := machine.goal.x - a*machine.a.x
		if xDistanceForB%machine.b.x != 0 {
			continue
		}
		b := xDistanceForB / machine.b.x
		if b < 0 || b > 100 {
			continue
		}

		if machine.goal.y == a*machine.a.y+b*machine.b.y {
			hasSolution = true
			tokens = min(tokens, 3*a+b)
		}

	}
	return tokens, hasSolution
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func upTo100(a int) int {
	if a > 100 {
		return 100
	}

	return a
}

func parseButton(input string) grid {
	parts := strings.Split(input, " ")
	partsCount := len(parts)

	xPart := parts[partsCount-2]
	x, _ := strconv.Atoi(xPart[2 : len(xPart)-1])

	yPart := parts[partsCount-1]
	y, _ := strconv.Atoi(yPart[2:len(yPart)])

	return grid{x, y}
}

type prize struct {
	a, b, goal grid
}

type grid struct {
	x, y int
}
