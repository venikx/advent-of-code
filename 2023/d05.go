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

	var smallest_location int64 = math.MaxInt64
	for seedIndex := 0; seedIndex < len(seeds); seedIndex++ {
		seed := seeds[seedIndex]

		for i := 0; i < len(steps); i++ {
			step := steps[i]

			for _, step_struct := range step {
				min := step_struct.source
				max := step_struct.source + step_struct.step_range
				// 79    99
				// source = 50 (min) + 48 = 98 (max)
				// destincation = 52 + 48 = 100
				// destination - source = 2
				// 79 + 2 = 81

				if seed >= min && seed < max {
					seed = seed + step_struct.destination - step_struct.source
					//fmt.Println(seed, min, max, step_struct)
					break
				}
			}
		}

		if seed < smallest_location {
			smallest_location = seed
		}
	}

	fmt.Println("Smallest location: ", smallest_location)
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

func calculate_next_step() {

}
