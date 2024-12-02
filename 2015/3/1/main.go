package main

import (
	"fmt"
	"os"
)

// Learned how sets work in golang, or the lack of them!
func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	santa := coordinates{0, 0}
	locations := make(map[coordinates]struct{})
	locations[santa] = struct{}{}
	for _, c := range content {
		santa = move(santa, c)
		locations[santa] = struct{}{}
	}
	fmt.Println(len(locations))
}

func move(santa coordinates, c byte) coordinates {
	if c == '^' {
		return coordinates{santa.X + 1, santa.Y}
	}
	if c == 'v' {
		return coordinates{santa.X - 1, santa.Y}
	}
	if c == '<' {
		return coordinates{santa.X, santa.Y - 1}
	}
	if c == '>' {
		return coordinates{santa.X, santa.Y + 1}
	}
	return coordinates{santa.X, santa.Y}
}

type coordinates struct {
	X int
	Y int
}
