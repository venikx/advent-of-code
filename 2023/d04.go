package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const (
	Card_Details int = 0
	Winning_Numbers
	Owned_Numbers
)

func main() {
	readFile, err := os.Open("./d04_prod")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sumBasedOnElf := 0 // part 1
	idx := 0           // part 2

	for fileScanner.Scan() {
		card_split := strings.Split(fileScanner.Text(), ":")
		numbers_split := strings.Split(card_split[1], "|")

		idx++
		if points := get_points_based_from_elf(numbers_split[0], numbers_split[1]); points != 0 {
			sumBasedOnElf += int(math.Pow(2, float64(points)-1))
		}
	}

	readFile.Close()
	fmt.Println(sumBasedOnElf)
}

func get_points_based_from_elf(winning_numbers string, numbers string) int {
	points := 0
	for _, v := range strings.Split(numbers, " ") {
		if v == "" || v == " " {
			continue
		}

		padded_value := " " + v + " "

		if strings.Contains(winning_numbers, padded_value) {
			points += 1
		}
	}

	return points
}
