package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed sample1
	sample1 string

	//go:embed input1
	input1 string
)

func main() {
	// part1(sample1)
	part1(input1)
	//part2(sample1)
	part2(input1)
}

func pi(val string) int {
	v, _ := strconv.Atoi(val)
	return v
}

func part1(data string) {
	var n1 []int
	var n2 []int

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "   ")
		n1 = append(n1, pi(parts[0]))
		n2 = append(n2, pi(parts[1]))
	}
	sort.Slice(n1, func(i, j int) bool {
		return n1[i] < n1[j]
	})
	sort.Slice(n2, func(i, j int) bool {
		return n2[i] < n2[j]
	})

	total := 0
	for i, n := range n1 {
		diff := n - n2[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	fmt.Println(total)
}

func part2(data string) {
	var n1 []int
	n2 := map[int]int{}

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "   ")
		n1 = append(n1, pi(parts[0]))
		n2[pi(parts[1])]++
	}

	total := 0
	for _, n := range n1 {
		total += n * n2[n]
	}
	fmt.Println(total)
}
