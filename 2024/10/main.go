package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed sample
	sample string
	//go:embed input1
	input1 string
)

func main() {
	part1(sample)
	part1(input1)
	part2(sample)
	part2(input1)
}

func part1(data string) {
	m := strings.Split(strings.TrimSpace(data), "\n")

	Y := len(m)
	X := len(m[0])

	type job struct {
		startY, startX int
		y, x           int
	}

	queue := []job{}

	for y, line := range m {
		for x, ch := range line {
			if ch == '0' {
				queue = append(queue, job{y, x, y, x})
			}
		}
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	found := map[job]int{}
	total := 0
	for {
		if len(queue) == 0 {
			break
		}
		q := queue[0]
		queue = queue[1:]
		if m[q.y][q.x] == '9' {
			found[q]++
			total++
			continue
		}
		for _, d := range dirs {
			y := q.y + d[0]
			x := q.x + d[1]
			if x >= 0 && x < X && y >= 0 && y < Y {
				if m[y][x] == m[q.y][q.x]+1 {
					queue = append(queue, job{
						startY: q.startY,
						startX: q.startX,
						y:      y,
						x:      x,
					})
				}
			}
		}
	}
	fmt.Println(len(found))
}

func part2(data string) {
	m := strings.Split(strings.TrimSpace(data), "\n")

	Y := len(m)
	X := len(m[0])

	type job struct {
		startY, startX int
		y, x           int
	}

	queue := []job{}

	for y, line := range m {
		for x, ch := range line {
			if ch == '0' {
				queue = append(queue, job{y, x, y, x})
			}
		}
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	found := map[job]int{}
	total := 0
	for {
		if len(queue) == 0 {
			break
		}
		q := queue[0]
		queue = queue[1:]
		if m[q.y][q.x] == '9' {
			found[q]++
			total++
			continue
		}
		for _, d := range dirs {
			y := q.y + d[0]
			x := q.x + d[1]
			if x >= 0 && x < X && y >= 0 && y < Y {
				if m[y][x] == m[q.y][q.x]+1 {
					queue = append(queue, job{
						startY: q.startY,
						startX: q.startX,
						y:      y,
						x:      x,
					})
				}
			}
		}
	}
	fmt.Println(total)
}
