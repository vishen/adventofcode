package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {

	data, err := ioutil.ReadFile("./input2.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

var (
	tileIDMap = map[string][]string{}
)

type state int

const (
	state_tile_id state = iota
	state_tiles
)

func run(data []byte) {

	tiles := []tile{}

	s := state_tile_id

	var curTileID int
	var curTile tile
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		switch s {
		case state_tile_id:
			if bytes.HasPrefix(line, []byte("Tile ")) {
				curTileID = convertToInt(line[5 : len(line)-1])
				curTile = tile{id: curTileID}
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

	sideHash := map[string][]string{}

	for _, t := range tiles {
		// Top row
		toHash := make([]byte, t.size)
		for i := 0; i < t.size; i++ {
			toHash[i] = t.data[i]
		}
		h := hash(toHash)
		sideHash[h] = append(sideHash[h], fmt.Sprintf("%d.top", t.id))

		// Bottom row
		toHash = make([]byte, t.size)
		for i := 0; i < t.size; i++ {
			j := (len(t.data) - t.size) + i
			toHash[i] = t.data[j]
		}
		h = hash(toHash)
		sideHash[h] = append(sideHash[h], fmt.Sprintf("%d.bottom", t.id))

		// Left column
		toHash = make([]byte, t.size)
		for i := 0; i < t.size; i++ {
			j := (t.size * i)
			toHash[i] = t.data[j]
		}
		h = hash(toHash)
		sideHash[h] = append(sideHash[h], fmt.Sprintf("%d.left", t.id))

		// Right column
		toHash = make([]byte, t.size)
		for i := 0; i < t.size; i++ {
			j := (t.size * i) + (t.size - 1)
			toHash[i] = t.data[j]
		}
		h = hash(toHash)
		sideHash[h] = append(sideHash[h], fmt.Sprintf("%d.right", t.id))
	}

	for h, tileIDs := range sideHash {
		fmt.Println(h, len(tileIDs), tileIDs)
		for _, tileID := range tileIDs {
			tileIDMap[tileID] = append(tileIDMap[tileID], tileIDs...)
		}
	}

	fmt.Println()
	for tileID, tileIDs := range tileIDMap {
		fmt.Println(tileID, tileIDs)
	}
}

func hash(data []byte) string {
	h := 1
	mid := len(data) / 2
	for i, c := range data {
		b := 3
		if c == '#' {
			b = 5
		}

		if i/mid == 0 {
			h += b * i
		} else {
			h += b * ((len(data) - 1) - i)
		}
	}
	return strconv.Itoa(h)
}

type tile struct {
	id   int
	data []byte
	size int
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
