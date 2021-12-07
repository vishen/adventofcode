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
)

func main() {
	part1(test)
	part1(p1)
	part2(test)
	part2(p1)
}

func part1(data []byte) {
	crabs := map[int]int{}
	for _, crab := range bytes.Split(data, []byte{','}) {
		if len(crab) == 0 {
			continue
		}
		crabs[convertToInt(crab)]++
	}

	idealFuelUsage := 0
	idealHorMark := 0

	for hor := range crabs {
		fuel := 0
		for crab, count := range crabs {
			if crab == hor {
				continue
			}
			f := (diff(crab, hor) * count)
			fuel += f
		}
		if idealFuelUsage == 0 && idealHorMark == 0 {
			idealFuelUsage = fuel
			idealHorMark = hor
			continue
		}
		if fuel < idealFuelUsage {
			idealFuelUsage = fuel
			idealHorMark = hor
		}
	}

	fmt.Printf("Part 1: fuel_used=%d hor=%d\n", idealFuelUsage, idealHorMark)
}

func part2(data []byte) {
	crabs := map[int]int{}
	max := 0
	for _, crab := range bytes.Split(data, []byte{','}) {
		if len(crab) == 0 {
			continue
		}
		crabI := convertToInt(crab)
		crabs[crabI]++
		if crabI > max {
			max = crabI
		}
	}

	idealFuelUsage := 0
	idealHorMark := 0

	for hor := 0; hor < max; hor++ {
		fuel := 0
		for crab, count := range crabs {
			if crab == hor {
				continue
			}
			f := sum(diff(crab, hor)) * count
			fuel += f
		}
		if idealFuelUsage == 0 && idealHorMark == 0 {
			idealFuelUsage = fuel
			idealHorMark = hor
			continue
		}
		if fuel < idealFuelUsage {
			idealFuelUsage = fuel
			idealHorMark = hor
		}
	}

	fmt.Printf("Part 2: fuel_used=%d hor=%d\n", idealFuelUsage, idealHorMark)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sum(a int) int {
	total := 0
	for i := 1; i <= a; i++ {
		total += i
	}
	return total
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		if val[i] < '0' || val[i] > '9' {
			continue
		}
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
