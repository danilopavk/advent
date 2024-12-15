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
	scanner := bufio.NewScanner(file)

	x := 101
	y := 103
	robots := []robot{}
	for scanner.Scan() {
		robot := parseRobot(scanner.Text())
		robot = move(robot, 100, x, y)
		robots = append(robots, robot)
	}
	halfX := x / 2
	halfY := y / 2
	a := countIn(robots, 0, halfX - 1, 0, halfY - 1)
	b := countIn(robots, 0, halfX - 1, halfY + 1, y - 1)
	c := countIn(robots, halfX + 1, x - 1, 0, halfY - 1)
	d := countIn(robots, halfX + 1, x - 1, halfY + 1, y - 1)

	fmt.Println(a * b * c * d)
}

func parseRobot(input string) robot {
	parts := strings.Split(input, " ")
	
	p := strings.Split(parts[0], "=")[1]
	pParts := strings.Split(p, ",")
	x, _ := strconv.Atoi(pParts[0])
	y, _ := strconv.Atoi(pParts[1])

	v := strings.Split(parts[1], "=")[1]
	vParts := strings.Split(v, ",")
	velX, _ := strconv.Atoi(vParts[0])
	velY, _ := strconv.Atoi(vParts[1])

	return robot{x, y, velX, velY}
}

func move(robot robot, moves int, xLength int, yLength int) robot {
	x := robot.x + moves * robot.velX
	y := robot.y + moves * robot.velY

	x = x % xLength
	if x < 0 {
		x += xLength
	}
	robot.x = x

	y = y % yLength
	if y < 0 {
		y += yLength
	}
	robot.y = y

	return robot
}

func countIn(robots []robot, fromX, toX, fromY, toY int) int {
	count := 0
	for _, robot := range robots {
		if robot.x >= fromX && robot.x <= toX && robot.y >= fromY && robot.y <= toY {
			count++
		}
	}
	return count
}

type robot struct {
	x, y, velX, velY int
}
