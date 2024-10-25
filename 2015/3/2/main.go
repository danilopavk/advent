package main

import (
	"fmt"
	"os"
)

// Small tweek from puzzle 1, alternating between 2 santas.
func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	santa := coordinates{0, 0}
	robot := coordinates{0, 0}
	locations := make(map[coordinates]struct{})
	locations[santa] = struct{}{}
	santaMoves := true
	for _, c := range content {
		if santaMoves {
			santa = move(santa, c)
			locations[santa] = struct{}{}
		} else {
			robot = move(robot, c)
			locations[robot] = struct{}{}
		}
		santaMoves = !santaMoves
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