package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input1.txt
var input1 string

//go:embed test1.txt
var input2 string

func main() {
	part1()
	part2()
}

func part1() {
	cur := 0
	max := 0
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}
		v, _ := strconv.Atoi(line)
		cur += v
	}
	fmt.Println(max)
}

func part2() {
	var totals []int

	cur := 0
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			totals = append(totals, cur)
			cur = 0
			continue
		}
		v, _ := strconv.Atoi(line)
		cur += v
	}
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})
	fmt.Println(totals[:3])
	total := 0
	for _, t := range totals[:3] {
		total += t
	}
	fmt.Println(total)
}
