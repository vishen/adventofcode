package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var (
	//go:embed test.txt
	test []byte

	//go:embed input1.txt
	input1 []byte

	directions = []int{-1, 0, 1}
)

func main() {
	part1(test)
	part1(input1)
	part2(test)
	part2(input1)
}

func part1(data []byte) {
	ndata := make([]byte, len(data))
	copy(ndata, data)

	g := &grid{}
	for _, line := range bytes.Split(ndata, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		g.board = append(g.board, line)
		g.size++
	}
	steps := 100
	for s := 0; s < steps; s++ {
		for y := 0; y < g.size; y++ {
			for x := 0; x < g.size; x++ {
				g.step(y, x)
			}
		}
		g.resolve()
	}
	fmt.Printf("Part 1: %d flashes\n", g.flashes)
}

func part2(data []byte) {
	ndata := make([]byte, len(data))
	copy(ndata, data)

	g := &grid{}
	for _, line := range bytes.Split(ndata, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		g.board = append(g.board, line)
		g.size++
	}
	s := 0
	for ; ; s++ {
		for y := 0; y < g.size; y++ {
			for x := 0; x < g.size; x++ {
				g.step(y, x)
			}
		}
		if g.resolve() {
			break
		}
	}
	fmt.Printf("Part 2: %d steps\n", s+1)
}

type grid struct {
	size    int
	board   [][]byte
	flashes int
}

func (g *grid) resolve() bool {
	found := 0
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			if g.board[y][x] == '*' {
				g.board[y][x] = '0'
				g.flashes++
				found++
			}
		}
	}
	return found == g.size*g.size
}

func (g *grid) step(y, x int) {
	if g.board[y][x] == '*' {
		return
	}
	if g.board[y][x] != '9' {
		g.board[y][x]++
		return
	}

	g.board[y][x] = '*'

	for _, dirX := range directions {
		for _, dirY := range directions {
			if dirX == 0 && dirY == 0 {
				continue
			}
			if y+dirY < 0 || y+dirY >= g.size || x+dirX < 0 || x+dirX >= g.size {
				continue
			}
			g.step(y+dirY, x+dirX)
		}
	}
}

func (g *grid) printBoard() {
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			fmt.Printf("%c", g.board[y][x])
		}
		fmt.Println()
	}
}
