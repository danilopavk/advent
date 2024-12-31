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

	for scanner.Scan() {
		gate := scanner.Text()
		if gate == "" {
			break
		}
	}

	operationsPerResults := map[string]operation{}
	operationsPerOperations := map[string][]operation{
		"AND": {},
		"OR":  {},
		"XOR": {},
	}
	for scanner.Scan() {
		operation := parseOperation(scanner.Text())
		operationsPerResults[operation.result] = operation
		operationsPerOperations[operation.operation] = append(operationsPerOperations[operation.operation], operation)

	}
	invalidGates := map[string]bool{}
	for i := 1; i < 44; i++ {
		gate := toGateName(i, "z")
		operation := operationsPerResults[gate]
		if operation.operation != "XOR" {
			invalidGates[gate] = true
			continue
		}
		left := operation.left[:1]
		if left == "x" || left == "y" {
			invalidGates[gate] = true
			continue
		}
		right := operation.right[:1]
		if right == "x" || right == "y" {
			invalidGates[gate] = true
			continue
		}

	}

	for _, orOperation := range operationsPerOperations["OR"] {
		leftOperation := operationsPerResults[orOperation.left]
		if leftOperation.operation != "AND" {
			invalidGates[orOperation.left] = true
		}

		rightOperation := operationsPerResults[orOperation.right]
		if rightOperation.operation != "AND" {
			invalidGates[orOperation.right] = true
		}
	}

	for _, andOperation := range operationsPerOperations["AND"] {
		leftOperation := operationsPerResults[andOperation.left]
		if leftOperation.left != "x00" && leftOperation.right != "x00" && leftOperation.operation == "AND" {
			invalidGates[andOperation.left] = true
			continue
		}
		rightOperation := operationsPerResults[andOperation.right]
		if rightOperation.left != "x00" && rightOperation.right != "x00" && rightOperation.operation == "AND" {
			invalidGates[andOperation.right] = true
			continue
		}
	}

	for _, xorOperation := range operationsPerOperations["XOR"] {
		if xorOperation.result[:1] == "z" {
			continue
		}

		leftPreOperation := operationsPerResults[xorOperation.left].operation
		rightPreOperation := operationsPerResults[xorOperation.right].operation
		if (leftPreOperation == "OR" && rightPreOperation == "XOR") ||
			(leftPreOperation == "XOR" && rightPreOperation == "OR") {
			invalidGates[xorOperation.result] = true
		}

	}

	fmt.Println(format(invalidGates))
}

func toGateName(counter int, letter string) string {
	if counter < 10 {
		return fmt.Sprintf("%s0%d", letter, counter)
	} else {
		return fmt.Sprintf("%s%d", letter, counter)
	}
}

func format(gates map[string]bool) string {
	gatesSlice := []string{}
	for gate := range gates {
		gatesSlice = append(gatesSlice, gate)
	}
	sort.Strings(gatesSlice)
	return strings.Join(gatesSlice, ",")
}

func parseOperation(input string) operation {
	inputParts := strings.Split(input, " -> ")
	operationParts := strings.Split(inputParts[0], " ")
	return operation{
		operationParts[0],
		operationParts[2],
		operationParts[1],
		inputParts[1],
	}
}

type operation struct {
	left, right, operation, result string
}
