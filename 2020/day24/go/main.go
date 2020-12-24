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

type dir struct {
	row, col int
}

// e, se, sw, w, nw, and ne
// https://www.redblobgames.com/grids/hexagons/
var knownDirections = map[string]dir{
	"e":  {0, 1},
	"se": {-1, 0},
	"ne": {1, 1},

	"w":  {0, -1},
	"sw": {-1, -1},
	"nw": {1, 0},
}

func run(data []byte) {

	tiles := map[dir]bool{}

	for linei, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		i := 0
		row, col := 0, 0
		for {
			var (
				dir   dir
				ok    bool
				found string
			)
			if i >= len(line) {
				break
			}
			if len(line) >= i+2 {
				found = string(line[i : i+2])
				dir, ok = knownDirections[found]
				if ok {
					i++
				}
			}
			if !ok {
				found = string(line[i])
				dir, ok = knownDirections[found]
				if !ok {
					panic(fmt.Sprintf("can't find direction for line %d at pos %d: %s or %c not known", linei+1, i, line[i:i+2], line[i]))
				}
			}
			i++
			row += dir.row
			col += dir.col
		}
		h := dir{row, col}
		tiles[h] = !tiles[h]
	}

	// Part 1
	total := 0
	for _, flipped := range tiles {
		if flipped {
			total += 1
		}
	}
	fmt.Println("Flipped:", total)
}
