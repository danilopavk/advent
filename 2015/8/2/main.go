package main

import (
	"bufio"
	"fmt"
	"os"
)

// Simpler version: we start with adding 2 chars for every line
// since extra " is added before and after, and then we are adding
// 1 for every \ or " out there
func main() {
	file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() 

	total := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    line := scanner.Text()
		total += evaluate(line)
	}

	fmt.Printf("%d", total)

}

func evaluate(line string) int {
	increase := 2
	for i := 0; i < len(line); i++ {
		if (line[i] == '\\' || line[i] == '"') {
			increase++
		}
	}
	return increase
}