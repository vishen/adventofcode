package main

import (
	"bytes"
	"fmt"

	_ "embed"
)

var (
	//go:embed test.txt
	test []byte

	//go:embed test2.txt
	test2 []byte

	//go:embed part1.txt
	p1 []byte
)

func main() {
	part1(test)
	part1(p1)
	part2(test2)
	part2(test)
	part2(p1)
}

func part1(data []byte) {
	found := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		lineSplit := bytes.Split(line, []byte{'|'})
		for _, number := range bytes.Split(lineSplit[1], []byte{' '}) {
			switch len(number) {
			case 2, 4, 3, 7:
				found++
			}
		}
	}
	fmt.Printf("Part 1: %d easy numbers\n", found)
}

func part2(data []byte) {
	total := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		lineSplit := bytes.Split(line, []byte{'|'})
		output := determineOutput(lineSplit[0], lineSplit[1])
		total += output
	}
	fmt.Printf("Part 2: %d total output\n", total)
}

func determineOutput(numbers []byte, output []byte) int {

	/*
		  0:      1:      2:      3:      4:
		 aaaa    ....    aaaa    aaaa    ....
		b    c  .    c  .    c  .    c  b    c
		b    c  .    c  .    c  .    c  b    c
		 ....    ....    dddd    dddd    dddd
		e    f  .    f  e    .  .    f  .    f
		e    f  .    f  e    .  .    f  .    f
		 gggg    ....    gggg    gggg    ....

		  5:      6:      7:      8:      9:
		 aaaa    aaaa    aaaa    aaaa    aaaa
		b    .  b    .  .    c  b    c  b    c
		b    .  b    .  .    c  b    c  b    c
		 dddd    dddd    ....    dddd    dddd
		.    f  e    f  .    f  e    f  .    f
		.    f  e    f  .    f  e    f  .    f
		 gggg    gggg    ....    gggg    gggg


		 0: abcefg 	= 6
		 1: cf 		= 2
		 2: acdef	= 5
		 3: acdfg	= 5
		 4: bcdf	= 4
		 5: abdfg	= 5
		 6: abdefg	= 6
		 7: acf 	= 3
		 8: abcdefg = 7
		 9: abcdfg 	= 6

		# 1: len == 2
		# 7: len == 3
		# 8: len == 7
		# 4: len == 4

		# 3: contains 1 and is len==5
		# 6: len == 6 and doesn't contain any of 1

		# 9: len == 6 and contains 4
		# 5: len == 5 and 6 contains 5

		# 2: len == 5
		# 0: len == 6

	*/

	unknown := [][]byte{}
	known := map[int][]byte{}

	for _, number := range bytes.Split(numbers, []byte{' '}) {
		if len(number) == 0 {
			continue
		}
		switch len(number) {
		case 2:
			known[1] = number
		case 4:
			known[4] = number
		case 3:
			known[7] = number
		case 7:
			known[8] = number
		default:
			unknown = append(unknown, number)
		}
	}

	// Look for 3 and 6
	unknown1 := [][]byte{}
	for _, number := range unknown {
		switch len(number) {
		case 5:
			if contains(number, known[1]) {
				known[3] = number
				continue
			}
		case 6:
			if !contains(number, known[1]) {
				known[6] = number
				continue
			}
		}
		unknown1 = append(unknown1, number)
	}

	// Look for 5 and 9
	unknown2 := [][]byte{}
	for _, number := range unknown1 {
		switch len(number) {
		case 6:
			if contains(number, known[4]) {
				known[9] = number
				continue
			}
		case 5:
			if contains(known[6], number) {
				known[5] = number
				continue
			}
		}
		unknown2 = append(unknown2, number)
	}

	// Look for 2 and 0
	for _, number := range unknown2 {
		switch len(number) {
		case 6:
			known[0] = number
		case 5:
			known[2] = number
		}
	}

	/*
		for n, d := range known {
			fmt.Printf("%d: %s\n", n, d)
		}
	*/

	total := 0
	base := 1000
	for _, number := range bytes.Split(output, []byte{' '}) {
		if len(number) == 0 {
			continue
		}

		for n, val := range known {
			if equals(val, number) {
				total += base * n
				base /= 10
				break
			}
		}
	}
	return total
}

func contains(b, sub []byte) bool {
	found := 0
	for _, b1 := range b {
		for _, s := range sub {
			if s == b1 {
				found++
				break
			}
		}
	}
	return found == len(sub)
}

func equals(b, c []byte) bool {
	if len(b) != len(c) {
		return false
	}
	found := 0
	for _, b1 := range b {
		for _, c1 := range c {
			if c1 == b1 {
				found++
				break
			}
		}
	}
	return len(b) == found
}
