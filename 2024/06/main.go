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
	lines := strings.Split(strings.TrimSpace(data), "\n")

	Y := len(lines)
	X := len(lines[0])

	dir := "UP"
	dirY, dirX := -1, 0

	startX, startY := 0, 0
	for y, line := range lines {
		for x, ch := range line {
			if ch == '^' {
				startX = x
				startY = y
			}
		}
	}

	type key struct {
		x, y int
	}
	seen := map[key]bool{}

	curX, curY := startX, startY
	for {
		seen[key{curY, curX}] = true

		curY += dirY
		curX += dirX

		if curY < 0 || curY >= Y {
			break
		}
		if curX < 0 || curX >= X {
			break
		}

		if lines[curY][curX] == '#' {
			curY -= dirY
			curX -= dirX
			switch dir {
			case "UP":
				dir = "RIGHT"
				dirY = 0
				dirX = 1
			case "RIGHT":
				dir = "DOWN"
				dirY = 1
				dirX = 0
			case "DOWN":
				dir = "LEFT"
				dirY = 0
				dirX = -1
			case "LEFT":
				dir = "UP"
				dirY = -1
				dirX = 0
			}
		}

	}
	fmt.Println(len(seen))
}

func part2(data string) {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	startX, startY := 0, 0
	for y, line := range lines {
		for x, ch := range line {
			if ch == '^' {
				startX = x
				startY = y
			}
		}
	}

	loops := 0
	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				continue
			}
			nlines := make([]string, len(lines))
			copy(nlines, lines)
			nl := []byte(nlines[y])
			nl[x] = '#'
			nlines[y] = string(nl)
			if run(startY, startX, nlines) {
				loops++
			}
		}
	}
	fmt.Println(loops)
}

func run(startY, startX int, lines []string) bool {

	Y := len(lines)
	X := len(lines[0])

	dir := "UP"
	dirY, dirX := -1, 0

	type key struct {
		x, y int
		dir  string
	}
	seen := map[key]bool{}

	curX, curY := startX, startY
	for {
		k := key{curY, curX, dir}
		if _, ok := seen[k]; ok {
			return true
		}
		seen[k] = true

		curY += dirY
		curX += dirX

		if curY < 0 || curY >= Y {
			break
		}
		if curX < 0 || curX >= X {
			break
		}

		if lines[curY][curX] == '#' {
			curY -= dirY
			curX -= dirX
			switch dir {
			case "UP":
				dir = "RIGHT"
				dirY = 0
				dirX = 1
			case "RIGHT":
				dir = "DOWN"
				dirY = 1
				dirX = 0
			case "DOWN":
				dir = "LEFT"
				dirY = 0
				dirX = -1
			case "LEFT":
				dir = "UP"
				dirY = -1
				dirX = 0
			}
		}
	}
	return false
}
