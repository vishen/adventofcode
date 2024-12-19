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
	part2(sample)
	part2(input)
}

func run(target string, towels map[string]bool, seen map[string]int) int {
	if target == "" {
		return 1
	}
	if found, ok := seen[target]; ok {
		return found
	}
	found := 0
	for t := range towels {
		if strings.HasPrefix(target, t) {
			found += run(target[len(t):], towels, seen)
		}
	}
	seen[target] = found
	return found
}

func part1(data string) {
	parts := strings.Split(data, "\n\n")
	towels := map[string]bool{}
	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel] = true
	}

	seen := make(map[string]int)
	possible := 0
	for _, target := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		if run(target, towels, seen) > 0 {
			possible++
		}
	}
	fmt.Println(possible)
}

func part2(data string) {
	parts := strings.Split(data, "\n\n")
	towels := map[string]bool{}
	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel] = true
	}

	seen := make(map[string]int)
	combos := 0
	for _, target := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		combos += run(target, towels, seen)
	}
	fmt.Println(combos)

}
