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

	var machines []prize
	for true {
		scanner.Scan()
		a := parseButton(scanner.Text())
		scanner.Scan()
		b := parseButton(scanner.Text())
		scanner.Scan()
		goal := parseButton(scanner.Text())
		increase := 10000000000000
		increasedGoal := grid{goal.x + increase, goal.y + increase}
		machines = append(machines, prize{a, b, increasedGoal})
		if !scanner.Scan() {
			break
		}
		scanner.Text()
	}

	var tokens int
	for _, machine := range machines {
		tokens += play(machine)
	}

	fmt.Println(tokens)
}

func play(machine prize) int {
	upper := machine.goal.y*machine.a.x - machine.goal.x*machine.a.y
	lower := machine.a.x*machine.b.y - machine.a.y*machine.b.x

	if lower == 0 {
		return 0
	}

	if upper%lower != 0 {
		return 0
	}
	b := upper / lower
	if b < 0 {
		return 0
	}
	leftForA := machine.goal.x - machine.b.x*b
	if leftForA%machine.a.x != 0 {
		return 0
	}
	a := leftForA / machine.a.x
	if a*machine.a.y+b*machine.b.y != machine.goal.y {
		return 0
	}
	return a*3 + b
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
