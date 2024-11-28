package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Playing with pointers in order to modify the content of an array
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	raindeers := []*Raindeer{}
	for scanner.Scan() {
		line := scanner.Text()
		raindeers = append(raindeers, parseLine(line))
	}

	maxPoints := 0
	maxDistance := 0
	for time := 0; time < 2503; time++ {
		for _, raindeerPtr := range raindeers {
			raindeer := *raindeerPtr
			raindeer.lastTimeSwitch++
			if raindeer.isRunning {
				raindeer.distance += raindeer.speed
				if raindeer.distance > maxDistance {
					maxDistance = raindeer.distance
				}
				if raindeer.lastTimeSwitch == raindeer.run {
					raindeer.isRunning = false
					raindeer.lastTimeSwitch = 0
				}
			} else if raindeer.lastTimeSwitch == raindeer.rest {
				raindeer.isRunning = true
				raindeer.lastTimeSwitch = 0
			}
			*raindeerPtr = raindeer
		}
		for _, raindeer := range raindeers {
			if raindeer.distance == maxDistance {
				raindeer.points++
				if raindeer.points > maxPoints {
					maxPoints = raindeer.points
				}
			}
		}
	}

	fmt.Println(maxPoints)
}

type Raindeer struct {
	speed, run, rest, distance, points, lastTimeSwitch int
	isRunning                                          bool
}

func parseLine(line string) *Raindeer {
	parts := strings.Split(line, " ")
	speed, _ := strconv.Atoi(parts[3])
	run, _ := strconv.Atoi(parts[6])
	rest, _ := strconv.Atoi(parts[13])
	return &Raindeer{speed, run, rest, 0, 0, 0, true}
}
