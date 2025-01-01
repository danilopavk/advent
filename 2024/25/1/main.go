package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	locks := []code{}
	keys := []code{}

	for scanner.Scan() {
		firstLine := scanner.Text()
		if firstLine[0] == '#' {
			locks = parse(locks, '#', scanner)
		} else {
			keys = parse(keys, '.', scanner)
		}
	}

	matchedPairs := 0
	for _, lock := range locks {
		for _, key := range keys {
			matched := true

			if lock.a > key.a {
				matched = false
			}
			if lock.b > key.b {
				matched = false
			}
			if lock.c > key.c {
				matched = false
			}
			if lock.d > key.d {
				matched = false
			}
			if lock.e > key.e {
				matched = false
			}

			if matched {
				matchedPairs++
			}
		}
	}

	fmt.Println(matchedPairs)

}

func parse(codes []code, schemaSign byte, scanner *bufio.Scanner) []code {
	row := 0
	code := code{-1, -1, -1, -1, -1}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if code.a == -1 && line[0] != schemaSign {
			code.a = row
		}
		if code.b == -1 && line[1] != schemaSign {
			code.b = row
		}
		if code.c == -1 && line[2] != schemaSign {
			code.c = row
		}
		if code.d == -1 && line[3] != schemaSign {
			code.d = row
		}
		if code.e == -1 && line[4] != schemaSign {
			code.e = row
		}
		row++
	}
	return append(codes, code)
}

type code struct {
	a, b, c, d, e int
}
