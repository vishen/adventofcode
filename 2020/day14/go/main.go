package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	data, err := ioutil.ReadFile("./input2.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	run(data)
}

func run(data []byte) {
	mem := map[int]uint64{}
	// var mask uint64 = math.MaxUint64
	var mask uint64 = 0

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		// Handle mask
		if bytes.Equal(line[0:4], []byte("mask")) {
			// Ignore the "mask = "
			for i, b := range line[7:] {
				switch b {
				case 'X':
					continue
				case '1':
					mask |= 1 << (35 - i)
				case '0':
					var m uint64 = ^(1 << (35 - i))
					mask &= m
				}
			}
			continue
		}

		// Handle memory

		var loc int
		var val uint64
		fmt.Sscanf(string(line), "mem[%d] = %d", &loc, &val)
		fmt.Println(loc, val)

		// fmt.Printf("mask=%b\n", mask)
		// fmt.Printf("val=%b\n", val)

		m := math.MaxUint64 & mask
		fmt.Printf("mask=%b\n", mask)
		fmt.Printf("m=%b\n", m)
		mem[loc] = (val & m) | mask
		fmt.Println(mem)
	}

	fmt.Println(mask)
	fmt.Println(mem)

}
