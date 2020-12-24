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
	total := flippedTiles(tiles)
	fmt.Println("Flipped:", total)

	// Part 2
	days := 100
	for i := 0; i < days; i++ {
		tiles = update(tiles)
		total := flippedTiles(tiles)
		fmt.Println("Day", i+1, ": ", total)
	}
}

func flippedTiles(tiles map[dir]bool) int {
	total := 0
	for _, flipped := range tiles {
		if flipped {
			total += 1
		}
	}
	return total
}

func update(tiles map[dir]bool) map[dir]bool {
	newTiles := map[dir]bool{}

	// Handle black tiles
	for tile, ok := range tiles {
		if !ok {
			continue
		}
		adjacent := 0
		for _, d := range knownDirections {
			if f := tiles[dir{tile.row + d.row, tile.col + d.col}]; f {
				adjacent += 1
			}
		}

		if adjacent > 0 && adjacent <= 2 {
			newTiles[tile] = true
		}
	}

	// Handle white tiles
	whiteTiles := []dir{}
	for tile, ok := range tiles {
		if !ok {
			continue
		}
		for _, d := range knownDirections {
			dir := dir{tile.row + d.row, tile.col + d.col}
			f := tiles[dir]
			if !f {
				whiteTiles = append(whiteTiles, dir)
			}
		}
	}

	for _, tile := range whiteTiles {
		adjacent := 0
		for _, d := range knownDirections {
			if f := tiles[dir{tile.row + d.row, tile.col + d.col}]; f {
				adjacent += 1
			}
		}

		if adjacent == 2 {
			newTiles[tile] = true
		}
	}

	return newTiles
}

// Naive approach
func updateNaive(tiles map[dir]bool) map[dir]bool {
	newTiles := map[dir]bool{}
	size := len(tiles)
	for i := -size; i < size; i++ {
		for j := -size; j < size; j++ {
			flipped := tiles[dir{i, j}]

			adjacent := 0
			for _, d := range knownDirections {
				if f := tiles[dir{i + d.row, j + d.col}]; f {
					adjacent += 1
				}
			}

			if flipped {
				if adjacent > 0 && adjacent <= 2 {
					newTiles[dir{i, j}] = true
				}
			} else {
				if adjacent == 2 {
					newTiles[dir{i, j}] = true
				}
			}
		}
	}
	return newTiles
}
