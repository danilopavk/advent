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

	garden := make(map[point]rune)
	var x int
	y := 0
	for scanner.Scan() {
		input := scanner.Text()
		x = len(input)
		parse(input, garden, y)
		y++
	}

	price := 0
	visited := make(map[point]bool)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			plot := point{i, j}
			if visited[plot] {
				continue
			}

			price += calculatePrice(garden, visited, plot, x, y)
		}
	}

	fmt.Println(price)
}

func parse(input string, garden map[point]rune, y int) {
	for x, r := range input {
		garden[point{x, y}] = r
	}
}

func calculatePrice(garden map[point]rune, visited map[point]bool, plot point, x int, y int) int {
	area := 0
	perimeter := make(map[fence]bool)
	visitedNow := make(map[point]bool)
	plant := garden[plot]
	bfs := []point{plot}

	for len(bfs) > 0 {
		nextBfs := make(map[point]bool)
		for _, plot := range bfs {
			visited[plot] = true
			visitedNow[plot] = true
			area++

			for i, next := range next(plot) {
				var nextFence fence
				if i%2 == 0 {
					nextFence = fence{next, plot}
				} else {
					nextFence = fence{plot, next}
				}

				if perimeter[nextFence] || nextBfs[next] || visitedNow[next] {
					continue
				}

				if next.x >= 0 && next.x < x && next.y >= 0 && next.y < y && garden[next] == plant {
					nextBfs[next] = true
				} else {
					perimeter[nextFence] = true
				}
			}
		}
		bfs = []point{}
		for nextPoint, _ := range nextBfs {
			bfs = append(bfs, nextPoint)
		}
	}
	return area * sides(perimeter)
}

func sides(perimeter map[fence]bool) int {
	xCrosses := make(map[int]map[int]bool)
	yCrosses := make(map[int]map[int]bool)

	for fence, _ := range perimeter {
		if _, exists := yCrosses[fence.a.y]; !exists {
			yCrosses[fence.a.y] = make(map[int]bool)
		}
		if _, exists := xCrosses[fence.a.x]; !exists {
			xCrosses[fence.a.x] = make(map[int]bool)
		}

		if fence.a.x == fence.b.x {
			xCrosses[fence.a.x][fence.a.y] = true
		} else {
			yCrosses[fence.a.y][fence.a.x] = true
		}
	}

	return disconnected(xCrosses, yCrosses) + disconnected(yCrosses, xCrosses)
}

func disconnected(lines map[int]map[int]bool, opposites map[int]map[int]bool) int {
	sides := 0
	for x, vals := range lines {
		if len(vals) == 0 {
			continue
		}
		for y, _ := range vals {
			// crosses up or down
			if opposites[y][x] || opposites[y+1][x] {
				sides++
			}
		}
	}
	return sides
}

func next(plot point) []point {
	return []point{
		point{plot.x - 1, plot.y},
		point{plot.x + 1, plot.y},
		point{plot.x, plot.y - 1},
		point{plot.x, plot.y + 1}}
}

type point struct {
	x, y int
}

type fence struct {
	a, b point
}
