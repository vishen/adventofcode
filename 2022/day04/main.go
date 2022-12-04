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
	overlaps := 0

	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			continue
		}

		sectionParts := strings.Split(line, ",")

		var sections [][2]int
		for i := 0; i < 2; i++ {
			sectionIDs := strings.Split(sectionParts[i], "-")
			id1, _ := strconv.Atoi(sectionIDs[0])
			id2, _ := strconv.Atoi(sectionIDs[1])
			sections = append(sections, [2]int{id1, id2})
		}

		if sections[0][0] >= sections[1][0] && sections[1][1] >= sections[0][1] {
			// fmt.Println(sections[0], "IN", sections[1])
			overlaps += 1
		} else if sections[1][0] >= sections[0][0] && sections[0][1] >= sections[1][1] {
			// fmt.Println(sections[1], "IN", sections[0])
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}

func part2() {
	overlaps := 0

	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			continue
		}

		sectionParts := strings.Split(line, ",")

		var sections [][2]int
		for i := 0; i < 2; i++ {
			sectionIDs := strings.Split(sectionParts[i], "-")
			id1, _ := strconv.Atoi(sectionIDs[0])
			id2, _ := strconv.Atoi(sectionIDs[1])
			sections = append(sections, [2]int{id1, id2})
		}

		if sections[0][0] >= sections[1][0] && sections[0][0] <= sections[1][1] {
			//fmt.Println(sections[0], "IN1", sections[1])
			overlaps += 1
		} else if sections[0][1] >= sections[1][0] && sections[0][1] <= sections[1][1] {
			//fmt.Println(sections[0], "IN2", sections[1])
			overlaps += 1
		} else if sections[1][0] >= sections[0][0] && sections[1][0] <= sections[0][1] {
			//fmt.Println(sections[0], "IN3", sections[1])
			overlaps += 1
		} else if sections[1][1] >= sections[0][0] && sections[1][1] <= sections[0][1] {
			//fmt.Println(sections[0], "IN4", sections[1])
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}

func assert(expr bool, msg string, args ...interface{}) {
	if !expr {
		panic(fmt.Sprintf(msg, args...))
	}
}
