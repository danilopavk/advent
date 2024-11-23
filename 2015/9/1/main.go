package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Brute force shortest path algorithm.
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
	shortestPath := math.MaxInt
	for city := range cities {
		delete(cities, city)
		shortestPath = min(findShortestPath(graph, cities, city), shortestPath)
		cities[city] = true
	}

	fmt.Println(shortestPath)
}

func findShortestPath(graph map[string]map[string]int, unvisited map[string]bool, current string) int {
	if len(unvisited) == 0 {
		return 0
	}
	shortestPath := math.MaxInt
	for next := range unvisited {
		delete(unvisited, next)
		path := findShortestPath(graph, unvisited, next)
		pathWithNext := path + graph[current][next]
		shortestPath = min(shortestPath, pathWithNext)
		unvisited[next] = true
	}
	return shortestPath
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
