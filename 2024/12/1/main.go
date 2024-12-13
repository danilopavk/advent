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
	perimeter := 0
	plant := garden[plot]
	bfs := []point{plot}
	visitedNow := make(map[point]map[point]bool)

	for len(bfs) > 0 {
		nextBfs := make(map[point]bool)
		for _, plot := range bfs {
			visited[plot] = true
			area++

			for _, next := range next(plot) {
				if _, exists := visitedNow[next]; !exists {
					visitedNow[next] = make(map[point]bool)
				}
				if _, exists := visitedNow[plot]; !exists {
					visitedNow[plot] = make(map[point]bool)
				}

				if visitedNow[next][plot] || visitedNow[plot][next] {
					continue
				}

				visitedNow[plot][next] = true
				visitedNow[next][plot] = true

				if next.x >= 0 && next.x < x && next.y >= 0 && next.y < y && garden[next] == plant {
					nextBfs[next] = true
				} else {
					perimeter++
				}
			}
		}
		bfs = []point{}
		for nextPoint, _ := range nextBfs {
			bfs = append(bfs, nextPoint)
		}
	}
	return area * perimeter
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
