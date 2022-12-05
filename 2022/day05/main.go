package main

import (
	"fmt"
	"log"
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
	stacks := map[int][]string{}

	readingStacks := true
	for _, line := range strings.Split(input1, "\n") {
		if len(line) == 0 {
			readingStacks = false
			continue
		}
		if readingStacks {
			stack := 1
			i := 0
			for {
				if len(line) <= i {
					break
				}
				if line[i] == '[' {
					stacks[stack] = append(stacks[stack], string(line[i+1]))
				}
				i += 4
				stack += 1
			}
		} else {
			var amount, from, to int
			if _, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to); err != nil {
				log.Fatal(err)
			}
			fromStack := stacks[from]
			for i := 0; i < amount; i++ {
				stacks[to] = append([]string{fromStack[i]}, stacks[to]...)
			}
			stacks[from] = fromStack[amount:]
		}
	}
	for i := 0; i < len(stacks); i++ {
		fmt.Print(stacks[i+1][0])
	}
	fmt.Println()
}

func part2() {
	stacks := map[int][]string{}

	readingStacks := true
	for _, line := range strings.Split(input2, "\n") {
		if len(line) == 0 {
			readingStacks = false
			continue
		}
		if readingStacks {
			stack := 1
			i := 0
			for {
				if len(line) <= i {
					break
				}
				if line[i] == '[' {
					stacks[stack] = append(stacks[stack], string(line[i+1]))
				}
				i += 4
				stack += 1
			}
		} else {
			var amount, from, to int
			if _, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to); err != nil {
				log.Fatal(err)
			}
			fromStack := stacks[from][:amount]
			for i := 0; i < amount; i++ {
				stacks[to] = append([]string{fromStack[len(fromStack)-1-i]}, stacks[to]...)
			}
			stacks[from] = stacks[from][amount:]
		}
	}
	for i := 0; i < len(stacks); i++ {
		fmt.Print(stacks[i+1][0])
	}
	fmt.Println()
}
