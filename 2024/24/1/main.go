package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	gates := map[string]bool{}

	for scanner.Scan() {
		gate := scanner.Text()
		if gate == "" {
			break
		}
		parts := strings.Split(gate, ": ")
		gateValue := parts[1] == "1"
		gates[parts[0]] = gateValue

	}

	operations := []operation{}
	for scanner.Scan() {
		operation := parseOperation(scanner.Text())
		operations = append(operations, operation)
	}
	resolveGates(gates, operations)
	binary := getBinary(gates)
	fmt.Println(getDecimal(binary))
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

func resolveGates(gates map[string]bool, operations []operation) {
	for len(operations) > 0 {
		operations = resolvePossibleGates(gates, operations)
	}
}

func resolvePossibleGates(gates map[string]bool, operations []operation) []operation {
	unresolvedOperations := []operation{}
	for _, o := range operations {
		if val, resolved := o.resolve(gates); resolved {
			gates[o.result] = val
		} else {
			unresolvedOperations = append(unresolvedOperations, o)
		}
	}
	return unresolvedOperations
}

func getBinary(gates map[string]bool) string {
	zGates := []string{}
	for gate := range gates {
		if gate[:1] == "z" {
			zGates = append(zGates, gate)
		}
	}
	slices.Sort(zGates)
	slices.Reverse(zGates)

	binary := ""
	for _, zGate := range zGates {
		if gates[zGate] {
			binary = binary + "1"
		} else {
			binary = binary + "0"
		}
	}

	return binary
}

func getDecimal(binary string) int64 {
	decimal, _ := strconv.ParseInt(binary, 2, 64)
	return decimal
}

// first return value is a value of a resolved operation, second value wether it's resolved
func (o operation) resolve(gates map[string]bool) (bool, bool) {
	if _, exists := gates[o.left]; !exists {
		return false, false
	}

	if _, exists := gates[o.right]; !exists {
		return false, false
	}

	switch o.operation {
	case "AND":
		{
			if gates[o.right] && gates[o.left] {
				return true, true
			} else {
				return false, true
			}
		}
	case "OR":
		{
			if gates[o.right] || gates[o.left] {
				return true, true
			} else {
				return false, true
			}
		}
	case "XOR":
		{
			if gates[o.right] && !gates[o.left] {
				return true, true
			}
			if !gates[o.right] && gates[o.left] {
				return true, true
			}
			return false, true
		}
	default:
		return false, false
	}
}

type operation struct {
	left, right, operation, result string
}
