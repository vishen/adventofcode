package main

import (
	"fmt"
	"log"
	"strings"

	_ "embed"
)

var (
	//go:embed test.txt
	test string

	//go:embed part1.txt
	p1 string

	//go:embed part2.txt
	p2 string
)

func main() {
	part1(p1)
	part2(p2)
}

func part1(data string) {
	hor := 0
	depth := 0

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		var command string
		var amount int

		if _, err := fmt.Sscanf(line, "%s %d", &command, &amount); err != nil {
			log.Fatalf("unable to scane %q: %v", line, err)
		}

		switch command {
		case "forward":
			hor += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	fmt.Printf("Part 1: %d\n", hor*depth)
}

func part2(data string) {
	hor := 0
	depth := 0
	aim := 0

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		var command string
		var amount int

		if _, err := fmt.Sscanf(line, "%s %d", &command, &amount); err != nil {
			log.Fatalf("unable to scane %q: %v", line, err)
		}

		switch command {
		case "forward":
			hor += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	fmt.Printf("Part 2: %d\n", hor*depth)
}
