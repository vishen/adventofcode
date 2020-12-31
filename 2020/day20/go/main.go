package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
)

func main() {

	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

var (
	tiles = []*tile{}
	edges = map[uint][]*tile{}
)

type state int

const (
	state_tile_id state = iota
	state_tiles
)

func run(data []byte) {

	s := state_tile_id

	var curTileID int
	var curTile *tile
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		switch s {
		case state_tile_id:
			if bytes.HasPrefix(line, []byte("Tile ")) {
				curTileID = convertToInt(line[5 : len(line)-1])
				curTile = &tile{id: curTileID}
				s = state_tiles
				continue
			}
		case state_tiles:
			if len(line) == 0 {
				s = state_tile_id
				tiles = append(tiles, curTile)
				continue
			}
			curTile.data = append(curTile.data, line...)
			curTile.size = len(line)
		}
	}

	addEdge := func(e uint, t *tile) {
		edges[e] = append(edges[e], t)
		re := flipped(e)
		if e != re {
			edges[re] = append(edges[re], t)
		}
	}

	for _, t := range tiles {
		// Top row
		for i := 0; i < t.size; i++ {
			t.n |= shifted(t.data[i], i)
		}
		addEdge(t.n, t)

		// Bottom row
		for i := 0; i < t.size; i++ {
			j := (len(t.data) - t.size) + i
			t.s |= shifted(t.data[j], i)
		}
		addEdge(t.s, t)

		// Left column
		for i := 0; i < t.size; i++ {
			j := (t.size * i)
			t.e |= shifted(t.data[j], i)
		}
		addEdge(t.e, t)

		// Right column
		for i := 0; i < t.size; i++ {
			j := (t.size * i) + (t.size - 1)
			t.w |= shifted(t.data[j], i)
		}
		addEdge(t.w, t)
	}

	part1()
}

func part1() {
	total := 1
	for _, t := range tiles {
		neighbours := map[int]bool{}
		for _, dir := range []uint{t.n, t.s, t.e, t.w} {
			for _, t1 := range edges[dir] {
				if t.id != t1.id {
					neighbours[t1.id] = true
				}
			}
			for _, t1 := range edges[flipped(dir)] {
				if t.id != t1.id {
					neighbours[t1.id] = true
				}
			}
		}
		if len(neighbours) == 2 {
			total *= t.id
		}
	}
	fmt.Println("Part 1:", total)
}

func flipped(val uint) uint {
	return bits.Reverse(val) >> 54
}

func shifted(c byte, i int) uint {
	if c == '#' {
		return 1 << (9 - i)
	}
	return 0
}

type tile struct {
	id   int
	data []byte
	size int

	n uint
	e uint
	s uint
	w uint
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
