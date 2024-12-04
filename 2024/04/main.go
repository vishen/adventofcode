package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed sample
	sample string

	//go:embed input1
	input1 string
)

func main() {
	part1(sample)
	part1(input1)
	part2(sample)
	part2(input1)
}

func part1(data string) {
	words := []string{
		"XMAS",
		"SAMX",
	}

	total := 0
	lines := strings.Split(data, "\n")
	lines = lines[:len(lines)-1]

	for row, line := range lines {
		for cur, _ := range line {
			for _, w := range words {
				across, down, diag1, diag2 := 0, 0, 0, 0
				for wi, chr := range w {
					ch := byte(chr)
					if cur+wi < len(line) && line[cur+wi] == ch {
						across++
					}
					if row+wi < len(lines) && lines[row+wi][cur] == ch {
						down++
					}
					if cur+wi < len(line) && row+wi < len(lines) && lines[row+wi][cur+wi] == ch {
						diag1++
					}
					if cur-wi >= 0 && row+wi < len(lines) && lines[row+wi][cur-wi] == ch {
						diag2++
					}
				}
				if across == len(w) {
					total++
				}
				if down == len(w) {
					total++
				}
				if diag1 == len(w) {
					total++
				}
				if diag2 == len(w) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func part2(data string) {
	words := []string{
		"MAS",
		"SAM",
	}

	lines := strings.Split(data, "\n")
	lines = lines[:len(lines)-1]

	totals := map[int]int{}

	for row, line := range lines {
		for cur, _ := range line {
			for _, w := range words {
				diag1, diag2 := 0, 0
				for wi, chr := range w {
					ch := byte(chr)
					if cur+wi < len(line) && row+wi < len(lines) && lines[row+wi][cur+wi] == ch {
						diag1++
					}
					if cur-wi >= 0 && row+wi < len(lines) && lines[row+wi][cur-wi] == ch {
						diag2++
					}
				}
				if diag1 == len(w) {
					totals[key(row+1, cur+1)]++
				}
				if diag2 == len(w) {
					totals[key(row+1, cur-1)]++
				}
			}
		}
	}
	total := 0
	for _, t := range totals {
		if t >= 2 {
			total++
		}
	}
	fmt.Println(total)
}

func key(row, col int) int {
	return (row * 1000) + col
}
