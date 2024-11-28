package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Brute force shortest path algorithm on a weighted graph. Added "me"
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	attendies := make(map[string]bool)
	graph := make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line, graph, attendies)
	}

	addMe(attendies, graph)

	maxPath := math.MinInt
	for attendee := range attendies {
		delete(attendies, attendee)
		maxPath = findMax(maxPath, findMaxPath(graph, attendies, attendee, attendee))
		attendies[attendee] = true
	}

	fmt.Println(maxPath)
}

func addMe(attendies map[string]bool, graph map[string]map[string]int) {
	me := "Me"
	attendies[me] = true
	graph[me] = make(map[string]int)
	for attendee := range attendies {
		graph[me][attendee] = 0
		graph[attendee][me] = 0
	}
}

func findMaxPath(graph map[string]map[string]int, attendies map[string]bool, current string, first string) int {
	if len(attendies) == 0 {
		return graph[current][first] + graph[first][current]
	}
	maxPath := math.MinInt
	for next := range attendies {
		delete(attendies, next)
		nextMaxPath := graph[current][next] + graph[next][current] + findMaxPath(graph, attendies, next, first)
		maxPath = findMax(maxPath, nextMaxPath)
		attendies[next] = true
	}
	return maxPath
}

func findMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseLine(line string, graph map[string]map[string]int, attendies map[string]bool) {
	parts := strings.Split(line, " ")
	left := parts[0]
	right := parts[len(parts)-1]
	right = right[:len(right)-1]
	attendies[left] = true
	attendies[right] = true
	weight, _ := strconv.Atoi(parts[3])
	if parts[2] == "lose" {
		weight *= -1
	}
	if graph[left] == nil {
		graph[left] = make(map[string]int)
	}
	graph[left][right] = weight
}
