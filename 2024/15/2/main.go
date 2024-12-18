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
	boxesLeft := make(map[cell]bool)
	boxesRight := make(map[cell]bool)

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		for i, c := range line {
			x := 2 * i
			currentCell := cell{x, y}
			nextCell := cell{x + 1, y}

			switch c {
			case '#':
				walls[currentCell] = true
				walls[nextCell] = true
			case 'O':
				boxesLeft[currentCell] = true
				boxesRight[nextCell] = true
			case '@':
				robot = currentCell
			}
		}

		y++
	}

	for scanner.Scan() {
		for _, direction := range scanner.Text() {
			robot = move(robot, direction, walls, boxesLeft, boxesRight)
		}
	}

	fmt.Println(score(boxesLeft))
}

func move(robot cell, direction rune, walls, boxesLeft map[cell]bool, boxesRight map[cell]bool) cell {
	nextCells := make(map[cell]bool)
	firstNextCell := next(robot, direction)
	nextCells[firstNextCell] = true

	boxesToMove := []cell{}
	for len(nextCells) > 0 {
		for nextCell, _ := range nextCells {
			delete(nextCells, nextCell)

			if walls[nextCell] {
				return robot
			}

			if boxesLeft[nextCell] {
				rightCell := cell{nextCell.x + 1, nextCell.y}
				boxesToMove = append(boxesToMove, nextCell)
				if direction == '^' || direction == 'v' {
					nextCells[next(nextCell, direction)] = true
					nextCells[next(rightCell, direction)] = true
				}
				if direction == '>' {
					nextCells[next(next(nextCell, direction), direction)] = true
				}
			}

			if boxesRight[nextCell] {
				leftCell := cell{nextCell.x - 1, nextCell.y}
				boxesToMove = append(boxesToMove, leftCell)
				if direction == '^' || direction == 'v' {
					nextCells[next(nextCell, direction)] = true
					nextCells[next(leftCell, direction)] = true
				}
				if direction == '<' {
					nextCells[next(next(nextCell, direction), direction)] = true
				}
			}

		}
	}

	for i := len(boxesToMove) - 1; i >= 0; i-- {
		boxToMove := boxesToMove[i]
		moveTo := next(boxToMove, direction)

		boxesLeft[moveTo] = true
		delete(boxesLeft, boxToMove)

		boxesRight[cell{moveTo.x + 1, moveTo.y}] = true
		delete(boxesRight, cell{boxToMove.x + 1, boxToMove.y})
	}

	return firstNextCell
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
