package main

import "fmt"

func main() {
	/*
		input := []int{9, 6, 2, 7, 1, 3, 8, 5, 4}
		part1 := run(input, len(input), 100)
		total := 0
		for i := part1[1]; i != 1; i = part1[i] {
			total = 10*total + i
		}
		fmt.Println("Part 1", total)
	*/

	// input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//run(input, 15, 100)
	// input = []int{9, 6, 2, 7, 1, 3, 8, 5, 4}
	//found := run(input, 1_000_000, 10_000_000)
	found := play("962713854", 1_000_000, 10_000_000)
	fmt.Println(found[1] * found[found[1]])
}

func run(input []int, size, turns int) []int {
	numbers := make([]int, size)

	last := input[0]
	numbers[0] = last
	for _, n := range input[1:] {
		numbers[last] = n
		last = n
	}

	for i := len(input) + 1; i < size; i++ {
		numbers[last] = i
		last = i
	}

	current := numbers[0]
	numbers[last] = current

	for i := 0; i < turns; i++ {
		if i%100_000 == 0 {
			fmt.Println("Turn", i/100_000, "K")
		}

		hold1 := numbers[current]
		hold2 := numbers[hold1]
		hold3 := numbers[hold2]

		destination := current - 1
		for {
			if destination == 0 {
				destination = size - 1
			}
			if destination == hold1 || destination == hold2 || destination == hold3 {
				destination--
				continue
			}
			break
		}

		next := numbers[hold3]
		numbers[current] = next
		current = next
		numbers[hold3] = numbers[destination]
		numbers[destination] = hold1
	}
	return numbers
}

func printNumbers(numbers []int) {
	n := numbers[0]
	fmt.Print(n, " ")
	for i := 0; i < len(numbers)-2; i++ {
		n = numbers[n]
		fmt.Print(n, " ")
	}
	fmt.Println()
}

func play(seed string, cups, rounds int) []int {
	// next[i] is the label of the cup that comes next (clockwise)
	// after the cup labeled i
	next := make([]int, 1+cups)

	last := 0
	for i := 0; i < cups; i++ {
		x := i + 1
		if i < len(seed) {
			x = int(seed[i] - '0')
		}

		next[last] = x
		last = x
	}

	current := next[0]
	next[last] = current
	next[0] = -1e9 // poison

	for i := 0; i < rounds; i++ {
		next1 := next[current]
		next2 := next[next1]
		next3 := next[next2]
		next4 := next[next3]

		dest := current
		for {
			dest--
			if dest == 0 {
				dest = cups
			}
			if dest != next1 && dest != next2 && dest != next3 {
				break
			}
		}

		next[current] = next4
		current = next4

		next[next3] = next[dest]
		next[dest] = next1
	}

	return next
}
