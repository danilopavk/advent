package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Adjusted for calories = 500
func main() {
	file, err := os.Open("../../../input")
	if err != nil {
		return
	}
	defer file.Close()

	ingredients := []Ingredient{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ingredients = append(ingredients, parseLine(line))
	}
	fmt.Println(iterateVariations(ingredients))
}

func iterateVariations(ingredients []Ingredient) int {
	max := 0
	for first := 0; first <= 100; first++ {
		for second := 0; second <= 100-first; second++ {
			for third := 0; third <= 100-first-second; third++ {
				for forth := 0; forth <= 100-first-second-third; forth++ {
					value := calculate(ingredients, first, second, third, forth)
					if value > max {
						max = value
					}
				}
			}
		}
	}
	return max
}

func calculate(ingredients []Ingredient, first, second, third, forth int) int {
	calories := ingredients[0].calories*first + ingredients[1].calories*second + ingredients[2].calories*third + ingredients[3].calories*forth
	if calories != 500 {
		return 0
	}
	capacity := ingredients[0].capacity*first + ingredients[1].capacity*second + ingredients[2].capacity*third + ingredients[3].capacity*forth
	durability := ingredients[0].durability*first + ingredients[1].durability*second + ingredients[2].durability*third + ingredients[3].durability*forth
	flavor := ingredients[0].flavor*first + ingredients[1].flavor*second + ingredients[2].flavor*third + ingredients[3].flavor*forth
	texture := ingredients[0].texture*first + ingredients[1].texture*second + ingredients[2].texture*third + ingredients[3].texture*forth
	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0
	}
	return capacity * durability * flavor * texture
}

type Ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func parseLine(line string) Ingredient {
	parts := strings.Split(line, " ")
	capacity, _ := strconv.Atoi(parts[2][:len(parts[2])-1])
	durability, _ := strconv.Atoi(parts[4][:len(parts[4])-1])
	flavor, _ := strconv.Atoi(parts[6][:len(parts[6])-1])
	texture, _ := strconv.Atoi(parts[8][:len(parts[8])-1])
	calories, _ := strconv.Atoi(parts[10])
	return Ingredient{capacity, durability, flavor, texture, calories}
}
