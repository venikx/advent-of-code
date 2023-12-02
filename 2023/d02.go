package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// for part 1
var possible_cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	readFile, err := os.Open("./d02_prod")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	idx := 0
	sum_of_possible_games := 0
	sum_of_minimal_cube_set := 0

	for fileScanner.Scan() {
		idx++

		is_possible := true
		game := strings.Split(fileScanner.Text(), ":")
		parts_of_game := strings.Split(game[1], ";")

		// for part 2
		least_possible_cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, part := range parts_of_game {
			for _, cubes_of_one_color := range strings.Split(part, ",") {
				cube := strings.Split(strings.Trim(cubes_of_one_color, " "), " ")
				cube_value, err := strconv.Atoi(cube[0])

				if err != nil {
					log.Fatal(err)
				} else {
					if cube_value > possible_cubes[cube[1]] {
						is_possible = false
					}

					if cube_value > least_possible_cubes[cube[1]] {
						least_possible_cubes[cube[1]] = cube_value
					}
				}
			}
		}

		if is_possible {
			sum_of_possible_games += idx
		}

		product := least_possible_cubes["red"] * least_possible_cubes["green"] * least_possible_cubes["blue"]
		sum_of_minimal_cube_set += product
	}

	readFile.Close()

	fmt.Println(sum_of_possible_games)
	fmt.Println(sum_of_minimal_cube_set)
}
