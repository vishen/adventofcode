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

	// N is positive vertial
	// S is negative vertical
	// E is positive horizontal
	// W is negative horizontal

	waypointH := 10
	waypointV := 1

	shipH := 0
	shipV := 0

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		c := line[0]
		num := convertToInt(line[1:])
		fmt.Println(string(c), num)

		if c == 'F' {
			shipV += waypointV * num
			shipH += waypointH * num
		}

		r := 0
		dir := 1
		switch c {
		case 'N':
			waypointV += num
		case 'S':
			waypointV -= num
		case 'E':
			waypointH += num
		case 'W':
			waypointH -= num
		case 'R':
			r = num
		case 'L':
			r = num
			dir = -1
		}

		if r%360 == 0 {
			continue
		}

		waypointV, waypointH = rotate(r, dir, waypointV, waypointH)
	}

	if shipV < 0 {
		shipV = -shipV
	}
	if shipH < 0 {
		shipH = -shipH
	}

	fmt.Println(shipV, shipH, shipV+shipH)
}

func rotate(r, dir, v, h int) (int, int) {
	for i := 0; i < r/90; i++ {
		dh, dv := 1*dir, -1*dir
		v, h = h*dv, v*dh
	}
	return v, h
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
