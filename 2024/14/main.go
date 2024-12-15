package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed sample
	sample string

	//go:embed input1
	input string
)

func main() {
	part1(input)
	part2(input)
}

type robot struct {
	x, y   int
	vx, vy int
}

func part1(data string) {
	m := map[robot]bool{}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		pv := strings.Split(line, " ")
		p := nums(strings.Split(strings.Split(pv[0], "=")[1], ","))
		v := nums(strings.Split(strings.Split(pv[1], "=")[1], ","))
		m[robot{p[0], p[1], v[0], v[1]}] = true
	}

	X := 101
	Y := 103

	// X = 11
	// Y = 7

	for i := 0; i < 100; i++ {
		nm := map[robot]bool{}
		for r := range m {
			nr := r
			nr.x = (nr.x + nr.vx) % X
			if nr.x < 0 {
				nr.x += X
			}
			nr.y = (nr.y + nr.vy) % Y
			if nr.y < 0 {
				nr.y += Y
			}
			nm[nr] = true
		}
		m = nm
	}
	qs := [4]int{}
	for r := range m {
		if r.x == X/2 || r.y == Y/2 {
			continue
		}

		if r.x <= X/2 && r.y <= Y/2 {
			qs[0]++
		} else if r.x >= X/2 && r.y <= Y/2 {
			qs[1]++
		} else if r.x <= X/2 && r.y >= Y/2 {
			qs[2]++
		} else if r.x >= X/2 && r.y >= Y/2 {
			qs[3]++
		}
	}

	total := 1
	for _, q := range qs {
		total *= q
	}
	fmt.Println(total)
}

func nums(nums []string) []int {
	var ret []int
	for _, n := range nums {
		ni, _ := strconv.Atoi(n)
		ret = append(ret, ni)
	}
	return ret
}

func part2(data string) {
	m := map[robot]bool{}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		pv := strings.Split(line, " ")
		p := nums(strings.Split(strings.Split(pv[0], "=")[1], ","))
		v := nums(strings.Split(strings.Split(pv[1], "=")[1], ","))
		m[robot{p[0], p[1], v[0], v[1]}] = true
	}

	X := 101
	Y := 103

	// X = 11
	// Y = 7

	for i := 0; i < 10000; i++ {
		pos := map[robot]bool{}

		nm := map[robot]bool{}
		for r := range m {
			nr := r
			nr.x = (nr.x + nr.vx) % X
			if nr.x < 0 {
				nr.x += X
			}
			nr.y = (nr.y + nr.vy) % Y
			if nr.y < 0 {
				nr.y += Y
			}
			nm[nr] = true
			pos[robot{nr.x, nr.y, 0, 0}] = true
		}
		m = nm
		possibleTree := false
		for y := 0; y < Y; y++ {
			t := 0
			for x := 0; x < X; x++ {
				if _, ok := pos[robot{x, y, 0, 0}]; ok {
					t += 1
				} else {
					if t > 10 {
						possibleTree = true
						break
					}
					t = 0
				}
			}
			if possibleTree {
				break
			}
		}
		if !possibleTree {
			continue
		}

		fmt.Println("##########################")
		fmt.Println(i + 1)
		for y := 0; y < Y; y++ {
			for x := 0; x < X; x++ {
				found := 0
				for r := range m {
					if r.y == y && r.x == x {
						found += 1
					}
				}
				if found == 0 {
					fmt.Print(".")
				} else {
					fmt.Printf("%d", found)
				}
			}
			fmt.Println()
		}
		fmt.Println("-----------------------------")
	}
}
