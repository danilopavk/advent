package main

import (
	"bufio"
	"fmt"
	"os"
)

// The algorithm was a bit more fun!
func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var score int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			score++
		}
	}

	fmt.Println(score)
}

func isNice(line string) bool {
	return hasSkipPair(line) && hasDoublePair(line)
}

func hasSkipPair(line string) bool {
	if len(line) < 3 {
		return false
	}
	secondLast := line[0]
	last := line[1]
	for i := 2; i < len(line); i++ {
		next := line[i]
		if next == secondLast {
			return true
		}
		secondLast = last
		last = next
	}
	return false
}

func hasDoublePair(line string) bool {
	fmt.Println("Checking " + line)
	if len(line) < 4 {
		return false
	}
	lastChar := line[1]
	lastSeq := string(line[0]) + string(lastChar)
	allPairs := map[string]struct{}{lastSeq: {}}
	lastSeqIsOnlySeq := true
	for i := 2; i < len(line); i++ {
		nextChar := line[i]
		nextSeq := string(lastChar) + string(nextChar)
		_, isDoublePair := allPairs[nextSeq]
		if isDoublePair {
			// if the last 2 seq are not the same, then we have a double pair
			if nextSeq != lastSeq {
				return true
			}
			// if last seq is repeated, then also it's a double pair
			// because we are doubling someting from before
			if !lastSeqIsOnlySeq {
				return true
			}
		}

		// prepare next iteration
		lastSeqIsOnlySeq = !isDoublePair
		allPairs[nextSeq] = struct{}{}
		lastChar = nextChar
		lastSeq = nextSeq
	}
	return false
}
