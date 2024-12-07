package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
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

	var sum uint64 = 0
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		nums := strings.Split(line, ":")
		target := parseInt(nums[0])
		eqs := parseInts(strings.Split(strings.TrimSpace(nums[1]), " "))
		if valid(target, 0, eqs) {
			sum += uint64(target)
		}
	}
	fmt.Println(sum)
}

func valid(target, cur int, ns []int) bool {
	if len(ns) == 1 {
		return target == cur*ns[0] || target == cur+ns[0]
	}
	if valid(target, cur+ns[0], ns[1:]) {
		return true
	}
	if valid(target, cur*ns[0], ns[1:]) {
		return true
	}
	return false
}

func part2(data string) {

	var sum uint64 = 0
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		nums := strings.Split(line, ":")
		target := parseInt(nums[0])
		eqs := parseInts(strings.Split(strings.TrimSpace(nums[1]), " "))
		if validPart2(target, 0, eqs) {
			sum += uint64(target)
		}
	}
	fmt.Println(sum)
}

func validPart2(target, cur int, ns []int) bool {
	if len(ns) == 1 {
		return target == cur*ns[0] || target == cur+ns[0] || target == concat(cur, ns[0])
	}
	if validPart2(target, cur+ns[0], ns[1:]) {
		return true
	}
	if validPart2(target, cur*ns[0], ns[1:]) {
		return true
	}
	if validPart2(target, concat(cur, ns[0]), ns[1:]) {
		return true
	}
	return false
}

func concat(n1, n2 int) int {
	n, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2))
	return n
}

func parseInt(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func parseInts(nums []string) []int {
	var ints []int
	for _, n := range nums {
		ints = append(ints, parseInt(n))
	}
	return ints
}
