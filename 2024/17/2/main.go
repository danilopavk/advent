package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var program []int
	for scanner.Scan() {
		scanner.Text()
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		program = parseProgram(scanner.Text())
	}

	fmt.Println(getA(program))

}

// works only with my input, since I reverse engineered the solution from it
// maybe I'll build a generic solution one day (probably not)
// my input: 2,4,1,1,7,5,0,3,1,4,4,0,5,5,3,0
func getA(program []int) int {
	a, b, c := 0, 0, 0
	previouses := []int{a}
	for i := len(program) - 1; i >= 0; i-- {
		nextPreviouses := []int{}
		for _, previous := range previouses {
			for t := 0; t < 8; t++ {
				a = (previous * 8) + t
				b = a % 8
				b = b ^ 1
				c = int(float64(a) / (math.Pow(2, float64(b))))
				b = b ^ 4
				b = b ^ c
				b = b % 8
				if b == program[i] {
					nextPreviouses = append(nextPreviouses, a)
				}
			}
		}
		previouses = nextPreviouses
	}
	min := math.MaxInt
	for _, a := range previouses {
		if a < min {
			min = a
		}
	}
	return min
}

func parseProgram(line string) []int {
	trimmed := line[9:]
	parts := strings.Split(trimmed, ",")
	program := []int{}
	for _, part := range parts {
		instruction, _ := strconv.Atoi(part)
		program = append(program, instruction)
	}
	return program
}
