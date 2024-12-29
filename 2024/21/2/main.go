package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numericalMap := map[rune]cell{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	}

	numericalMoves := expandNumericalMoves(numericalMap)

	complexities := 0

	directionalMoves := map[rune]map[rune]string {
		'<': {
				'^': ">^A",
				'A': ">>^A",
				'v': ">A",
				'>': ">>A",
				'<': "A",
		},
		'^': {
				'<': "v<A",
				'A': ">A",
				'v': "vA",
				'>': "v>A",
				'^': "A",
		},
		'v': {
				'<': "<A",
				'A': "^>A",
				'^': "^A",
				'>': ">A",
				'v': "A",
		},
		'A': {
				'<': "v<<A",
				'^': "<A",
				'v': "<vA",
				'>': "vA",
				'A': "A",
		},
		'>': {
				'<': "<<A",
				'A': "^A",
				'v': "<A",
				'^': "<^A",
				'>': "A",
		},
	}

	for scanner.Scan() {
		input := scanner.Text()
		directions := moveDirections(input, numericalMoves)

		length := calculateLength(directions, directionalMoves, 0, map[cacheKey]int{})
		
		numericalInput, _ := strconv.Atoi(input[:3])
		complexity := numericalInput * length
		
		complexities += complexity
	}

	fmt.Println(complexities)
}

func calculateLength(section string, directionalMoves map[rune]map[rune]string, depth int, cache map[cacheKey]int) int {
	if depth == 25 {
		return len(section)
	}

	memoVal := cacheKey{section, depth}
	if val, exists := cache[memoVal]; exists {
		return val
	}

	length := 0
	from := 'A'
	for _, to:= range section {
		if from == to {
			length += 1
		} else {
			nextSection := directionalMoves[from][to]
			length += calculateLength(nextSection, directionalMoves, depth + 1, cache)
		}
		from = to
	}

	cache[memoVal] = length 
	return length

}

func expandNumericalMoves(pads map[rune]cell) map[rune]map[rune]string {
	moves := make(map[rune]map[rune]string, 5)
	for runeFrom, cellFrom := range pads {
		moves[runeFrom] = make(map[rune]string, 5)
		for runeTo, cellTo := range pads {
			if runeFrom == runeTo {
				moves[runeFrom][runeTo] = "A"
				continue
			}

			code := ""

			if cellFrom.x == cellTo.x {
				code += upDown(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			}

			if cellFrom.y == cellTo.y {
				code += leftRight(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			}

			if cellFrom.y == 3  && cellTo.x == 0 {
				code += upDown(cellFrom, cellTo)
				code += leftRight(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			}

			if cellFrom.x == 0 && cellTo.y == 3 {
				code += leftRight(cellFrom, cellTo)
				code += upDown(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			}

			if cellFrom.x > cellTo.x {
				code += leftRight(cellFrom, cellTo)
				code += upDown(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			} else {
				code += upDown(cellFrom, cellTo)
				code += leftRight(cellFrom, cellTo)
				code += "A"
				moves[runeFrom][runeTo] = code
				continue
			}

		}
	}
	return moves
}

func leftRight(from cell, to cell) string {
	code := ""	
	if from.x > to.x {
		for i := 0; i < from.x - to.x; i++ {
			code += "<"
		}
	} else {
		for i := 0; i < to.x - from.x; i++ {
			code += ">"
		}
	}		
	return code
}

func upDown(from cell, to cell) string {
	code := ""	
	if from.y > to.y {
		for i := 0; i < from.y - to.y; i++ {
			code += "^"
		}
	} else {
		for i := 0; i < to.y - from.y; i++ {
			code += "v"
		}
	}		
	return code
}

func moveDirections(input string, moves map[rune]map[rune]string) string {
	previous := 'A'
	direction := ""

	for _, r := range input {
		direction += moves[previous][r]
		previous = r
	}
	return direction
}

type cell struct {
	x, y int
}

type cacheKey struct {
	section string
	depth int
}
