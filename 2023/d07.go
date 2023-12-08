package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var cards = map[rune]uint8{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func main() {
	b, err := os.ReadFile("./d07_prod")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1]

	// part 1
	qs(lines, 0, len(lines)-1, false) // last line is empty char
	sum := 0
	for i, line := range lines {
		_, _, bet := parse_line(line, false)
		sum += (i + 1) * bet
	}
	fmt.Println(sum)

	// part 2
	qs(lines, 0, len(lines)-1, true) // last line is empty char
	sum_with_joker := 0

	for i, line := range lines {
		_, _, bet := parse_line(line, false)
		sum_with_joker += (i + 1) * bet
	}
	fmt.Println(sum_with_joker)

}

func parse_line(line string, with_joker bool) (string, int8, int) {
	s := strings.Split(line, " ")
	var bid int

	if v, err := strconv.Atoi(s[1]); err == nil {
		bid = v
	}

	hand := [13]int8{}

	var max_card rune
	for _, c := range s[0] {
		hand[cards[c]] = hand[cards[c]] << 1
		hand[cards[c]] = hand[cards[c]] | 1

		if with_joker && c != 'J' && hand[cards[c]] > hand[cards[max_card]] {
			max_card = c
		}
	}

	if with_joker {
		for hand[cards['J']] > 0 {
			hand[cards[max_card]] = hand[cards[max_card]] << 1
			hand[cards[max_card]] = hand[cards[max_card]] | 1
			hand[cards['J']] = hand[cards['J']] >> 1
		}
	}

	var sum int8 = 0
	for _, b := range hand {
		sum += b
	}

	return s[0], sum, bid
}

func needs_sorting(line_a string, line_b string, with_joker bool) bool {
	hand_a, hand_value_a, _ := parse_line(line_a, with_joker)
	hand_b, hand_value_b, _ := parse_line(line_b, with_joker)

	if hand_value_a < hand_value_b {
		return true
	} else if hand_value_a == hand_value_b {
		for idx := 0; idx < len(hand_a); idx++ {
			card_a, _ := utf8.DecodeRune(append([]byte{}, hand_a[idx]))
			card_b, _ := utf8.DecodeRune(append([]byte{}, hand_b[idx]))

			if cards[card_a] == cards[card_b] {
				continue
			}

			if with_joker && (card_a == 'J' || card_b == 'J') {
				return card_a == 'J'

			}

			return cards[card_a] < cards[card_b]
		}
	}

	return false
}

func partition(arr []string, lo int, hi int, with_joker bool) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if needs_sorting(arr[i], pivot, with_joker) {
			idx++

			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func qs(arr []string, lo int, hi int, with_joker bool) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi, with_joker)
	qs(arr, lo, pivotIdx-1, with_joker)
	qs(arr, pivotIdx, hi, with_joker)
}
