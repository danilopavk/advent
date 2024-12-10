package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	geomap := make(map[point]int)
	trailheads := make(map[point]bool)

	scanner := bufio.NewScanner(file)
	var x int
	y := 0
	for scanner.Scan() {
		input := scanner.Text()
		x = len(input)
		parse(input, geomap, trailheads, y)
		y++
	}
	score := 0
	for trailhead, _ := range trailheads {
		score += traverse(trailhead, geomap, x, y)
	}

	fmt.Println(score)
}

func traverse(trailhead point, geomap map[point]int, x int, y int) int {
	candidates := make(map[point]int)
	candidates[trailhead] = 1
	for i := 1; i <= 9; i++ {
		nextCandidates := make(map[point]int)
		for here, hereVal := range candidates {
			next := next(here, x, y)
			for _, nextCandidate := range next {
				if geomap[nextCandidate] == i {
					nextCandidateValue, exists := nextCandidates[nextCandidate]
					if !exists {
						nextCandidateValue = 0
					}
					nextCandidates[nextCandidate] = nextCandidateValue + hereVal
				}
			}
		}
		if len(nextCandidates) == 0 {
			return 0
		} else {
			candidates = nextCandidates
		}
	}

	score := 0
	for _, val := range candidates {
		score += val
	}
	return score
}

func next(current point, x int, y int) []point {
	next := []point{}
	if current.x > 0 {
		next = append(next, point{current.x - 1, current.y})
	}
	if current.x < x-1 {
		next = append(next, point{current.x + 1, current.y})
	}
	if current.y > 0 {
		next = append(next, point{current.x, current.y - 1})
	}
	if current.y < y-1 {
		next = append(next, point{current.x, current.y + 1})
	}
	return next
}

func parse(input string, geomap map[point]int, trailheads map[point]bool, y int) {
	for x, c := range input {
		n, _ := strconv.Atoi(string(c))
		here := point{x, y}
		geomap[here] = n
		if n == 0 {
			trailheads[here] = true
		}
	}
}

type point struct {
	x, y int
}
