package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed sample
	sample string

	//go:embed input
	input string
)

func main() {
	part1(sample)
	part1(input)
}

func part1(data string) {

	var keys [][5]int
	var locks [][5]int

	for _, lines := range strings.Split(strings.TrimSpace(data), "\n\n") {
		isKey := true
		var found [5]int
		for i, line := range strings.Split(lines, "\n") {
			if i == 0 {
				if line == "#####" {
					isKey = false
					continue
				}
			}
			for chi, ch := range line {
				if ch == '#' {
					found[chi]++
				}
			}
		}
		if isKey {
			for i := 0; i < len(found); i++ {
				found[i]--
			}
			keys = append(keys, found)
		} else {
			locks = append(locks, found)
		}
	}

	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			overlap := false
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					overlap = true
					break
				}
			}
			if !overlap {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
