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

func part1() {
	x := 1
	cycles := 0
	ss := 0
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}

		cmds := strings.Split(line, " ")

		oc := cycles
		ox := x
		switch cmds[0] {
		case "noop":
			cycles += 1
		case "addx":
			cycles += 2
			val, _ := strconv.Atoi(cmds[1])
			x += val
		}

		for _, v := range []int{20, 60, 100, 140, 180, 220} {
			if oc < v && cycles >= v {
				ss += (ox * v)
			}
		}
	}
	fmt.Println(ss)
}

func part2() {
	x := 1
	cycles := 0
	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}

		cmds := strings.Split(line, " ")

		oc := cycles
		ox := x
		switch cmds[0] {
		case "noop":
			cycles += 1
		case "addx":
			cycles += 2
			val, _ := strconv.Atoi(cmds[1])
			x += val
		}
		for i := oc; i < cycles; i++ {
			if i > 0 && i%40 == 0 {
				fmt.Println()
				ox += 40
				x += 40
			}
			switch i {
			case ox - 1, ox, ox + 1:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
	}
}
