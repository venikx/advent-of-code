package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	destination, source, step_range int64
}

func main() {
	b, err := os.ReadFile("./d05_prod")

	if err != nil {
		log.Fatal(err)
	}

	var seeds = []int64{}
	var steps = [][]Step{}
	lines := strings.Split(string(b), "\n")

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		if i == 0 {
			seeds = parse_numbers(strings.Split(lines[i], ": ")[1])
			continue
		}

		if strings.Contains(lines[i], ":") {
			steps = append(steps, make([]Step, 0))
		} else {
			step := len(steps) - 1
			stepAsNumbers := parse_numbers(lines[i])
			steps[step] = append(steps[step], Step{stepAsNumbers[0], stepAsNumbers[1], stepAsNumbers[2]})
		}
	}

	// part 1
	var smallest_location int64 = math.MaxInt64

	for i := 0; i < len(seeds); i++ {
		identifier := calculate_location(steps, seeds[i])

		if identifier < smallest_location {
			smallest_location = identifier
		}
	}

	fmt.Println("Part 1: ", smallest_location)

	// part 2
	var smallest_location_from_range int64 = math.MaxInt64

	for i := int64(0); i < int64(len(seeds)); i += 2 {
		for j := int64(0); j < seeds[i+1]; j++ {
			seed := seeds[i] + j
			identifier := calculate_location(steps, seed)

			if identifier < smallest_location_from_range {
				smallest_location_from_range = identifier
			}
		}

	}

	fmt.Println("Part 2: ", smallest_location_from_range)
}

func parse_numbers(line string) []int64 {
	var numbers []int64
	numberStrings := strings.Split(line, " ")
	for _, s := range numberStrings {
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			numbers = append(numbers, v)
		}
	}

	return numbers
}

func calculate_location(steps [][]Step, currentId int64) int64 {
	for i := 0; i < len(steps); i++ {
		step := steps[i]

		for _, step_struct := range step {
			min := step_struct.source
			max := step_struct.source + step_struct.step_range

			if currentId >= min && currentId < max {
				currentId = currentId + step_struct.destination - step_struct.source
				break
			}
		}
	}
	return currentId
}
