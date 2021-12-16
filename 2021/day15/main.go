package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed test2.txt
	test2 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	// part1(test2)
	part1(test1)
	part1(input1)
	part2(input1)
}

func part1(data []byte) {
	c := &cavern{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		c.cave = append(c.cave, line)
		c.height++
	}
	c.width = len(c.cave[0])

	lowest := c.paths()
	fmt.Printf("Part 1: lowest path %d\n", lowest)
}

func part2(data []byte) {
	c := &cavern{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		c.cave = append(c.cave, line)
		c.height++
	}
	c.width = len(c.cave[0])

	lowest := c.paths2()
	fmt.Printf("Part 2: lowest path %d\n", lowest)
}

type point struct {
	x, y int
}

type cavern struct {
	cave   [][]byte
	width  int
	height int
}

func (c *cavern) paths() int {
	seen := map[point]int{}

	type qj struct {
		p     point
		total int
	}

	q := []qj{qj{}}
	done := false
	for {
		if len(q) == 0 || done {
			break
		}

		nq := []qj{}
		for _, j := range q {
			x, y := j.p.x, j.p.y
			if y < 0 || y >= c.height {
				continue
			}
			if x < 0 || x >= c.width {
				continue
			}

			curVal := int(c.cave[y][x] - '0')
			cost := j.total + curVal
			if dist, ok := seen[j.p]; ok {
				if dist > cost {
					seen[j.p] = cost
				} else {
					continue
				}
			} else {
				seen[j.p] = cost
			}

			if x == c.width && y == c.height {
				done = true
				break
			}
			dirs := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			}
			for _, dir := range dirs {
				nq = append(nq, qj{
					p: point{
						x: x + dir[1],
						y: y + dir[0],
					},
					total: cost,
				})
			}
		}
		q = nq
		sort.Slice(q, func(i, j int) bool {
			return q[i].total < q[j].total
		})
	}
	return seen[point{y: c.height - 1, x: c.width - 1}] - int(c.cave[0][0]-'0')
}

func (c *cavern) paths2() int {
	seen := map[point]int{}

	type qj struct {
		p     point
		total int
	}

	q := []qj{qj{}}
	done := false
	for {
		if len(q) == 0 || done {
			break
		}

		nq := []qj{}
		for _, j := range q {
			x, y := j.p.x, j.p.y
			if y < 0 || y >= c.height*5 {
				continue
			}
			if x < 0 || x >= c.width*5 {
				continue
			}

			ny := y % c.height
			nx := x % c.width
			curVal := (int(c.cave[ny][nx]-'0') + y/c.height + x/c.width)
			cost := j.total + (curVal % 10)
			if dist, ok := seen[j.p]; ok {
				if dist > cost {
					seen[j.p] = cost
				} else {
					continue
				}
			} else {
				seen[j.p] = cost
			}

			if x == c.width*5 && y == c.height*5 {
				done = true
				break
			}
			dirs := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			}
			for _, dir := range dirs {
				nq = append(nq, qj{
					p: point{
						x: x + dir[1],
						y: y + dir[0],
					},
					total: cost,
				})
			}
		}
		q = nq
		sort.Slice(q, func(i, j int) bool {
			return q[i].total < q[j].total
		})
	}
	return seen[point{y: (c.height * 5) - 1, x: (c.width * 5) - 1}] - int(c.cave[0][0]-'0')
}
