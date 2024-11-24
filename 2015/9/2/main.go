package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Small adjustment to make the distance longest.
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	graph := make(map[string]map[string]int)
	cities := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	var edge edge

	// read through and fill graph
	for scanner.Scan() {
		line := scanner.Text()
		edge = parseEdge(line)
		fill(graph, edge)
		cities[edge.from] = true
		cities[edge.to] = true
	}

	// find shortest path
	longestPath := -1
	for city := range cities {
		delete(cities, city)
		longestPath = max(findLongestPath(graph, cities, city), longestPath)
		cities[city] = true
	}

	fmt.Println(longestPath)
}

func findLongestPath(graph map[string]map[string]int, unvisited map[string]bool, current string) int {
	if len(unvisited) == 0 {
		return 0
	}
	longestPath := -1
	for next := range unvisited {
		delete(unvisited, next)
		path := findLongestPath(graph, unvisited, next)
		pathWithNext := path + graph[current][next]
		longestPath = max(longestPath, pathWithNext)
		unvisited[next] = true
	}
	return longestPath
}

func fill(graph map[string]map[string]int, edge edge) {
	_, exists := graph[edge.from]
	if !exists {
		graph[edge.from] = make(map[string]int)
	}
	graph[edge.from][edge.to] = edge.weight

	_, exists = graph[edge.to]
	if !exists {
		graph[edge.to] = make(map[string]int)
	}
	graph[edge.to][edge.from] = edge.weight

}

// edge: Faerun to Norrath = 129
func parseEdge(line string) edge {
	parts := strings.Split(line, " ")
	from := parts[0]
	to := parts[2]
	weightStr := parts[4]
	weight, _ := strconv.Atoi(weightStr)
	return edge{from, to, weight}
}

type edge struct {
	from, to string
	weight   int
}
