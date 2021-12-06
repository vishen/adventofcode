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
	part1(test)
	part1(p1)
	part2(test)
	part2(p2)
}

func part1(data []byte) {
	fishes := make([]int8, 0, len(data))
	for _, fish := range bytes.Split(data, []byte{','}) {
		if len(fish) == 0 {
			continue
		}
		fishes = append(fishes, int8(fish[0]-'0'))
	}

	days := 80
	reset := int8(6)
	spawn := int8(8)

	for day := 1; day <= days; day++ {
		// log.Printf("day %d: %d fish\n", day, len(fishes))

		for i, fish := range fishes {
			if fish == 0 {
				fishes[i] = reset
				fishes = append(fishes, spawn)
				continue
			}
			fishes[i]--
		}
	}
	fmt.Printf("Part 1: %d fish after %d days\n", len(fishes), days)
}

func part2(data []byte) {

	fishes := map[int]int{}

	for _, fish := range bytes.Split(data, []byte{','}) {
		if len(fish) == 0 {
			continue
		}
		fishes[int(fish[0]-'0')]++
	}

	days := 256
	reset := 6
	spawn := 8

	for day := 1; day <= days; day++ {
		// log.Printf("day %d: %d fish\n", day, len(fishes))

		newFishes := map[int]int{}
		for fish, count := range fishes {
			if fish == 0 {
				newFishes[reset] += count
				newFishes[spawn] += count
			} else {
				newFishes[fish-1] += count
			}
			fishes = newFishes
		}
	}
	total := 0
	for _, count := range fishes {
		total += count
	}
	fmt.Printf("Part 2: %d fish after %d days\n", total, days)
}
