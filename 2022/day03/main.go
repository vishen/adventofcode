package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input1.txt
var input1 string

//go:embed test1.txt
var input2 string

func main() {
	//part1()
	part2()
}

func assert(expr bool, msg string, args ...interface{}) {
	if !expr {
		panic(fmt.Sprintf(msg, args...))
	}
}

func priority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - 96
	}
	return int(c) - 38
}

func part1() {
	score := 0
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}
		assert(len(line)&1 == 0, "line is not even: %d", len(line))

		h := len(line) / 2

		box1 := make(map[rune]struct{}, h)
		for _, x := range line[:h] {
			box1[x] = struct{}{}
		}

		for _, x := range line[h:] {
			if _, ok := box1[x]; ok {
				// fmt.Printf("%c, %d, %d\n", x, x, priority(x))
				score += priority(x)
				break
			}
		}
	}
	fmt.Println(score)
}

func part2() {
	score := 0
	var boxes []string
	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}

		boxes = append(boxes, line)
		if len(boxes) < 3 {
			continue
		}

		box1 := map[rune]struct{}{}
		for _, x := range boxes[0] {
			box1[x] = struct{}{}
		}

		box2 := map[rune]struct{}{}
		for _, x := range boxes[1] {
			box2[x] = struct{}{}
		}

		for _, x := range boxes[2] {
			_, ok1 := box1[x]
			_, ok2 := box2[x]
			if ok1 && ok2 {
				// fmt.Printf("%c, %d, %d\n", x, x, priority(x))
				score += priority(x)
				break
			}
		}
		boxes = []string{}
	}
	fmt.Println(score)
}
