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

	startingMembers := make(map[cell]bool)
	startingMembers[start] = true
	startingPath := path{start, startingMembers, direction, 0}

	stack := make(stack, 0)
	stack = stack.Push(startingPath)
	var current path

	bestEndScore := math.MaxInt
	bestPaths := []path{}

	// store best scores for each cell to spead up the developemnt
	bestScores := make(map[cellWithDirection]int)
	bestScores[cellWithDirection{start, direction}] = 0

	for !stack.Empty() {
		stack, current = stack.Pop()
		currentScore := current.score

		var previousNext cell
		useCurrentCellsMap := true
		for _, nextDirection := range allDirections {
			// opposite means coming back to same
			if isOpposite(nextDirection, current.direction) {
				continue
			}

			next := current.last.next(nextDirection)
			if walls[next] {
				continue
			}

			if current.cells[next] {
				continue
			}

			var nextScore int
			if current.direction == nextDirection {
				nextScore = currentScore + 1
			} else {
				nextScore = currentScore + 1001
			}
			if nextScore > bestEndScore {
				continue
			}

			// we've already been here and with a better score
			nextWithDirection := cellWithDirection{next, nextDirection}
			if val, exists := bestScores[nextWithDirection]; exists {
				if val < nextScore {
					continue
				}
			}

			if next == end && nextScore > bestEndScore {
				continue
			}

			var nextMap map[cell]bool
			if useCurrentCellsMap {
				// make the current map same as for the next item
				nextMap = current.cells
				nextMap[next] = true
				previousNext = next
				useCurrentCellsMap = false
			} else {
				// if we already used current map, we need to create a new one and copy its content
				nextMap = make(map[cell]bool)
				for cell, _ := range current.cells {
					nextMap[cell] = true
				}
				nextMap[next] = true
				delete(nextMap, previousNext)
			}

			nextPath := path{next, nextMap, nextDirection, nextScore}

			if next != end {
				bestScores[nextWithDirection] = nextScore
				stack = stack.Push(nextPath)
			}

			if next == end {
				if bestEndScore > nextScore {
					bestEndScore = nextScore
					bestPaths = []path{nextPath}
				} else if bestEndScore == nextScore {
					bestPaths = append(bestPaths, nextPath)
				}
			}
		}
	}

	bestCells := make(map[cell]bool)
	for _, path := range bestPaths {
		for cell, _ := range path.cells {
			bestCells[cell] = true
		}
	}

	fmt.Println(len(bestCells))
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

func isOpposite(a rune, b rune) bool {
	switch a {
	case '>':
		return b == '<'
	case '<':
		return b == '>'
	case '^':
		return b == 'v'
	case 'v':
		return b == '^'
	default:
		return false
	}
}

type cell struct {
	x, y int
}

type cellWithDirection struct {
	cell      cell
	direction rune
}

type path struct {
	last      cell
	cells     map[cell]bool
	direction rune
	score     int
}

type stack []path

func (s stack) Push(v path) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, path) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Empty() bool {
	return len(s) == 0
}
