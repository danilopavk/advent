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

	corrupted := make(map[cell]bool)
	for scanner.Scan() {
		corrupted[parseCell(scanner.Text())] = true
		if len(corrupted) == 1024 {
			break
		}
	}
	start := cell{0, 0}
	end := cell{70, 70}
	minPath := math.MaxInt

	visited := make(map[cell]int)
	visited[start] = 0
	stack := make(stack, 0)
	stack = stack.Push(start)
	var current cell

	for !stack.Empty() {
		stack, current = stack.Pop()
		currentLength := visited[current]
		if currentLength >= minPath {
			continue
		}
		for _, next := range current.next(corrupted, end, visited) {
			nextLength := currentLength + 1
			if next == end {
				minPath = nextLength
				break
			}
			visited[next] = nextLength
			stack = stack.Push(next)
		}
	}

	fmt.Println(minPath)
}

func parseCell(line string) cell {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return cell{x, y}
}

type cell struct {
	x, y int
}

type stack []cell

func (s stack) Push(v cell) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, cell) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Empty() bool {
	return len(s) == 0
}

func (current cell) next(corrupted map[cell]bool, end cell, visited map[cell]int) []cell {
	next := []cell{}
	candidates := []cell{
		cell{current.x + 1, current.y},
		cell{current.x - 1, current.y},
		cell{current.x, current.y + 1},
		cell{current.x, current.y - 1},
	}
	for _, candidate := range candidates {
		if corrupted[candidate] {
			continue
		}
		if candidate.x < 0 || candidate.y < 0 || candidate.x > end.x || candidate.y > end.y {
			continue
		}
		currentPathCount := visited[current]
		if candidateCount, exists := visited[candidate]; exists {
			if candidateCount <= currentPathCount+1 {
				continue
			}
		}
		next = append(next, candidate)
	}
	return next
}
