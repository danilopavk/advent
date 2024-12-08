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

	antinodes := make(map[point]bool)
	nodes := make(map[rune][]point)
	var xLength int
	yLength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		xLength = len(input)

		for x, c := range input {
			r := rune(c)
			if r == '.' {
				continue
			}

			if _, exists := nodes[r]; !exists {
				nodes[r] = []point{}
			}
			nodes[r] = append(nodes[r], point{x, yLength})

		}

		yLength++
	}

	for _, points := range nodes {
		if len(points) == 1 {
			continue
		}

		for _, node := range points {
			for _, node2 := range points {
				if node == node2 {
					continue
				}

				xDiff := node.x - node2.x
				yDiff := node.y - node2.y

				antinode := node
				for antinode.x >= 0 && antinode.x < xLength && antinode.y >= 0 && antinode.y < yLength {
					antinodes[antinode] = true

					antinode = point{antinode.x + xDiff, antinode.y + yDiff}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

type point struct {
	x, y int
}
