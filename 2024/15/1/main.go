package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var robot cell
	walls := make(map[cell]bool)
	boxes := make(map[cell]bool)

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		for x, c := range scanner.Text() {
			currentCell := cell{x, y}

			switch c {
			case '#':
				walls[currentCell] = true
			case 'O':
				boxes[currentCell] = true
			case '@':
				robot = currentCell
			}
		}

		y++
	}

	for scanner.Scan() {
		for _, direction := range scanner.Text() {
			robot = move(robot, direction, walls, boxes)
		}
	}

	fmt.Println(score(boxes))
}

func move(robot cell, direction rune, walls, boxes map[cell]bool) cell {
	nextCell := next(robot, direction)

	if walls[nextCell] {
		return robot
	}

	if boxes[nextCell] {
		firstBox := nextCell
		lastBox := nextCell
		canMove := false
		for true {
			lastBox = next(lastBox, direction)
			if walls[lastBox] == true {
				break
			}
			if boxes[lastBox] == true {
				continue
			}

			canMove = true
			break
		}

		if canMove {
			boxes[lastBox] = true
			boxes[firstBox] = false
			return firstBox
		}

		return robot
	}

	return nextCell
}

func next(current cell, direction rune) cell {
	switch direction {
	case '>':
		return cell{current.x + 1, current.y}
	case '<':
		return cell{current.x - 1, current.y}
	case 'v':
		return cell{current.x, current.y + 1}
	case '^':
		return cell{current.x, current.y - 1}
	}

	return current
}

func score(boxes map[cell]bool) int {
	score := 0
	for box, isThere := range boxes {
		if isThere {
			score += box.y * 100
			score += box.x
		}

	}
	return score
}

type cell struct {
	x, y int
}
