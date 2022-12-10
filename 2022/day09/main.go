package main

import (
	"fmt"
	"strconv"
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

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func key(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

func part1() {
	visited := map[string]struct{}{
		key(0, 0): struct{}{},
	}

	headX, headY := 0, 0
	pheadX, pheadY := 0, 0
	tailX, tailY := 0, 0

	move := func(dx, dy int) {
		pheadX = headX
		pheadY = headY
		headX += dx
		headY += dy
		if abs(headX, tailX) >= 2 || abs(headY, tailY) >= 2 {
			tailX = pheadX
			tailY = pheadY
			visited[key(tailX, tailY)] = struct{}{}
		}
	}

	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}
		cmds := strings.Split(line, " ")
		dir := cmds[0]
		amount, _ := strconv.Atoi(cmds[1])
		for i := 0; i < amount; i++ {
			switch dir {
			case "L":
				move(-1, 0)
			case "R":
				move(1, 0)
			case "U":
				move(0, -1)
			case "D":
				move(0, 1)
			}
		}
	}
	fmt.Println(len(visited))
}

func diff(a, b int) int {
	switch abs(a, b) {
	case 1:
		return a - b
	case 2:
		return (a - b) / 2
	}
	return 0
}

func part2() {
	visited := map[string]struct{}{
		key(0, 0): struct{}{},
	}

	rope := [10]struct {
		x, y int
	}{}

	move := func(dx, dy int) {
		rope[0].x += dx
		rope[0].y += dy

		for i := 1; i < 10; i++ {
			pr := rope[i-1]
			cr := rope[i]
			if abs(pr.x, cr.x) >= 2 || abs(pr.y, cr.y) >= 2 {

				cr.x += diff(pr.x, cr.x)
				cr.y += diff(pr.y, cr.y)
				rope[i] = cr

				if i == 9 {
					visited[key(cr.x, cr.y)] = struct{}{}
				}
			}
		}
	}

	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}
		cmds := strings.Split(line, " ")
		dir := cmds[0]
		amount, _ := strconv.Atoi(cmds[1])
		for i := 0; i < amount; i++ {
			switch dir {
			case "L":
				move(-1, 0)
			case "R":
				move(1, 0)
			case "U":
				move(0, -1)
			case "D":
				move(0, 1)
			}
		}
	}
	fmt.Println(len(visited))
}
