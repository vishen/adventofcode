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
	dataSplit := bytes.Split(data, []byte{'\n'})

	startTime := convertToInt(dataSplit[0])

	max := 0
	pos := 0

	buses := []int{}
	busPos := map[int]int{}

	for i, b := range bytes.Split(dataSplit[1], []byte{','}) {
		if b[0] == 'x' {
			continue
		}
		bi := convertToInt(b)
		buses = append(buses, bi)
		busPos[bi] = i

		if max < bi {
			max = bi
			pos = i
		}
	}

	fmt.Println(startTime, buses)

	// Part 1
	if false {
		nextBus := buses[0]
		min := buses[0] - (startTime % buses[0])
		for _, b := range buses[1:] {
			diff := b - (startTime % b)
			if diff < min {
				nextBus = b
				min = diff
			}
		}
		fmt.Println(nextBus, min, nextBus*min)
	}

	// Part 2

	fmt.Println(max, pos)
	fmt.Println(busPos)

	if false {
		// Naive approach, take days
		offset := 100000000000000
		starting := offset - (offset % max)
	Outer:
		for i := starting; ; i += max {
			t := i - pos
			for b, j := range busPos {
				if (t+j)%b != 0 {
					continue Outer
				}
			}
			fmt.Println("Found", t)
			break
		}
	}

	t := 0
	mod := 1
	for bus, pos := range busPos {
		for (t+pos)%bus != 0 {
			t += mod
		}
		mod *= bus
	}
	fmt.Println(mod, t)
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
