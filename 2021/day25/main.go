package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	part1(test1)
	part1(input1)
}

func part1(data []byte) {
	sea := map[[2]int]byte{}
	height := 0
	width := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		for i, c := range line {
			if c != '.' {
				sea[[2]int{height, i}] = c
			}
		}
		width = len(line)
		height++
	}
	s := 0
	for ; ; s++ {
		changed := false
		nsea := map[[2]int]byte{}
		for pt, c := range sea {
			if c != '>' {
				continue
			}

			npt := pt
			npt[1]++
			if npt[1] == width {
				npt[1] = 0
			}

			if _, ok := sea[npt]; !ok {
				nsea[npt] = c
				changed = true
			} else {
				nsea[pt] = c
			}
		}

		for pt, c := range sea {
			if c != 'v' {
				continue
			}

			npt := pt
			npt[0]++
			if npt[0] == height {
				npt[0] = 0
			}

			if _, ok := nsea[npt]; ok {
				nsea[pt] = c
				continue
			}

			if oc, ok := sea[npt]; ok && oc == c {
				nsea[pt] = c
			} else {
				nsea[npt] = c
				changed = true
			}
		}

		sea = nsea

		if !changed {
			break
		}

	}
	fmt.Printf("Part 1: %d steps\n", s+1)
}

func printSea(sea map[[2]int]byte, height, width int) {
	fmt.Println("--------------------------------")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if c, ok := sea[[2]int{y, x}]; ok {
				fmt.Printf("%c", c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------")
}
