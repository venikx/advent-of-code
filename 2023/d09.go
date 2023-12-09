package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	destination, source, step_range int64
}

func main() {
	readFile, err := os.Open("./d09_prod")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total_sum := 0
	total_backwards_sum := 0

	for fileScanner.Scan() {
		numbers := parse_numbers(fileScanner.Text())
		last_numbers := []int{}
		first_numbers := []int{}

		for {
			numbers_sum := 0
			new_numbers := []int{}

			for i := 0; i < len(numbers); i++ {
				numbers_sum += numbers[i]

				if i < len(numbers)-1 {
					new_numbers = append(new_numbers, numbers[i+1]-numbers[i])
				}
			}

			last_numbers = append(last_numbers, numbers[len(numbers)-1])
			first_numbers = append(first_numbers, numbers[0])

			if numbers_sum == 0 {
				break
			}

			numbers = new_numbers
		}

		sum := 0
		backwards_sum := 0
		for i := len(last_numbers) - 1; i >= 0; i-- {
			sum += last_numbers[i]
			backwards_sum = first_numbers[i] - backwards_sum
		}

		total_sum += sum
		total_backwards_sum += backwards_sum
	}

	fmt.Println(total_sum)
	fmt.Println(total_backwards_sum)
}

func parse_numbers(line string) []int {
	var numbers []int
	numberStrings := strings.Split(line, " ")
	for _, s := range numberStrings {
		if v, err := strconv.Atoi(s); err == nil {
			numbers = append(numbers, v)
		}
	}

	return numbers
}
