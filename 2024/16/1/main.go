package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var start cell
	var end cell
	walls := make(map[cell]bool)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			cell := cell{x, y}
			switch c {
			case '#':
				walls[cell] = true
			case 'S':
				start = cell
			case 'E':
				end = cell
			}
		}
		y++
	}

	direction := '>'
	allDirections := []rune{'<', '>', '^', 'v'}
	startWithDirection := cellDirection{start, direction}

	scores := make(map[cellDirection]int)
	scores[startWithDirection] = 0

	stack := make(stack, 0)
	stack = stack.Push(startWithDirection)
	var current cellDirection

	for !stack.Empty() {
		stack, current = stack.Pop()
		currentScore := scores[current]

		for _, nextDirection := range allDirections {
			next := current.cell.next(nextDirection)
			if walls[next] {
				continue
			}

			var nextScore int
			if current.direction == nextDirection {
				nextScore = currentScore + 1
			} else {
				nextScore = currentScore + 1001
			}

			nextWithDirection := cellDirection{next, nextDirection}
			var valueAssigned = false
			if _, exists := scores[nextWithDirection]; !exists {
				scores[nextWithDirection] = nextScore
				valueAssigned = true
			}
			if nextScore < scores[nextWithDirection] {
				scores[nextWithDirection] = nextScore
				valueAssigned = true
			}

			if next != end && valueAssigned {
				stack = stack.Push(nextWithDirection)
			}
		}
	}

	minEnd := math.MaxInt
	for _, direction := range allDirections {
		if val, exists := scores[cellDirection{end, direction}]; exists && val < minEnd {
			minEnd = val
		}
	}
	fmt.Println(minEnd)
}

func (c cell) next(direction rune) cell {
	switch direction {
	case '>':
		return cell{c.x + 1, c.y}
	case '<':
		return cell{c.x - 1, c.y}
	case '^':
		return cell{c.x, c.y - 1}
	case 'v':
		return cell{c.x, c.y + 1}
	default:
		return c
	}
}

type cell struct {
	x, y int
}

type cellDirection struct {
	cell      cell
	direction rune
}

type stack []cellDirection

func (s stack) Push(v cellDirection) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, cellDirection) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Empty() bool {
	return len(s) == 0
}
