package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	links := map[string]map[string]bool{}
	for scanner.Scan() {
		link := scanner.Text()
		parts := strings.Split(link, "-")
		if _, exists := links[parts[0]]; !exists {
			links[parts[0]] = map[string]bool{}
		}
		if _, exists := links[parts[1]]; !exists {
			links[parts[1]] = map[string]bool{}
		}
		links[parts[0]][parts[1]] = true
		links[parts[1]][parts[0]] = true
	}
	fmt.Println(findThreeSets(links))
}

func findThreeSets(links map[string]map[string]bool) int {
	threeSets := 0
	visited := map[string]bool{}
	for first, firstLinks := range links {
		if first[0] != 't' {
			continue
		}
		secondVisited := map[string]bool{}
		for second, _ := range firstLinks {
			if visited[second] {
				continue
			}
			secondLinks := links[second]
			for third, _ := range secondLinks {
				if visited[third] || secondVisited[third] {
					continue
				}
				if firstLinks[third] {
					threeSets++
				}
			}
			secondVisited[second] = true
		}
		visited[first] = true
	}
	return threeSets
}
