package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed sample1
	sample string

	//go:embed input1
	input string
)

func main() {
	part1(sample)
	part1(input)
	part2(input)
}

func part1(data string) {
	var stones []int

	for _, num := range strings.Split(strings.TrimSpace(data), " ") {
		n, _ := strconv.Atoi(num)
		stones = append(stones, n)
	}

	blinks := 25
	for blink := 0; blink < blinks; blink++ {
		// fmt.Println(blink, stones)
		/*
					If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
					If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

			If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
		*/

		var nstones []int
		for _, s := range stones {
			if s == 0 {
				nstones = append(nstones, 1)
			} else if ss := strconv.Itoa(s); len(ss)%2 == 0 {
				s1, _ := strconv.Atoi(ss[:len(ss)/2])
				s2, _ := strconv.Atoi(ss[len(ss)/2:])
				nstones = append(nstones, s1, s2)
			} else {
				nstones = append(nstones, s*2024)
			}
			stones = nstones
		}
	}
	fmt.Println(len(stones))
}

func part2(data string) {
	stones := 0
	for _, num := range strings.Split(strings.TrimSpace(data), " ") {
		n, _ := strconv.Atoi(num)
		stones += solve(n, 75)
	}
	fmt.Println(stones)
}

type key struct {
	num, t int
}

var cache = map[key]int{}

func solve(num, t int) int {
	if found, ok := cache[key{num, t}]; ok {
		return found
	}
	var l int
	if t == 0 {
		return 1
	} else if num == 0 {
		l = solve(1, t-1)
	} else if ss := strconv.Itoa(num); len(ss)%2 == 0 {
		s1, _ := strconv.Atoi(ss[:len(ss)/2])
		s2, _ := strconv.Atoi(ss[len(ss)/2:])
		l = solve(s1, t-1) + solve(s2, t-1)
	} else {
		l = solve(num*2024, t-1)
	}
	cache[key{num, t}] = l
	return l
}
