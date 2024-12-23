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

	var start cell
	var end cell
	path := make(map[cell]bool)
	y := 0
	for scanner.Scan() {
		parse(scanner.Text(), &start, &end, path, y)
		y++
	}

	distances := distances(start, end, path)
	distancesInversed := inverse(distances)

	minSaved := 100
	saved := 0
	for from := minSaved; from <= distances[end]; from++ {
		fromCell := distancesInversed[from]
		candidates := []cell{
			{fromCell.x - 2, fromCell.y},
			{fromCell.x - 1, fromCell.y - 1},
			{fromCell.x - 1, fromCell.y + 1},
			{fromCell.x, fromCell.y - 2},
			{fromCell.x, fromCell.y + 2},
			{fromCell.x + 1, fromCell.y - 1},
			{fromCell.x + 1, fromCell.y + 1},
			{fromCell.x + 2, fromCell.y},
		}
		for _, candidate := range candidates {
			if to, exists := distances[candidate]; exists && to+minSaved+2 <= from {
				saved++
			}
		}
	}

	fmt.Println(saved)
}

func inverse(source map[cell]int) map[int]cell {
	target := make(map[int]cell)

	for cell, distance := range source {
		target[distance] = cell
	}

	return target
}

func distances(start cell, end cell, path map[cell]bool) map[cell]int {
	distances := make(map[cell]int)
	distances[start] = 0
	distance := 0
	current := start
	var previous cell
	for current != end {
		next := current.next(path, previous, end)
		distance++
		distances[next] = distance
		previous = current
		current = next
	}
	return distances
}

func parse(line string, start *cell, end *cell, path map[cell]bool, y int) {
	for x, c := range line {
		cell := cell{x, y}
		switch c {
		case '.':
			path[cell] = true
		case 'S':
			*start = cell
		case 'E':
			*end = cell
		}
	}
}

func (current cell) next(path map[cell]bool, previous cell, end cell) cell {
	candidates := []cell{
		{current.x + 1, current.y},
		{current.x - 1, current.y},
		{current.x, current.y + 1},
		{current.x, current.y - 1},
	}

	for _, candidate := range candidates {
		if candidate == previous {
			continue
		}
		if candidate == end {
			return candidate
		}
		if path[candidate] {
			return candidate
		}
	}

	panic("Cannot find next")
}

type cell struct {
	x, y int
}
