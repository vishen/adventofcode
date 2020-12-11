package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	data   []byte
	rowLen int
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(d []byte) {
	data = d
	for i, c := range data {
		if c == '\n' {
			rowLen = i + 1
			break
		}
	}

	prevState := state(data)
	// fmt.Printf("b1:\n%s\n", data)
	for depth := 0; ; depth++ {
		dataNext := make([]byte, len(data))
		copy(dataNext, data)
		for i := range data {
			if d, ok := canChange(i); ok {
				dataNext[i] = d
			}
		}
		// fmt.Printf("b2:\n%s\n", dataNext)
		newState := state(dataNext)
		if bytes.Equal(prevState, newState) {
			break
		}
		// fmt.Println("prev state and current state not the same")
		prevState = newState
		copy(data, dataNext)
	}

	occupiedSeats := 0
	for i := range data {
		if data[i] == '#' {
			occupiedSeats += 1
		}
	}
	fmt.Println(occupiedSeats)
}

func canChange(pos int) (byte, bool) {
	seen := map[string]byte{}

	// Out of bounds check
	oobR := false
	oobL := false
	oobU := false
	oobD := false

	for i := 1; i <= rowLen; i++ {

		// up
		upDelta := -rowLen * i
		up := pos + upDelta
		if up < 0 {
			oobU = true
		} else if !oobU {
			if occupied(up) > 0 {
				if _, ok := seen["up"]; !ok {
					seen["up"] = data[up]
				}
			}
		}

		// down
		downDelta := rowLen * i
		down := pos + downDelta
		if down > rowLen*rowLen {
			oobD = true
		} else if !oobD {
			if occupied(down) > 0 {
				if _, ok := seen["down"]; !ok {
					seen["down"] = data[down]
				}
			}
		}

		// left
		leftDelta := -i
		left := pos + leftDelta
		if left%rowLen == rowLen-1 {
			oobL = true
		} else if !oobL {
			if occupied(left) > 0 {
				if _, ok := seen["left"]; !ok {
					seen["left"] = data[left]
				}
			}
		}

		// right
		rightDelta := i
		right := pos + rightDelta
		if right%rowLen == 0 {
			oobR = true
		} else if !oobR {
			if occupied(right) > 0 {
				if _, ok := seen["right"]; !ok {
					seen["right"] = data[right]
				}
			}
		}

		// up-left
		if !oobU && !oobL {
			if occupied(pos+upDelta+leftDelta) > 0 {
				if _, ok := seen["up-left"]; !ok {
					seen["up-left"] = data[pos+upDelta+leftDelta]
				}
			}
		}

		if !oobU && !oobR {
			// up-right
			if occupied(pos+upDelta+rightDelta) > 0 {
				if _, ok := seen["up-right"]; !ok {
					seen["up-right"] = data[pos+upDelta+rightDelta]
				}
			}
		}

		// down-left
		if !oobD && !oobL {
			if occupied(pos+downDelta+leftDelta) > 0 {
				if _, ok := seen["down-left"]; !ok {
					seen["down-left"] = data[pos+downDelta+leftDelta]
				}
			}
		}

		if !oobD && !oobR {
			// down-right
			if occupied(pos+downDelta+rightDelta) > 0 {
				if _, ok := seen["down-right"]; !ok {
					seen["down-right"] = data[pos+downDelta+rightDelta]
				}
			}
		}

		if false {
			fmt.Println("----------------------------------------------")
			fmt.Println(seen)
			fmt.Println(pos)
			fmt.Println("right delta", rightDelta)
			fmt.Println("left delta", leftDelta)
			fmt.Println("up delta", upDelta)
			fmt.Println("down delta", downDelta)
			fmt.Println("************************************************")
		}

		if oobR && oobL && oobU && oobD {
			break
		}

	}

	switch data[pos] {
	case 'L':
		for _, v := range seen {
			if v == '#' {
				return '0', false
			}
		}
		return '#', true
	case '#':
		c := 0
		for _, v := range seen {
			if v == '#' {
				c += 1
			}
		}
		if c >= 5 {
			return 'L', true
		}
	}
	return '0', false
}

func occupied(pos int) int {
	if pos < 0 || pos >= len(data) {
		return 0
	}

	switch data[pos] {
	case '#', 'L':
		return 1
	}
	return 0
}

func state(data []byte) []byte {
	s := sha1.New()
	s.Write(data)
	return s.Sum(nil)
}
