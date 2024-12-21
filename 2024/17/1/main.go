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

	var a, b, c int
	var program []int
	for scanner.Scan() {
		a = parseRegister(scanner.Text())
		scanner.Scan()
		b = parseRegister(scanner.Text())
		scanner.Scan()
		c = parseRegister(scanner.Text())
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		program = parseProgram(scanner.Text())
	}
	output := []string{}

	i := 0
	for i < len(program)-1 {
		opcode := program[i]
		operand := program[i+1]
		combo := combo(operand, a, b, c)
		pow := int(float64(a) / (math.Pow(2, float64(combo))))
		switch opcode {
		case 0:
			a = pow
		case 1:
			b = b ^ operand
		case 2:
			b = combo % 8
		case 3:
			{
				if a != 0 {
					i = operand
					continue
				}
			}
		case 4:
			b = b ^ c
		case 5:
			output = append(output, strconv.Itoa(combo%8))
		case 6:
			b = pow
		case 7:
			c = pow
		}
		i += 2
	}

	fmt.Println(strings.Join(output, ","))
}

func combo(instruction, a, b, c int) int {
	switch instruction {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return instruction
	}
}

func parseRegister(line string) int {
	trimmed := line[12:]
	register, _ := strconv.Atoi(trimmed)
	return register
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
