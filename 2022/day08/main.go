package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input1.txt
var input1 string

//go:embed input1.txt
var input2 string

func main() {
	part1()
	part2()
}

func key(c, r int) string {
	return fmt.Sprintf("%d.%d", c, r)
}

func part1() {
	trees := map[string]int{}

	columns := 0
	rows := 0

	for column, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}
		for i, c := range line {
			trees[key(column, i)] = int(c - '0')
		}
		columns += 1
		rows = len(line)
	}

	isVisibleLeft := func(c, r int) bool {
		t := trees[key(c, r)]
		for i := c - 1; i >= 0; i-- {
			if trees[key(i, r)] >= t {
				return false
			}
		}
		return true
	}
	isVisibleRight := func(c, r int) bool {
		t := trees[key(c, r)]
		for i := c + 1; i <= columns; i++ {
			if trees[key(i, r)] >= t {
				return false
			}
		}
		return true
	}
	isVisibleUp := func(c, r int) bool {
		t := trees[key(c, r)]
		for i := r - 1; i >= 0; i-- {
			if trees[key(c, i)] >= t {
				return false
			}
		}
		return true
	}
	isVisibleDown := func(c, r int) bool {
		t := trees[key(c, r)]
		for i := r + 1; i <= rows; i++ {
			if trees[key(c, i)] >= t {
				return false
			}
		}
		return true
	}

	isVisibleFuncs := []func(int, int) bool{
		isVisibleDown,
		isVisibleRight,
		isVisibleLeft,
		isVisibleUp,
	}

	visible := rows*2 + ((columns - 2) * 2)
	for r := 1; r < rows-1; r++ {
		for c := 1; c < columns-1; c++ {
			for fni, fn := range isVisibleFuncs {
				_ = fni
				if fn(c, r) {
					// fmt.Println(c, r, key(c, r), trees[key(c, r)], fni)
					visible += 1
					break
				}
			}
		}
	}
	fmt.Println(visible)
}

func part2() {
	trees := map[string]int{}

	columns := 0
	rows := 0

	for column, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}
		for i, c := range line {
			trees[key(column, i)] = int(c - '0')
		}
		columns += 1
		rows = len(line)
	}

	best := 0
	for r := 1; r < rows-1; r++ {
		for c := 1; c < columns-1; c++ {
			score := 1

			t := trees[key(c, r)]
			s := 0
			for i := c - 1; i >= 0; i-- {
				s += 1
				if trees[key(i, r)] >= t {
					break
				}
			}
			score *= s
			s = 0
			for i := c + 1; i < columns; i++ {
				s += 1
				if trees[key(i, r)] >= t {
					break
				}
			}
			score *= s
			s = 0
			for i := r - 1; i >= 0; i-- {
				s += 1
				if trees[key(c, i)] >= t {
					break
				}
			}
			score *= s
			s = 0
			for i := r + 1; i < rows; i++ {
				s += 1
				if trees[key(c, i)] >= t {
					break
				}
			}
			score *= s

			if score > best {
				best = score
			}
		}
	}
	fmt.Println(best)
}
