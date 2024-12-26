package main

import (
	"bufio"
	"fmt"
	"math"
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

	numericalMap := map[rune]cell{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	}

	directionalMap := map[rune]cell{
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	}

	complexities := 0

	for scanner.Scan() {
		input := scanner.Text()
		directions := expandDirections(input, numericalMap, cell{0, 3})
		secondDirections := []string{}
		for _, direction := range directions {
			expanded := expandDirections(direction, directionalMap, cell{0, 0})
			secondDirections = append(secondDirections, expanded...)
		}
		thirdDirections := []string{}
		for _, direction := range secondDirections {
			expanded := expandDirections(direction, directionalMap, cell{0, 0})
			thirdDirections = append(thirdDirections, expanded...)
		}
		numericalInput, _ := strconv.Atoi(input[:3])
		minLength := math.MaxInt
		for _, direction := range thirdDirections {
			length := len(direction)
			if length < minLength {
				minLength = length
			}
		}
		complexities += numericalInput * minLength
	}

	fmt.Println(complexities)
}

func expandDirections(input string, numericalMap map[rune]cell, avoidingCell cell) []string {
	previous := numericalMap['A']
	directions := []direction{
		{"", move{'A', previous}},
	}
	for _, r := range input {
		goal := numericalMap[r]
		goalReached := false
		for !goalReached {
			newDirections := []direction{}
			for _, presentDirection := range directions {
				lastCell := presentDirection.move.next
				nextMoves := lastCell.towards(goal, avoidingCell)

				for _, nextMove := range nextMoves {
					newDirection := presentDirection.direction + string(nextMove.keypad)
					newDirections = append(newDirections, direction{newDirection, nextMove})
					if nextMove.keypad == 'A' {
						goalReached = true
					}
				}
			}
			directions = make([]direction, len(newDirections))
			copy(directions, newDirections)
		}
	}

	result := []string{}
	for _, direction := range directions {
		result = append(result, direction.direction)
	}
	return result
}

func (c cell) towards(goal cell, avoiding cell) []move {
	if c.x == goal.x && c.y == goal.y {
		return []move{
			{rune('A'), c},
		}
	}
	moves := []move{}
	if goal.x > c.x {
		next := cell{c.x + 1, c.y}
		if next != avoiding {
			moves = append(moves, move{'>', next})
		}
	}

	if goal.x < c.x {
		next := cell{c.x - 1, c.y}
		if next != avoiding {
			moves = append(moves, move{'<', next})
		}
	}

	if goal.y > c.y {
		next := cell{c.x, c.y + 1}
		if next != avoiding {
			moves = append(moves, move{'v', next})
		}
	}

	if goal.y < c.y {
		next := cell{c.x, c.y - 1}
		if next != avoiding {
			moves = append(moves, move{'^', next})
		}
	}

	return moves

}

type direction struct {
	direction string
	move      move
}

type move struct {
	keypad rune
	next   cell
}

type cell struct {
	x, y int
}
