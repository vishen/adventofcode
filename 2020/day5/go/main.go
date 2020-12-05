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
	boardingPasses := bytes.Split(data, []byte{'\n'})

	max := 0
	min := 128
	seatsTaken := make(map[int]struct{}, 128)
	for _, b := range boardingPasses {
		if len(b) == 0 {
			break
		}
		row, column := boardingPassSeat(b)
		sid := (row * 8) + column
		if sid > max {
			max = sid
		}
		if sid < min {
			min = sid
		}
		seatsTaken[sid] = struct{}{}
	}
	fmt.Println(max, min)

	for i := min + 1; i < max; i++ {
		if _, ok := seatsTaken[i]; !ok {
			fmt.Println(i)
		}
	}
}

func boardingPassSeat(boardingPass []byte) (int, int) {
	var row int = 0
	for i, c := range boardingPass[0:7] {
		if c == 'B' {
			row = row | (1 << (6 - i))
		}
	}

	var col int = 0
	for i, c := range boardingPass[7:] {
		if c == 'R' {
			col = col | (1 << (2 - i))
		}
	}

	return row, col
}

func boardingPassSeatYuck(boardingPass []byte) (int, int) {
	rows := 7
	seats := 3

	// Basic asset that boarding pass is always 10 digits
	if len(boardingPass) != rows+seats {
		panic(fmt.Sprintf("boarding pass length of %d is invalid", len(boardingPass)))
	}

	row := 0
	{
		min := 0
		max := (1<<rows - 1)
		l := 1 << rows

		for _, c := range boardingPass[0:rows] {
			l = l / 2
			switch c {
			case 'F':
				max -= l
			case 'B':
				min += l
			}
		}
		if min != max {
			panic(fmt.Sprintf("min=%d != max=%d", min, max))
		}
		row = min
	}

	column := 0
	{
		min := 0
		max := (1<<seats - 1)
		l := 1 << seats

		for _, c := range boardingPass[rows:] {
			l = l / 2
			switch c {
			case 'L':
				max -= l
			case 'R':
				min += l
			}
		}
		if min != max {
			panic(fmt.Sprintf("min=%d != max=%d", min, max))
		}
		column = min
	}

	return row, column
}
