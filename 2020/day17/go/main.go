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

type point struct {
	x, y, z, w int
}

func run(data []byte) {
	state := map[point]bool{}

	size := 0
	y, z, w := 0, 0, 0
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
				state[point{i, y, z, w}] = true
			}
		}
		y += 1
	}

	steps := 6
	for i := 0; i < steps; i++ {
		state = step(state)
	}

	active := 0
	for _, a := range state {
		if a {
			active += 1
		}
	}
	fmt.Println(active)
}

func step(state map[point]bool) map[point]bool {
	newState := map[point]bool{}

	// Handle active state
	for p, ok := range state {
		if !ok {
			continue
		}
		activeNextTo := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						p1 := point{p.x + x, p.y + y, p.z + z, p.w + w}
						if _, ok := state[p1]; ok {
							activeNextTo += 1
						}
					}
				}
			}
		}
		if activeNextTo == 2 || activeNextTo == 3 {
			newState[p] = true
		}
	}

	// Handle inactive state
	inactive := []point{}
	for p, ok := range state {
		if !ok {
			continue
		}
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						p1 := point{p.x + x, p.y + y, p.z + z, p.w + w}
						if set := state[p1]; !set {
							inactive = append(inactive, p1)
						}
					}
				}
			}
		}
	}

	for _, i := range inactive {
		adjacent := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						p1 := point{i.x + x, i.y + y, i.z + z, i.w + w}
						if _, ok := state[p1]; ok {
							adjacent += 1
						}
					}
				}
			}
		}

		if adjacent == 3 {
			newState[i] = true
		}
	}

	return newState
}
