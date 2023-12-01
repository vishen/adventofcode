package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed test
var test []byte

//go:embed test2
var test2 []byte

//go:embed input
var input []byte

func main() {
	part1(input)
	part2(input)
}

func part1(data []byte) {
	total := 0
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		var numbers []int
		for i := 0; i < len(line); i++ {
			c := line[i]

			if c >= '0' && c <= '9' {
				numbers = append(numbers, int(c-'0'))
			}
		}
		if len(numbers) == 1 {
			total += (numbers[0] * 10) + numbers[0]
		} else {
			total += (numbers[0] * 10) + numbers[len(numbers)-1]
		}
	}
	fmt.Println(total)
}

func part2(data []byte) {
	total := 0
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		var numbers []int
		i := 0

		checkWord := func(number string) bool {
			if i+len(number) > len(line) {
				return false
			}
			return string(line[i:i+len(number)]) == number
		}

		wordNumbers := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		for {
			if i >= len(line) {
				break
			}
			c := line[i]
			if c >= '0' && c <= '9' {
				numbers = append(numbers, int(c-'0'))
				i += 1
				continue
			}
			found := false
			for word, number := range wordNumbers {
				if checkWord(word) {
					numbers = append(numbers, number)
					i += len(word)
					found = true
					break
				}
			}
			if !found {
				i += 1
			}
		}
		// fmt.Println(numbers)
		if len(numbers) == 1 {
			total += (numbers[0] * 10) + numbers[0]
		} else {
			total += (numbers[0] * 10) + numbers[len(numbers)-1]
		}
	}
	fmt.Println(total)
}
