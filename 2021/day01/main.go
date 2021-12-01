package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed part1.txt
	p1 []byte

	//go:embed part2.txt
	p2 []byte

	//go:embed test.txt
	test []byte
)

func main() {
	part1(p1)
	part2(p2)
}

func part2(data []byte) {
	var vals []int

	start := 0
	for i, c := range data {
		if c != '\n' {
			continue
		}
		val := convertToInt(data[start:i])
		vals = append(vals, val)
		start = i + 1
	}

	incremented := 0
	windowSize := 3

	for i := 0; ; i++ {
		if i+windowSize+1 > len(vals) {
			break
		}

		window1 := sum(vals[i : i+windowSize])
		window2 := sum(vals[i+1 : i+windowSize+1])
		if window1 < window2 {
			incremented++
		}

		// fmt.Println(window1, window2, window1 < window2, incremented)
	}

	fmt.Printf("Part 2: %d\n", incremented)
}

func sum(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func part1(data []byte) {
	incremented := 0
	prev := 0

	start := 0
	for i, c := range data {
		if c != '\n' {
			continue
		}
		val := convertToInt(data[start:i])
		start = i + 1

		if prev > 0 && val > prev {
			incremented++
		}
		prev = val
	}

	fmt.Printf("Part 1: %d\n", incremented)
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
