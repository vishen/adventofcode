package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

var (
	//go:embed test1.txt
	test1 string

	//go:embed input1.txt
	input1 string
)

func main() {
	part1(test1)
	part1(input1)
	part2(test1)
	part2(input1)
}

func part1(data string) {
	points := [][2]int{}

	foldX := []int{}
	foldY := []int{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		lineS := strings.Split(line, ",")
		if len(lineS) > 1 {
			x := i(lineS[0])
			y := i(lineS[1])
			points = append(points, [2]int{x, y})
			continue
		}

		lineSS := strings.Split(line, "=")
		if lineSS[0][len(lineSS[0])-1] == 'x' {
			foldX = append(foldX, i(lineSS[1]))
		} else {
			foldY = append(foldY, i(lineSS[1]))
		}
	}

	uniq := map[[2]int]struct{}{}
	for _, point := range points {
		if point[0] > foldX[0] {
			diff := point[0] - foldX[0]
			point[0] = (foldX[0] - diff)
		}
		uniq[point] = struct{}{}
	}

	fmt.Printf("Part 1: %d points\n", len(uniq))
}

type point struct {
	x, y int
}

func part2(data string) {
	points := map[point]struct{}{}

	maxX, maxY := 0, 0
	folds := [][2]int{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		lineS := strings.Split(line, ",")
		if len(lineS) > 1 {
			x := i(lineS[0])
			y := i(lineS[1])
			points[point{x, y}] = struct{}{}

			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}
			continue
		}

		lineSS := strings.Split(line, "=")
		fold := [2]int{}
		if lineSS[0][len(lineSS[0])-1] == 'x' {
			fold[0] = i(lineSS[1])
		} else {
			fold[1] = i(lineSS[1])
		}
		folds = append(folds, fold)
	}

	for _, fold := range folds {
		newPoints := map[point]struct{}{}
		for point := range points {
			fx, fy := fold[0], fold[1]
			if fx > 0 && point.x > fx {
				diff := point.x - fx
				point.x = fx - diff
				maxX = fx
			} else if fy > 0 && point.y > fy {
				diff := point.y - fy
				point.y = fy - diff
				maxY = fy
			}
			newPoints[point] = struct{}{}
		}
		points = newPoints
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if _, ok := points[point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func i(val string) int {
	valI, _ := strconv.Atoi(val)
	return valI
}
