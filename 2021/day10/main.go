package main

import (
	"bytes"
	"fmt"
	"sort"

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
	total := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		ok, wanted, got := parse1(line)
		if !ok {
			fmt.Printf("Syntax error: expected %c, but found %c instead\n", wanted, got)
			switch got {
			case ')':
				total += 3
			case ']':
				total += 57
			case '}':
				total += 1197
			case '>':
				total += 25137
			}
		}
	}
	fmt.Printf("Part 1: score %d\n", total)
}

func parse1(line []byte) (bool, byte, byte) {
	lookingFor := []byte{}
	last := func() byte {
		return lookingFor[len(lookingFor)-1]
	}
	pop := func() {
		lookingFor = lookingFor[:len(lookingFor)-1]
	}

	for _, c := range line {
		switch c {
		case '(':
			lookingFor = append(lookingFor, ')')
		case '[':
			lookingFor = append(lookingFor, ']')
		case '{':
			lookingFor = append(lookingFor, '}')
		case '<':
			lookingFor = append(lookingFor, '>')

		case ')', ']', '}', '>':
			l := last()
			if l != c {
				return false, l, c
			}
			pop()
		}
	}
	return true, 0, 0
}

func part2(data []byte) {
	totals := []int{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		remaining := parse2(line)
		if len(remaining) > 0 {
			total := 0
			for i := len(remaining) - 1; i >= 0; i-- {
				total *= 5
				switch remaining[i] {
				case ')':
					total += 1
				case ']':
					total += 2
				case '}':
					total += 3
				case '>':
					total += 4
				}
			}
			fmt.Printf("remaining=%s, total=%d\n", remaining, total)
			totals = append(totals, total)
		}
	}

	sort.Ints(totals)
	fmt.Printf("Part 2: middle score is %d\n", totals[len(totals)/2])
}

func parse2(line []byte) []byte {
	lookingFor := []byte{}
	last := func() byte {
		return lookingFor[len(lookingFor)-1]
	}
	pop := func() {
		lookingFor = lookingFor[:len(lookingFor)-1]
	}

	for _, c := range line {
		switch c {
		case '(':
			lookingFor = append(lookingFor, ')')
		case '[':
			lookingFor = append(lookingFor, ']')
		case '{':
			lookingFor = append(lookingFor, '}')
		case '<':
			lookingFor = append(lookingFor, '>')

		case ')', ']', '}', '>':
			l := last()
			if l != c {
				return nil
			}
			pop()
		}
	}
	return lookingFor
}
