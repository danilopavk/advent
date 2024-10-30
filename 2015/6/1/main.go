package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

// Brute force algorithm got me the answer pretty fast, so I let it.
// I could probably instead store "lights on from/to" instead of individual lights
// Pending on input data, one might get me better or worse results.
// I learned here how to handle consts as enums in golang
func main() {
	file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() 

	// set up lights
	lights := map[int]map[int]struct{} {}
	for i := 0; i < 1000; i++ {
		lights[i] = map[int]struct{} {}
	}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		instruction := parse(line)
		switch instruction.method {
			case turnOn: handleTurnOn(instruction, lights)
			case turnOff: handleTurnOff(instruction, lights)
			case toggle: handleToggle(instruction, lights)
		}
    }
	// calculate total unlit
	score := 0
	for _, value := range lights {
		score += len(value)
	}
	fmt.Println(score)
}

func handleTurnOn(instruction instruction, lights  map[int]map[int]struct{}) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			lights[i][j] = struct{}{}
		}
	}
}

func handleTurnOff(instruction instruction, lights  map[int]map[int]struct{}) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			delete(lights[i], j)
		}
	}
}

func handleToggle(instruction instruction, lights  map[int]map[int]struct{}) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			if _, exists := lights[i][j]; exists {
				delete(lights[i], j)
			} else {
				lights[i][j] = struct{}{}
			}
		}
	}
}
func parse(line string) instruction {
	method := getMethod(line)
	rest := strings.Replace(line, string(method) + " ", "", 1)
	parts := strings.Split(rest, " ")
	from := strings.Split(parts[0], ",")
	fromX, _ := strconv.Atoi(from[0])
	toX, _ := strconv.Atoi(from[1])
	to := strings.Split(parts[2], ",")
	fromY, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])
	return instruction{fromX, fromY, toX, toY, method}
}

func getMethod(line string) method {
	if strings.HasPrefix(line, "turn on") {
		return turnOn
	}
	if strings.HasPrefix(line, "turn off") {
		return turnOff
	}
	return toggle
}

type instruction struct {
	fromX int
	fromY int
	toX int
	toY int
	method method
}

type method string
const (
	turnOn method = "turn on"
	turnOff method = "turn off"
	toggle method = "toggle"
)