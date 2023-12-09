package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var paths = map[string][2]string{}

func main() {
	b, err := os.ReadFile("./d08_prod")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	input := lines[0]
	ghost_paths := []string{} // part 2

	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		line := strings.Split(lines[i], " = ")
		line_path := strings.Split(line[1], ", ")
		line_path[0] = line_path[0][1:]
		line_path[1] = line_path[1][:len(line_path[1])-1]
		paths[line[0]] = [2]string(line_path)

		if strings.Contains(line[0], "A") {
			ghost_paths = append(ghost_paths, line[0])
		}
	}

	// part 1
	path := "AAA"
	jumps := 0
	for i := 0; path != "ZZZ"; i = (i + 1) % len(input) {
		c := input[i]

		if c == 'L' {
			path = paths[path][0]
		} else if c == 'R' {
			path = paths[path][1]
		}
		jumps++
	}

	fmt.Println(jumps)

	// part 2
	steps := 1
	for idx := range ghost_paths {
		jumps := 0

		for i := 0; !strings.Contains(ghost_paths[idx], "Z"); i = (i + 1) % len(input) {
			c := input[i]

			if c == 'L' {
				ghost_paths[idx] = paths[ghost_paths[idx]][0]
			} else if c == 'R' {
				ghost_paths[idx] = paths[ghost_paths[idx]][1]
			}
			jumps++
		}

		steps = lcm(steps, jumps)
	}

	fmt.Println(steps)
}

func gcd(a int, b int) int { // greater common denominator
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a int, b int) int { // least common denominator
	return (a * b) / gcd(a, b)
}
