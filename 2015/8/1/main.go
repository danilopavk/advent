package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Evaluating strings, fun! I learned ContainsRune function
func main() {
	file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() 

	total := 0
	evaluated := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    line := scanner.Text()
		total += len(line)
		evaluated += evaluate(line)
	}

	fmt.Printf("%d", total - evaluated)

}

func evaluate(line string) int {
    evaluated := line[1:len(line)-1]
	length := 0
	for i := 0; i < len(evaluated); i++ {
		length++
		if (evaluated[i] != '\\') {
			continue
		}
		if (evaluated[i + 1] == '\\' || evaluated[i + 1] == '"') {
			i++
			continue
		}
		if (evaluated[i + 1] == 'x' && isHex(evaluated[i + 2]) && isHex(evaluated[i + 3])) {
			i += 3
			continue
		}
	}
	return length
}

func isHex(b byte) bool {
	return strings.ContainsRune("0123456789abcdef", rune(b))
}