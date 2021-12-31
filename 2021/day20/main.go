package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	part1(test1)
	part1(input1)
}

type point struct {
	y, x int
}

func part1(data []byte) {
	var enhancement []byte
	input := map[point]struct{}{}
	height := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if enhancement == nil {
			enhancement = line
			continue
		}

		for i, c := range line {
			if c == '#' {
				input[point{height, i}] = struct{}{}
			}
		}
		height++
	}
	// draw(input)
	steps := 50
	for s := 0; s < steps; s++ {
		on := false
		if enhancement[0] == '#' {
			on = s%2 == 0
		}
		input = simulate(input, enhancement, on)
		// draw(input)
	}

	fmt.Printf("Part 1: %d\n", len(input))
}

func simulate(input map[point]struct{}, enhancement []byte, on bool) map[point]struct{} {
	minh, maxh, minw, maxw := dims(input)

	dirs := []int{-1, 0, 1}
	dist := 1

	ninput := map[point]struct{}{}
	for y := minh - dist; y <= maxh+dist; y++ {
		for x := minw - dist; x <= maxw+dist; x++ {
			total := 0
			for _, dy := range dirs {
				for _, dx := range dirs {
					total = total << 1
					if _, ok := input[point{dy + y, dx + x}]; ok == on {
						total++
					}
				}
			}
			if (enhancement[total] == '#') != on {
				ninput[point{y, x}] = struct{}{}
			}
		}
	}
	return ninput
}

func dims(input map[point]struct{}) (int, int, int, int) {
	minw, maxw, minh, maxh := 0, 0, 0, 0
	for p := range input {
		if p.y > maxh {
			maxh = p.y
		} else if p.y < minh {
			minh = p.y
		}
		if p.x > maxw {
			maxw = p.x
		} else if p.x < minw {
			minw = p.x
		}
	}
	return minh, maxh, minw, maxw
}

func draw(input map[point]struct{}) {
	minh, maxh, minw, maxw := dims(input)

	fmt.Println("--------------------------------")
	for y := minh; y <= maxh; y++ {
		for x := minw; x <= maxw; x++ {
			if _, ok := input[point{y, x}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------")
}
