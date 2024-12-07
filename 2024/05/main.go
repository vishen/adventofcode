package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
	parts := strings.Split(data, "\n\n")

	deps := map[string][]string{}

	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, "|")
		n1 := nums[0]
		n2 := nums[1]
		deps[n2] = append(deps[n2], n1)
	}

	total := 0
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, ",")

		good := true
		for ni, n := range nums {
			ds, ok := deps[n]
			if !ok {
				continue
			}
			for _, d := range ds {
				for _, n := range nums[ni:] {
					if n == d {
						good = false
						break
					}
				}
				if !good {
					break
				}
			}
		}
		if good {
			ni, _ := strconv.Atoi(nums[len(nums)/2])
			total += ni
		}
	}
	fmt.Println(total)
}

func part2(data string) {
	parts := strings.Split(data, "\n\n")

	deps := map[string][]string{}

	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, "|")
		n1 := nums[0]
		n2 := nums[1]
		deps[n2] = append(deps[n2], n1)
	}

	total := 0
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, ",")

		fixed := false
		bad := false
		for ni := 0; ni < len(nums); ni++ {
			n := nums[ni]
			ds, ok := deps[n]
			if !ok {
				continue
			}
			for _, d := range ds {
				for nni, nn := range nums[ni:] {
					if nn == d {
						fixed = true
						nums = move(nums, ni, ni+nni)
						bad = true
						break
					}
				}
				if fixed {
					ni--
					fixed = false
					break
				}
			}

		}
		if bad {
			ni, _ := strconv.Atoi(nums[len(nums)/2])
			total += ni
		}
	}
	fmt.Println(total)
}

func move(items []string, start, end int) []string {
	cur := items[start]
	for i := start + 1; i <= end; i++ {
		items[i-1] = items[i]
	}
	items[end] = cur
	return items
}
