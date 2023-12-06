package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./d06_prod")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	times, time := parse_numbers(lines[0])
	distances, distance := parse_numbers(lines[1])

	fmt.Println(calculate_total(times, distances))
	fmt.Println(calculate_winning_races(time, distance))
}

func parse_numbers(line string) ([]int, int) {
	var numbers []int
	numberStrings := strings.Split(strings.Split(line, ": ")[1], " ")
	var single_number string

	for _, s := range numberStrings {
		if v, err := strconv.Atoi(s); err == nil {
			numbers = append(numbers, v)
			single_number += fmt.Sprint(v) // part 2
		}
	}

	var number int
	if v, err := strconv.Atoi(single_number); err == nil {
		number = v
	}

	return numbers, number
}

func calculate_winning_races(time int, distance_to_break int) int {
	ways_to_win := 0
	for held_button := time / 2; held_button >= 0; held_button-- {
		distance_travelled := held_button * (time - held_button)

		if distance_travelled > distance_to_break {
			points := 2
			if time%2 == 0 && held_button == time/2 {
				points = 1
			}
			ways_to_win += points
		}
	}

	return ways_to_win
}

func calculate_total(times []int, distances []int) int {
	total_races_to_win := 1
	for idx, time := range times {
		winning_races := calculate_winning_races(time, distances[idx])

		if winning_races > 0 {
			total_races_to_win *= winning_races
		}
	}

	return total_races_to_win
}
