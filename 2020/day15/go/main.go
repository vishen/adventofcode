package main

import "fmt"

func main() {
	input := []int{20, 0, 1, 11, 6, 3}
	turns := 2020

	input = []int{1, 3, 2}
	turns = 100

	lastSpoken := map[int]int{}
	last := 0

	// Handle the initial starting numbers
	for i, num := range input {
		lastSpoken[num] = i
		fmt.Println(i+1, num)
		last = num
	}

	for i := len(input); i < turns; i++ {
		if i == len(input) {
			fmt.Println(i+1, 0)
			last = 0
			continue
		}

		pos, ok := lastSpoken[last]
		lastSpoken[last] = i - 1
		if ok {
			last = i - 1 - pos
		} else {
			last = 0
		}
		fmt.Println(i+1, last)
	}
}
