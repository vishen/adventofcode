package main

import (
	"fmt"
	"strings"

	_ "embed"
)

var (
	//go:embed test1.txt
	test1 string

	//go:embed test2.txt
	test2 string

	//go:embed input1.txt
	input1 string
)

func main() {
	part1(test1)
	part1(test2)
	part1(input1)
	part2(test1)
	part2(input1)
}

type room struct {
	name       string
	neighbours []string
}

func part1(data string) {

	rooms := map[string]*room{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		rs := strings.Split(line, "-")
		r1, r2 := rs[0], rs[1]
		{
			r, ok := rooms[r2]
			if !ok {
				r = &room{name: r2}
				rooms[r2] = r
			}
			r.neighbours = append(r.neighbours, r1)
		}

		r, ok := rooms[r1]
		if !ok {
			r = &room{name: r1}
			rooms[r1] = r
		}
		r.neighbours = append(r.neighbours, r2)
	}

	total := visit("start", rooms, nil)
	fmt.Printf("Part 1: %d paths\n", total)
}

func visit(room string, rooms map[string]*room, visited []string) int {
	if room == "end" {
		return 1
	}

	for _, v := range visited {
		if room == v {
			if v[0] >= 'a' && v[0] <= 'z' {
				return 0
			}
		}
	}

	v := append(visited, room)

	r := rooms[room]
	total := 0
	for _, n := range r.neighbours {
		if n == "start" || n == "" {
			continue
		}
		total += visit(n, rooms, v)
	}
	return total
}

func part2(data string) {
	rooms := map[string]*room{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		rs := strings.Split(line, "-")
		r1, r2 := rs[0], rs[1]

		{
			r, ok := rooms[r2]
			if !ok {
				r = &room{name: r2}
				rooms[r2] = r
			}
			r.neighbours = append(r.neighbours, r1)
		}

		r, ok := rooms[r1]
		if !ok {
			r = &room{name: r1}
			rooms[r1] = r
		}
		r.neighbours = append(r.neighbours, r2)
	}

	total := visit2("start", rooms, nil)
	fmt.Printf("Part 2: %d paths\n", total)
}

func visit2(room string, rooms map[string]*room, visited []string) int {
	if room == "end" {
		return 1
	}

	m := map[string]bool{}

	hasBeenVisited := false
	smallCaveVisitedMoreThanOnce := false

	for _, v := range visited {
		if v[0] >= 'a' && v[0] <= 'z' {
			if room == v {
				hasBeenVisited = true
			}
			if _, ok := m[v]; ok {
				smallCaveVisitedMoreThanOnce = true
			}
		}
		m[v] = true
	}

	if hasBeenVisited && smallCaveVisitedMoreThanOnce {
		return 0
	}

	v := append(visited, room)

	total := 0

	r := rooms[room]
	for _, n := range r.neighbours {
		if n == "start" || n == "" {
			continue
		}
		total += visit2(n, rooms, v)
	}
	return total
}
