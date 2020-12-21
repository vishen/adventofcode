package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input2.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

type point struct {
	x, y, z int
}

func run(data []byte) {
	state := map[point]bool{}

	size := 0
	y, z := 0, 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		if size == 0 {
			size = len(line)
		}

		// For the initial state, z is always 0
		for i, c := range line {
			if c == '#' {
				state[point{i, y, z}] = true
			}
		}
		y += 1
	}

	steps := 1
	for i := 0; i < steps; i++ {
		size, state = step(size, state)
		fmt.Println(size)
		fmt.Println(state)
	}

	active := 0
	for _, a := range state {
		if a {
			active += 1
		}
	}
	fmt.Println(active)
}

func step(size int, state map[point]bool) (int, map[point]bool) {
	newState := map[point]bool{}
	newSize := size + 2

	// Handle active state
	for p := range state {
		activeNextTo := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					if x == 0 && y == 0 && z == 0 {
						continue
					}
					p1 := point{p.x + x, p.y + y, p.z + z}
					if _, ok := state[p1]; ok {
						activeNextTo += 1
					}
				}
			}
		}
		if activeNextTo == 2 || activeNextTo == 3 {
			newState[p] = true
		}
	}

	// Handle inactive states
	for x := -size; x <= size; x++ {
		for y := -size; y <= size; y++ {
			for z := -size; z <= size; z++ {
				p := point{x, y, z}
				// Ignore active states
				if _, ok := state[p]; ok {
					continue
				}
				activeNextTo := 0
				for x1 := -1; x1 <= 1; x1++ {
					for y1 := -1; y1 <= 1; y1++ {
						for z1 := -1; z1 <= 1; z1++ {
							if x1 == 0 && y1 == 0 && z1 == 0 {
								continue
							}
							p1 := point{x + x1, y + y1, z + z1}
							if _, ok := state[p1]; ok {
								activeNextTo += 1
							}
						}
					}
				}
				if activeNextTo == 3 {
					newState[p] = true
				}
			}
		}
	}

	return newSize, newState
}
