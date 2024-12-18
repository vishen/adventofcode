package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed sample
	sample string

	//go:embed input
	input string
)

func main() {
	part1(sample, 6, 6, 12)
	part1(input, 70, 70, 1024)
	part2(sample, 6, 6, 12)
	part2(input, 70, 70, 1024)
}

type XY struct {
	x, y int
}

func part1(data string, X, Y, MAX int) {
	grid := map[XY]bool{}
	for i, line := range strings.Split(strings.TrimSpace(data), "\n") {
		if i == MAX {
			break
		}
		xy := nums(strings.Split(line, ","))
		grid[XY{xy[0], xy[1]}] = true
	}

	/*
		for y := 0; y <= Y; y++ {
			for x := 0; x <= X; x++ {
				if _, ok := grid[XY{x, y}]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	*/

	DIRS := [4]XY{
		XY{0, 1},
		XY{0, -1},
		XY{1, 0},
		XY{-1, 0},
	}

	start := XY{0, 0}
	end := XY{X, Y}

	type job struct {
		pos   XY
		score int
	}

	Q := []job{
		job{start, 0},
	}

	seen := map[XY]bool{}
	for {
		if len(Q) == 0 {
			break
		}

		sort.Slice(Q, func(i, j int) bool {
			return Q[i].score < Q[j].score
		})

		j := Q[0]
		Q = Q[1:]

		if _, ok := seen[j.pos]; ok {
			continue
		}
		seen[j.pos] = true

		if j.pos == end {
			fmt.Println(j.score)
			break
		}

		for _, d := range DIRS {
			dxy := XY{j.pos.x + d.x, j.pos.y + d.y}
			if _, ok := grid[dxy]; ok {
				continue
			}
			if dxy.x < 0 || dxy.x > X || dxy.y < 0 || dxy.y > Y {
				continue
			}
			Q = append(Q, job{dxy, j.score + 1})
		}
	}
}

func part2(data string, X, Y, start int) {
	grid := map[XY]bool{}
	for i, line := range strings.Split(strings.TrimSpace(data), "\n") {
		xy := nums(strings.Split(line, ","))
		grid[XY{xy[0], xy[1]}] = true
		if i >= start+1 {
			if !run(grid, X, Y) {
				fmt.Println(i, line)
				break
			}
		}
	}
}

func run(grid map[XY]bool, X, Y int) bool {
	DIRS := [4]XY{
		XY{0, 1},
		XY{0, -1},
		XY{1, 0},
		XY{-1, 0},
	}

	start := XY{0, 0}
	end := XY{X, Y}

	type job struct {
		pos   XY
		score int
	}

	Q := []job{
		job{start, 0},
	}

	seen := map[XY]bool{}
	for {
		if len(Q) == 0 {
			break
		}

		sort.Slice(Q, func(i, j int) bool {
			return Q[i].score < Q[j].score
		})

		j := Q[0]
		Q = Q[1:]

		if _, ok := seen[j.pos]; ok {
			continue
		}
		seen[j.pos] = true

		if j.pos == end {
			return true
		}

		for _, d := range DIRS {
			dxy := XY{j.pos.x + d.x, j.pos.y + d.y}
			if _, ok := grid[dxy]; ok {
				continue
			}
			if dxy.x < 0 || dxy.x > X || dxy.y < 0 || dxy.y > Y {
				continue
			}
			Q = append(Q, job{dxy, j.score + 1})
		}
	}
	return false
}

func num(n string) int {
	i, _ := strconv.Atoi(n)
	return i
}

func nums(ns []string) []int {
	var is []int
	for _, n := range ns {
		is = append(is, num(n))
	}
	return is
}
