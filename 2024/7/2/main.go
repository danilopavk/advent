package main

import (
	"bufio"
	"fmt"
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

	score := 0
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		calibration := parse(input)
		if isValid(calibration) {
			score += calibration.result
		}
	}

	
	fmt.Println(score)
}

func parse(input string) calibration {
	parts := strings.Split(input, ":")
	
	result, _ := strconv.Atoi(parts[0])

	values := []int{}

	trimmedValuesStr := strings.Trim(parts[1], " ")
	valuesParts := strings.Split(trimmedValuesStr, " ") 
	for i := 0; i < len(valuesParts); i++ {
		valueStr := valuesParts[i]
		value, _ := strconv.Atoi(valueStr)
		values = append(values, value)
	}

	return calibration{result, values}
}

func isValid(calibration calibration) bool {
	return scoreReached(calibration, 0)
}

func scoreReached(equation calibration, score int) bool {
	if (len(equation.values) == 0) {
		return equation.result == score
	}

	if (equation.result < score) {
		return false
	}
 
	leftover := calibration{equation.result, equation.values[1:]}
	val := equation.values[0]
	
	return scoreReached(leftover, score + val) || 
		scoreReached(leftover, score * val) || 
		scoreReached(leftover, concat(score, val))
}

func concat(a int, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	result, _ := strconv.Atoi(aStr + bStr)
	return result
}

type calibration struct {
	result int
	values []int
}