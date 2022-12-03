package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input1.txt
var input1 string

//go:embed input1.txt
var input2 string

func main() {
	// part1()
	part2()
}

var scoring = map[string]int{
	"X": 1, // "A" -- Rock
	"Y": 2, // "B" -- Paper
	"Z": 3, // "C" -- Scissors
}

var scoring2 = map[string]int{
	"A": 1, // "A" -- Rock
	"B": 2, // "B" -- Paper
	"C": 3, // "C" -- Scissors
}

func result(a, b string) int {
	switch {
	case a == "A" && b == "Y":
		return 6
	case a == "A" && b == "Z":
		return 0
	case a == "B" && b == "X":
		return 0
	case a == "B" && b == "Z":
		return 6
	case a == "C" && b == "X":
		return 6
	case a == "C" && b == "Y":
		return 0
	}
	return 3
}

func part1() {
	score := 0
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}
		lineS := strings.Split(line, " ")
		score += (scoring[lineS[1]] + result(lineS[0], lineS[1]))
		fmt.Println(lineS, score)
	}
	fmt.Println(score)
}

func result2(a, b string) int {
	/*"Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"*/
	if b == "Y" {
		return 3 + scoring2[a]
	}
	switch {
	case a == "A" && b == "X":
		return scoring2["C"]
	case a == "A" && b == "Z":
		return scoring2["B"] + 6
	case a == "B" && b == "X":
		return scoring2["A"]
	case a == "B" && b == "Z":
		return scoring2["C"] + 6
	case a == "C" && b == "X":
		return scoring2["B"]
	case a == "C" && b == "Z":
		return scoring2["A"] + 6
	}
	panic("Aaa")
}

func part2() {
	score := 0
	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}
		lineS := strings.Split(line, " ")
		score += result2(lineS[0], lineS[1])
		fmt.Println(lineS, score)
	}
	fmt.Println(score)
}
