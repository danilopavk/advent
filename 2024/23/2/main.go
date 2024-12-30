package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	fmt.Println(format(maxCluster(links)))
}

func format(computers []string) string {
	sort.Strings(computers)
	return strings.Join(computers, ",")
}

func maxCluster(links map[string]map[string]bool) []string {
	maxCluster := []string{}
	visited := map[string]bool{}
	for computer := range links {
		allLinks := []string{computer}
		for link := range links[computer] {
			allLinks = append(allLinks, link)
		}
		localMaxCluster := localMaxCluster(allLinks, links, len(maxCluster), visited)
		if len(localMaxCluster) > len(maxCluster) {
			maxCluster = localMaxCluster
		}
	}
	return maxCluster
}

func localMaxCluster(computers []string, links map[string]map[string]bool, best int, visited map[string]bool) []string {
	if len(computers) <= best {
		return []string{}
	}
	formatted := format(computers)
	if visited[formatted] {
		return []string{}
	}

	if isConnected(computers, links) {
		return computers
	}

	bestLocalCluster := []string{}
	for i := 0; i < len(computers); i++ {
		copyComputers := make([]string, len(computers))
		copy(copyComputers, computers)
		leftSide := copyComputers[:i]
		rightSide := copyComputers[i+1:]
		shortenedComputers := append(leftSide, rightSide...)
		localCluster := localMaxCluster(shortenedComputers, links, best, visited)
		if len(localCluster) > len(bestLocalCluster) {
			bestLocalCluster = localCluster
		}
	}
	visited[formatted] = true

	return bestLocalCluster
}

func isConnected(computers []string, links map[string]map[string]bool) bool {
	for _, co1 := range computers {
		for _, co2 := range computers {
			if co1 == co2 {
				continue
			}

			if !links[co1][co2] {
				return false
			}
		}
	}

	return true
}
