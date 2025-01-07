package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strings"
)

var (
	//go:embed sample
	sample string
	//go:embed sample2
	sample2 string
	//go:embed input
	input string
)

func main() {
	//part1(sample2)
	// part1(sample)
	// part1(input)
	// part2(sample, 2, 100)
	// part2(sample, 20, 50)
	part2(input, 2, 100)
	part2(sample, 20, 50)
	part2(input, 20, 100)
}

type pos struct {
	x, y int
}

func part2(data string, cheats, threshold int) {
	grid := map[pos]rune{}

	var start, end pos
	for y, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for x, ch := range line {
			switch ch {
			case 'S':
				start = pos{x, y}
				ch = '.'
			case 'E':
				end = pos{x, y}
				ch = '.'
			}

			grid[pos{x, y}] = ch
		}
	}

	dirs := [4]pos{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	type car struct {
		pos         pos
		dir         int
		picoseconds int
	}

	var q []car
	for d := range dirs {
		q = append(q, car{start, d, 0})
	}

	add := func(c car, d int) {
		nc := car{c.pos, d, c.picoseconds}
		q = append(q, nc)
	}

	type key struct {
		pos pos
		dir int
	}
	seen := map[key]bool{}

	path := []pos{
		start,
	}
	for {
		if len(q) == 0 {
			break
		}

		c := q[0]
		q = q[1:]

		c.pos.x += dirs[c.dir].x
		c.pos.y += dirs[c.dir].y
		c.picoseconds++

		if c.pos == end {
			path = append(path, c.pos)
			break
		}
		if ch, ok := grid[c.pos]; !ok {
			continue
		} else if ch == '#' {
			continue
		}

		k := key{c.pos, 0}
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = true
		path = append(path, c.pos)

		add(c, c.dir)
		switch c.dir {
		case 0, 1:
			add(c, 2)
			add(c, 3)
		case 2, 3:
			add(c, 0)
			add(c, 1)
		}
	}

	found := 0
	for i, p1 := range path {
		for j, p2 := range path[i+1:] {
			diff := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			if diff <= cheats {
				if (j+1)-diff >= threshold {
					found += 1
				}
			}
		}
	}
	fmt.Println(found)
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func part1(data string) {
	grid := map[pos]rune{}

	var start, end pos
	for y, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for x, ch := range line {
			switch ch {
			case 'S':
				start = pos{x, y}
				ch = '.'
			case 'E':
				end = pos{x, y}
				ch = '.'
			}

			grid[pos{x, y}] = ch
		}
	}

	initial := run(grid, start, end, 0, 0)
	if len(initial) != 1 {
		log.Panicf("expected only one finding without cheats, got %v", initial)
	}
	withoutCheats := 0
	for k := range initial {
		withoutCheats = k
	}
	fmt.Printf("without cheats %d\n", withoutCheats)

	type s struct {
		saved, count int
	}

	withCheats := run(grid, start, end, 1, withoutCheats)
	var sorted []s
	for taken, count := range withCheats {
		saved := withoutCheats - taken
		if saved > 0 {
			sorted = append(sorted, s{saved, count})
		}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].saved < sorted[j].saved
	})
	total := 0
	for _, s := range sorted {
		fmt.Printf("%d cheats saved %d picoseconds\n", s.count, s.saved)
		if s.saved >= 100 {
			total += s.count
		}
	}
	fmt.Println(total)
}

func run(grid map[pos]rune, start, end pos, cheats, max int) map[int]int {
	dirs := [4]pos{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	type car struct {
		pos         pos
		dir         int
		picoseconds int
		cheats      int
		path        string
	}

	path := func(c *car) {
		// c.path += fmt.Sprintf("{%v, %d, %d, %d}..", c.pos, c.dir, c.cheats, c.picoseconds)
	}

	var q []car
	for d := range dirs {
		nc := car{start, d, 0, cheats, ""}
		path(&nc)
		q = append(q, nc)

	}

	add := func(c car, d int) {
		nc := car{c.pos, d, c.picoseconds, c.cheats, c.path}
		path(&nc)
		q = append(q, nc)
	}

	found := map[int]int{}
	for {
		if len(q) == 0 {
			break
		}

		c := q[0]
		q = q[1:]

		if max > 0 && c.picoseconds > max {
			continue
		}

		c.pos.x += dirs[c.dir].x
		c.pos.y += dirs[c.dir].y
		c.picoseconds++

		if c.pos == end {
			found[c.picoseconds]++
			continue
		}
		ch, ok := grid[c.pos]
		if !ok {
			continue
		}

		cheating := false
		if ch == '#' {
			if c.cheats <= 0 {
				continue
			}
			c.cheats -= 1
			cheating = true
		} else {
			if c.cheats < cheats {
				c.cheats -= 1
			}
		}

		if cheating {
			for d := range dirs {
				add(c, d)
			}
		} else {
			add(c, c.dir)
			switch c.dir {
			case 0, 1:
				add(c, 2)
				add(c, 3)
			case 2, 3:
				add(c, 0)
				add(c, 1)
			}
		}

	}
	return found
}
