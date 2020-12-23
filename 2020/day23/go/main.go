package main

import (
	"fmt"
	"strconv"
)

func main() {
	//input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	input := []int{9, 6, 2, 7, 1, 3, 8, 5, 4}
	run(input)
}

func run(input []int) {
	turns := 100
	for i := 0; i < turns; i++ {

		iw := i % len(input)
		max := iw + len(input)
		inputExt := append(input, input...)
		cur := input[iw]

		fmt.Printf("Turn %d: %v\n", i+1, input)
		fmt.Println("destination:", cur)
		fmt.Println("pick up:", inputExt[1+iw:4+iw])

		hold := map[int]bool{}
		for _, n := range inputExt[1+iw : 4+iw] {
			hold[n] = true
		}

		num := previous(cur)
		for i := 0; i < len(hold); i++ {
			if _, ok := hold[num]; !ok {
				break
			}
			num = previous(num)
		}

		index := 0
		for j, n := range inputExt[iw+4 : max] {
			if n == num {
				index = j
				break
			}
		}

		newInput := make([]int, len(input))
		nj := 0
		for j := 0; j < iw+1; j++ {
			newInput[nj] = input[j]
			nj++
		}
		for j := iw + 4; j < iw+5+index; j++ {
			jw := nj % len(input)
			newInput[jw] = inputExt[j]
			nj++
		}
		for j := iw + 1; j < iw+4; j++ {
			jw := nj % len(input)
			newInput[jw] = inputExt[j]
			nj++
		}
		for j := iw + 5 + index; j < len(input); j++ {
			jw := nj % len(input)
			newInput[jw] = inputExt[j]
			nj++
		}
		fmt.Println("num:", num, index)
		input = newInput[:len(input)]
		fmt.Println(input)
	}

	// Part 1
	output1 := ""
	output2 := ""
	rev := false
	for _, num := range input {
		if num == 1 {
			rev = true
			continue
		}
		if !rev {
			output1 += strconv.Itoa(num)
		} else {
			output2 += strconv.Itoa(num)
		}
	}
	fmt.Println(output2 + output1)
}

func previous(cur int) int {
	// This is always the case in all games
	max := 9
	min := 1

	cur -= 1
	if cur < min {
		return max
	}
	return cur
}
