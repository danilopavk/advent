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

	var start cell
	var end cell
	path := make(map[cell]bool)
	y := 0
	for scanner.Scan() {
		input := scanner.Text()
		parse(input, &start, &end, path, y)
		y++
	}

	distances := distances(start, end, path)
	distancesInversed := inverse(distances)

	minSaved := 100
	saved := 0
	for from := minSaved; from <= distances[end]; from++ {
		for to := 0; to <= from - minSaved; to++ {
			fromCell := distancesInversed[from]
			toCell := distancesInversed[to]

			xDiff := abs(fromCell.x, toCell.x)
			yDiff := abs(fromCell.y, toCell.y)
			distance := xDiff + yDiff
			if distance > 20  {
				continue
			}

			newPath := distances[end] - from + to + distance
			if distances[end] - newPath >= minSaved {
				saved++
			} 
		}
	}

	fmt.Println(saved)
}

func abs(a int, b int) int {
	return int(math.Abs(float64(a - b)))
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

func (path cheat) neighbours(x int, y int) []cell {
	last := path.last
	neighbours :=  []cell{}
	candidates := []cell{ 	
		{last.x + 1, last.y},
		{last.x - 1, last.y},
		{last.x, last.y + 1},
		{last.x, last.y - 1},
	}
	for _, candidate := range candidates {
		if path.visited[candidate] {
			continue
		}

		if candidate.x < 0 || candidate.x >= x || candidate.y < 0 || candidate.y >= y {
			continue
		}

		neighbours = append(neighbours, candidate)
	}

	return neighbours
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

func asString(first cell, last cell) string {
	var b strings.Builder
	b.WriteString(strconv.Itoa(first.x))
	b.WriteString(",")
	b.WriteString(strconv.Itoa(first.y))
	b.WriteString(";")
	b.WriteString(strconv.Itoa(last.x))
	b.WriteString(",")
	b.WriteString(strconv.Itoa(last.y))
	return b.String()

}

type cheat struct {
	first cell
	last cell
	visited map[cell]bool
}

type cell struct {
	x, y int
}
