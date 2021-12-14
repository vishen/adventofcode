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
	part2(test1)
	part2(input1)
}

func part1(data []byte) {
	var poly []byte

	rules := map[[2]byte]byte{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		lineS := bytes.Split(line, []byte{'-', '>'})
		if len(lineS) == 1 {
			poly = lineS[0]
			continue
		}
		rules[[2]byte{lineS[0][0], lineS[0][1]}] = lineS[1][1]
	}

	steps := 10
	for s := 0; s < steps; s++ {
		newPoly := []byte{}
		for p := 0; p < len(poly)-1; p++ {
			newPoly = append(newPoly, poly[p])
			r := [2]byte{poly[p], poly[p+1]}
			if c, ok := rules[r]; ok {
				newPoly = append(newPoly, c)
			}
		}
		newPoly = append(newPoly, poly[len(poly)-1])
		poly = newPoly
	}

	counts := map[byte]int{}
	for _, p := range poly {
		counts[p]++
	}

	min, max := 0, 0
	set := false
	for _, count := range counts {
		if !set {
			min = count
			max = count
			set = true
		} else {
			if min > count {
				min = count
			} else if count > max {
				max = count
			}
		}
	}
	fmt.Printf("Part 1: max=%d - min=%d = %d\n", max, min, max-min)
}

type poly [2]byte

func part2(data []byte) {
	polys := map[poly]int{}
	rules := map[poly]byte{}
	counts := map[byte]int{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		lineS := bytes.Split(line, []byte{'-', '>'})
		if len(lineS) == 1 {
			for _, p := range lineS[0] {
				counts[p]++
			}
			polys = makePolys(lineS[0])
			continue
		}
		p1 := lineS[0][0]
		rules[[2]byte{p1, lineS[0][1]}] = lineS[1][1]
	}

	for s := 0; s < 40; s++ {
		newPolys := map[poly]int{}
		for p, count := range polys {
			c, ok := rules[p]
			if !ok {
				newPolys[p] += count
			} else {
				counts[c] += count
				newPolys[poly{p[0], c}] += count
				newPolys[poly{c, p[1]}] += count
			}
		}
		polys = newPolys
	}

	min, max := 0, 0
	set := false
	for _, count := range counts {
		if !set {
			min = count
			max = count
			set = true
		} else {
			if min > count {
				min = count
			} else if count > max {
				max = count
			}
		}
	}
	fmt.Printf("Part 2: max=%d - min=%d = %d\n", max, min, max-min)
}

func makePolys(p []byte) map[poly]int {
	polys := map[poly]int{}
	for i := 0; i < len(p)-1; i++ {
		polys[poly{p[i], p[i+1]}]++
	}
	return polys
}
