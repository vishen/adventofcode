package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input1.txt
	input1 string
)

func main() {
	part12(input1)
}

func part12(data string) {
	total := 0
	totalFuel := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		f := fuel(ci(line))
		totalFuel += f
		total += f
		for {
			f = fuel(f)
			total += f
			if f == 0 {
				break
			}
		}

	}
	fmt.Printf("Part 1: %d\n", totalFuel)
	fmt.Printf("Part 2: %d\n", total)
}

func fuel(v int) int {

	return abs((v / 3) - 2)
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return 0
}

func ci(v string) int {
	vi, _ := strconv.Atoi(v)
	return vi
}
