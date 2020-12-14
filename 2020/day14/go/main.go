package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	run(data)
}

func run(data []byte) {
	mem := map[int]int{}
	var mask map[int]byte

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		// Handle mask
		if bytes.Equal(line[0:4], []byte("mask")) {
			mask = map[int]byte{}
			// Ignore the "mask = "
			for i, b := range line[7:] {
				mask[35-i] = b
			}
			continue
		}

		// Handle memory

		var loc int
		var val int
		fmt.Sscanf(string(line), "mem[%d] = %d", &loc, &val)
		fmt.Println(loc, val)

		// Part 1
		if false {
			for pos, m := range mask {
				if m == '1' {
					val |= (1 << pos)
				} else if m == '0' {
					val &= ^(1 << pos)
				}
			}
			mem[loc] = val
		}

		// Part 2
		// First pass; apply the mask of 1 bits
		for pos, m := range mask {
			if m == '1' {
				loc |= (1 << pos)
			}
		}

		mems := map[int]bool{
			loc: true,
		}
		// Second pass; handle the X's variations
		for pos, m := range mask {
			if m == 'X' {
				for mem := range mems {
					mems[mem|(1<<pos)] = true  // set 1's
					mems[mem&^(1<<pos)] = true // set 0's
				}
			}
		}

		for m := range mems {
			mem[m] = val
		}
	}

	fmt.Println(mask)
	fmt.Println(mem)

	total := 0
	for _, val := range mem {
		total += val
	}
	fmt.Println(total)
}
