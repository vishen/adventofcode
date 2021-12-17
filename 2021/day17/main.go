package main

import (
	"bytes"
	"fmt"

	_ "embed"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	part1(test1)
	part1(input1)
}

func part1(data []byte) {
	lineS := bytes.Split(data, []byte{'='})
	x1, i := readNum(lineS[1])
	x2, i := readNum(lineS[1][i:])
	y1, i := readNum(lineS[2])
	y2, i := readNum(lineS[2][i:])

	fmt.Println(y1, y2, x1, x2)

	var found [][2]int
	maxY := 0
	for vy := -1000; vy < 1000; vy++ {
		for vx := -1000; vx < 1000; vx++ {
			probe := [2]int{0, 0}
			velocity := [2]int{vy, vx}
			my := 0
			for s := 0; s < 1000; s++ {
				probe[1] += velocity[1]
				probe[0] += velocity[0]
				if probe[0] > my {
					my = probe[0]
				}
				if velocity[1] > 0 {
					velocity[1]--
				} else if velocity[1] < 0 {
					velocity[1]++
				}
				velocity[0]--
				if x1 <= probe[1] && probe[1] <= x2 && y1 <= probe[0] && probe[0] <= y2 {
					found = append(found, [2]int{vy, vx})
					if maxY < my {
						maxY = my
					}
					break
				}
			}
		}
	}
	fmt.Printf("Part 1: highest y %d\n", maxY)

	uniq := map[[2]int]struct{}{}
	for _, f := range found {
		uniq[f] = struct{}{}
	}
	fmt.Println(len(found), len(uniq))
}

func readNum(data []byte) (int, int) {
	neg := false
	val := 0
	end := 0
	for i, c := range data {
		if c == '-' {
			neg = true
		} else if c >= '0' && c <= '9' {
			val *= 10
			val += int(c - '0')
		} else if val > 0 {
			end = i
			break
		}
	}
	if neg {
		val *= -1
	}
	return val, end
}
