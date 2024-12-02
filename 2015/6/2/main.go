package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

// little changes incalculation, nothing more
func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// set up lights
	lights := map[int]map[int]int{}
	for i := 0; i < 1000; i++ {
		lights[i] = map[int]int{}
		for j := 0; j < 1000; j++ {
			lights[i][j] = 0
		}
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instruction := parse(line)
		switch instruction.method {
		case turnOn:
			handleTurnOn(instruction, lights)
		case turnOff:
			handleTurnOff(instruction, lights)
		case toggle:
			handleToggle(instruction, lights)
		}
	}
	// calculate total unlit
	score := 0
	for _, row := range lights {
		for _, column := range row {
			score += column
		}
	}
	fmt.Println(score)
}

func handleTurnOn(instruction instruction, lights map[int]map[int]int) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			lights[i][j] = lights[i][j] + 1
		}
	}
}

func handleTurnOff(instruction instruction, lights map[int]map[int]int) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			if lights[i][j] > 0 {
				lights[i][j] = lights[i][j] - 1
			}
		}
	}
}

func handleToggle(instruction instruction, lights map[int]map[int]int) {
	for i := instruction.fromX; i <= instruction.fromY; i++ {
		for j := instruction.toX; j <= instruction.toY; j++ {
			lights[i][j] = lights[i][j] + 2
		}
	}
}
func parse(line string) instruction {
	method := getMethod(line)
	rest := strings.Replace(line, string(method)+" ", "", 1)
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
	fromX  int
	fromY  int
	toX    int
	toY    int
	method method
}

type method string

const (
	turnOn  method = "turn on"
	turnOff method = "turn off"
	toggle  method = "toggle"
)
