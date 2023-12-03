package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coordinate struct {
	x, y int
}

var directions = []coordinate{
	{0, -1},  // top
	{1, 0},   // right
	{0, 1},   // bottom
	{-1, 0},  // left
	{1, -1},  // top right
	{1, 1},   // bottom right
	{-1, 1},  // bottom left
	{-1, -1}, // top left
}

var symbols = map[string][]int{}
var connected_gears = map[coordinate][]int{}

type possible_number struct {
	index coordinate
	value int
}

var possible_numbers = []possible_number{}

func main() {
	b, err := os.ReadFile("./d03_prod")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	for i := 0; i < len(lines); i++ {
		digits := ""

		for j, ch := range lines[i] {
			digitDone := j >= len(lines[i])-1

			if unicode.IsNumber(ch) {
				digits += string(ch)
			} else {
				digitDone = true
			}

			if digitDone {
				digits_as_int, err := strconv.Atoi(digits)

				if err == nil {
					possible_numbers = append(possible_numbers, possible_number{
						coordinate{j - len(digits), i},
						digits_as_int})
					digits = ""
				}
			}
		}
	}

	for i := 0; i < len(lines); i++ {
		for j, ch := range lines[i] {
			if !unicode.IsNumber(ch) && ch != '.' {

				for idx, possible_number := range possible_numbers {
					if symbol_intersects_with_number(possible_number, coordinate{j, i}) {
						symbols[string(ch)] = append(symbols[string(ch)], idx)

						// part 2
						if ch == '*' {
							connected_gears[coordinate{j, i}] = append(connected_gears[coordinate{j, i}], idx)
						}

					}
				}
			}
		}
	}

	sum_of_parts := 0
	for _, indexes_of_part_numbers := range symbols {
		for _, part_number_index := range indexes_of_part_numbers {
			sum_of_parts += possible_numbers[part_number_index].value
		}
	}

	// part 2
	sum_of_product_of_gears := 0
	for _, indexes_of_gear := range connected_gears {
		product_of_gears := 0

		if len(indexes_of_gear) >= 2 {
			product_of_gears = 1
			for _, part_number_index := range indexes_of_gear {
				product_of_gears *= possible_numbers[part_number_index].value
			}
		}

		sum_of_product_of_gears += product_of_gears
	}

	fmt.Println("Part 1: ", sum_of_parts)
	fmt.Println("Part 2: ", sum_of_product_of_gears)
}

func symbol_intersects_with_number(num possible_number, symbol coordinate) bool {
	numberCoordinate := num.index
	numberValue := num.value
	does_intersect := false

	for _, dir := range directions {
		y_to_check := symbol.y + dir.y
		x_to_check := symbol.x + dir.x

		does_intersect = numberCoordinate.y == y_to_check &&
			numberCoordinate.x <= x_to_check &&
			x_to_check < (numberCoordinate.x+len(strconv.Itoa(numberValue)))

		if does_intersect {
			break
		}
	}

	return does_intersect
}
