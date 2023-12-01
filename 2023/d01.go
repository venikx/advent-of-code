package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digits_as_words = map[string]string{
	"zero":  "z0ero",
	"one":   "o1ne",
	"two":   "t2wo",
	"three": "t3hree",
	"four":  "f4our",
	"five":  "f5ive",
	"six":   "s6ix",
	"seven": "s7even",
	"eight": "e8ight",
	"nine":  "n9ine",
}

func main() {
	b, err := os.ReadFile("./d01_prod")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var digit1 int
		var digit2 int

		for k, v := range digits_as_words {
			line = strings.Replace(line, k, v, -1)
		}

		for _, ch := range line {
			if digit, err := strconv.Atoi(string(ch)); err == nil {
				digit1 = digit
				break
			}

		}

		for i := len(line) - 1; i >= 0; i-- {
			if digit, err := strconv.Atoi(string(line[i])); err == nil {
				digit2 = digit
				break
			}
		}

		digits, err := strconv.Atoi(fmt.Sprintf("%d%d", digit1, digit2))
		if err != nil {
			log.Fatal(err)
		}

		sum += digits
	}

	fmt.Println(sum)
}
