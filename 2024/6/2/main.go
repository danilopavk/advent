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

	var guard point
	var direction rune

	obstacles := make(map[point]bool)

	var xLength int
	var yLength int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()	
		xLength = len(input)
		for i, c := range input {
			r := rune(c)
			if r == '#' {
				obstacles[point{i, yLength}] = true
			}
			if r == '>' || r == '<' || r == '^' || r == 'v' {
				direction = r
				guard = point{i, yLength}
			}
		}

		yLength++
	}
	
	circular := 0
	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			newObstacle := point{x, y}
			if guard == newObstacle {
				continue
			}

			if obstacles[newObstacle] {
				continue
			}

			obstacles[newObstacle] = true
			if isCircular(guard, direction, obstacles, xLength, yLength) {
				circular++
			}

			delete(obstacles, newObstacle)
		}
	}
	fmt.Println(circular)
}

func isCircular(guard point, direction rune, obstacles map[point]bool, xLength int, yLength int) bool {
	visited := make(map[point]map[rune]bool)
	for guard.x >= 0 && guard.x < xLength && guard.y >= 0 && guard.y < yLength {
		if _, exists := visited[guard]; !exists {
			visited[guard] = make(map[rune]bool)
		}
		if visited[guard][direction] {
			return true
		}
		visited[guard][direction] = true

		moved := false
		for !moved {
			nextPoint := nextPoint(guard, direction)
			if obstacles[nextPoint] {
				direction = change(direction)
				continue
			} 
			
			guard = nextPoint
			moved = true
		}
 	}
	return false
}

func change(direction rune) rune {
	switch direction {
	case '>': return 'v'
	case 'v': return '<'
	case '<': return '^'
	case '^': return '>'
	}
	return '>'
}

func nextPoint(guard point, direction rune) point {
	switch direction {
	case '>': return point{guard.x + 1, guard.y}
	case 'v': return point{guard.x, guard.y + 1}
	case '<': return point{guard.x - 1, guard.y}
	case '^': return point{guard.x, guard.y - 1}
	}
	return guard
}

type point struct {
	x, y int
}
