package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// First time I solved the puzzle without asking AI anything.
// Algorithm isn't the most efficient, because time goes 1 by 1 instead of going the full
// run. This just makes the code more readable, so I sacrificed performance since more
// efficient algorithm isn't needed
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	maxDistance := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		raindeer := parseLine(line)
		maxDistance = max(maxDistance, calculateDistance(raindeer))
	}

	fmt.Println(maxDistance)
}

type Raindeer struct {
	speed, run, rest int
}

func parseLine(line string) Raindeer {
	parts := strings.Split(line, " ")
	speed, _ := strconv.Atoi(parts[3])
	run, _ := strconv.Atoi(parts[6])
	rest, _ := strconv.Atoi(parts[13])
	return Raindeer{speed, run, rest}
}

func calculateDistance(raindeer Raindeer) int {
	distance := 0
	isRunning := true
	end := 2503
	lastTimeSwitch := 0
	for time := 0; time < end; time++ {
		lastTimeSwitch++

		if isRunning {
			distance += raindeer.speed
			if lastTimeSwitch == raindeer.run {
				isRunning = false
				lastTimeSwitch = 0
			}
		} else if lastTimeSwitch == raindeer.rest {
			isRunning = true
			lastTimeSwitch = 0
		}

	}
	return distance
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
