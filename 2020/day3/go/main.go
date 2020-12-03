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

	path := bytes.Split(data, []byte{'\n'})

	pathsToTravel := [][2]int{
		// Right, Down
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	mul := 1
	for _, t := range pathsToTravel {
		hit := traversePath(path, t[0], t[1])
		fmt.Println(t, hit)
		mul *= hit
	}
	fmt.Println(mul)
}

func traversePath(path [][]byte, colTravel, rowTravel int) int {

	colLen := len(path[0])
	treesHit := 0

	row := 0
	col := 0
	for {
		row += rowTravel

		if row >= len(path) || len(path[row]) == 0 {
			break
		}

		col += colTravel
		if col >= colLen {
			col -= colLen
		}

		switch path[row][col] {
		case '#':
			treesHit += 1
		}

		if row == len(path)-1 {
			break
		}
	}

	return treesHit
}
