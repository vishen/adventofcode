package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

func main() {

	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	numbers := make([]int, 0, len(lines)+1)
	numbers = append(numbers, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers = append(numbers, convertToInt(line))
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	diffCount := map[int]int{}
	nextNumbers := map[int][]int{}

	maxJump := 3

	prev := 0
	for i, num := range numbers {
		diffCount[num-prev]++
		prev = num

		for j := 1; j < maxJump+1; j++ {
			n := i + j
			if len(numbers) <= n {
				break
			}
			val := numbers[n]
			if val-num > maxJump {
				break
			}
			nextNumbers[num] = append(nextNumbers[num], val)
		}
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers), diffCount)
	fmt.Println(diffCount[1] * (diffCount[3] + 1))

	// Part 2
	// Naive approach with recursive and memoization
	// total := paths(nextNumbers, 0)

	numbers = append(numbers, 0)
	fmt.Println(numbers)

	total := 1
	prevNum := numbers[0]
	c := 0
	for _, n := range numbers[1:] {
		if n-prevNum == 1 {
			c += 1
		} else {
			// SUPER HACKY
			if n == 0 {
				c += 1
			}
			if c > 2 {
				x := (pow(2, c-1) - pow(2, c-3)) + 1
				fmt.Println(n, c, x)
				total *= x
			} else if c > 0 {
				fmt.Println(n, c, c)
				total *= c
			}
			c = 0
		}
		prevNum = n
	}
	fmt.Println(total)
}

func pow(base, times int) int {
	total := 1
	for i := 0; i < times; i++ {
		total *= base
	}
	return total
}

var seen map[string]int

func hash(num, inc int) string {
	return fmt.Sprintf("%d.%d", num, inc)
}

func paths(nextNumbers map[int][]int, start int) int {
	if seen == nil {
		seen = map[string]int{}
	}
	total := 0
	next := nextNumbers[start]
	if len(next) == 0 {
		return 1
	}
	for _, n := range next {
		hn := hash(start, n)
		t, ok := seen[hn]
		if !ok {
			t = paths(nextNumbers, n)
			seen[hn] = t
		}
		total += t
	}
	return total
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
