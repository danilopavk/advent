package main

import (
	"bufio"
	"fmt"
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
		newCorrupted := parseCell(scanner.Text())
		corrupted[newCorrupted] = true
		if len(corrupted) < 1024 {
			continue
		}
		if (!hasExit(corrupted)) {
			fmt.Println(newCorrupted)
			break
		}
	}
}

func hasExit(corrupted map[cell]bool) bool {
	start := cell{0, 0}
	end := cell{70, 70}

	visited := make(map[cell]bool)
	visited[start] = true
	stack := make(stack, 0)
	stack = stack.Push(start)
	var current cell

	for !stack.Empty() {
		stack, current = stack.Pop()
		for _, next := range current.next(corrupted, end, visited) {
			if next == end {
				return true
			}
			visited[next] = true
			stack = stack.Push(next)
		}
	}

	return false
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

func (current cell) next(corrupted map[cell]bool, end cell, visited map[cell]bool) []cell {
	next := []cell{}
	candidates := []cell{
		{current.x + 1, current.y},
		{current.x - 1, current.y},
		{current.x, current.y + 1},
		{current.x, current.y - 1},
	}
	for _, candidate := range candidates {
		if corrupted[candidate] {
			continue
		}
		if visited[candidate] {
			continue
		}
		if candidate.x < 0 || candidate.y < 0 || candidate.x > end.x || candidate.y > end.y {
			continue
		}
		next = append(next, candidate)
	}
	return next
}
