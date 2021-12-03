package main

import (
	"bytes"
	"fmt"

	_ "embed"
)

var (
	//go:embed test.txt
	test []byte

	//go:embed part1.txt
	p1 []byte

	//go:embed part2.txt
	p2 []byte
)

func main() {
	part1(p1)
	part2(test)
	part2(p2)
}

func part1(data []byte) {

	var totals []int
	numberOfLines := 0
	lineLength := 0

	for i, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			lineLength = len(line)
			totals = make([]int, lineLength)
		}
		for j, c := range line {
			if c == '1' {
				totals[j]++
			}
		}
		numberOfLines++
	}

	midPoint := numberOfLines / 2

	gamma := 0
	epsilon := 0
	for i, total := range totals {
		if total >= midPoint {
			gamma += (1 << (lineLength - 1 - i))
		} else {
			epsilon += (1 << (lineLength - 1 - i))
		}
	}
	fmt.Printf("Part 1: gamma=%d * epsilon=%d == %d\n", gamma, epsilon, gamma*epsilon)

}

func part2(data []byte) {
	var lines [][]byte

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	oxygen := find(lines, true)

	co2 := find(lines, false)

	fmt.Printf("Part 2: oxygen=%d * co2=%d == %d\n", oxygen, co2, oxygen*co2)
}

func find(lines [][]byte, mostCommon bool) int {
	for i := 0; i < len(lines[0]); i++ {
		var zeroes [][]byte
		var ones [][]byte
		for _, line := range lines {
			if line[i] == '0' {
				zeroes = append(zeroes, line)
			} else {
				ones = append(ones, line)
			}
		}

		if mostCommon {
			if len(ones) >= len(zeroes) {
				lines = ones
			} else {
				lines = zeroes
			}
		} else {
			if len(ones) >= len(zeroes) {
				lines = zeroes
			} else {
				lines = ones
			}
		}

		if len(lines) == 1 {
			break
		}
	}
	total := 0
	for i, val := range lines[0] {
		if val == '1' {
			total += (1 << (len(lines[0]) - 1 - i))
		}
	}
	return total
}
