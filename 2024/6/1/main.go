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

	var x int
	var y int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		x = len(input)
		for i, c := range input {
			r := rune(c)
			if r == '#' {
				obstacles[point{i, y}] = true
			}
			if r == '>' || r == '<' || r == '^' || r == 'v' {
				direction = r
				guard = point{i, y}
			}
		}

		y++
	}

	visitedCount := 0
	visited := make(map[point]bool)

	for guard.x >= 0 && guard.x < x && guard.y >= 0 && guard.y < y {
		if !visited[guard] {
			visited[guard] = true
			visitedCount++
		}

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
	fmt.Println(visitedCount)
}

func change(direction rune) rune {
	switch direction {
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	case '^':
		return '>'
	}
	return '>'
}

func nextPoint(guard point, direction rune) point {
	switch direction {
	case '>':
		return point{guard.x + 1, guard.y}
	case 'v':
		return point{guard.x, guard.y + 1}
	case '<':
		return point{guard.x - 1, guard.y}
	case '^':
		return point{guard.x, guard.y - 1}
	}
	return guard
}

type point struct {
	x, y int
}
