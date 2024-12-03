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

func ints(vs []string) []int {
	var is []int
	for _, v := range vs {
		vi, _ := strconv.Atoi(v)
		is = append(is, vi)
	}
	return is
}

func part1(data string) {
	safe := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		arr := ints(strings.Split(line, " "))
		up := true
		if arr[0]-arr[1] > 0 {
			up = false
		}

		failed := false
		prev := arr[0]
		for _, a := range arr[1:] {
			diff := prev - a
			if up && diff > 0 {
				failed = true
				break
			}
			if !up && diff < 0 {
				failed = true
				break
			}
			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				failed = true
				break
			}
			prev = a
		}
		if !failed {
			safe++
		}
	}
	fmt.Println(safe)
}

func part2(data string) {
	safe := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		arr := ints(strings.Split(line, " "))
		failed := 0
		passed := 0
		for i := 0; i < len(arr); i++ {
			narr := make([]int, len(arr))
			copy(narr, arr)
			narr = append(narr[:i], narr[i+1:]...)
			if !check(narr) {
				failed++
			} else {
				passed++
			}
		}
		//fmt.Println(arr, failed, passed)
		if passed >= 1 {
			safe++
		}
	}
	fmt.Println(safe)
}

func check(arr []int) bool {

	ups, downs := 0, 0

	prev := arr[0]
	for _, a := range arr[1:] {
		if diff := prev - a; diff > 0 {
			downs++
		} else if diff < 0 {
			ups++
		}
	}

	up := ups > downs

	prev = arr[0]
	for _, a := range arr[1:] {
		diff := prev - a
		if up && diff > 0 {
			return false
		}
		if !up && diff < 0 {
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
		prev = a
	}
	return true
}
