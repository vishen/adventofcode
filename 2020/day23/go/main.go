package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//input := []int{9, 6, 2, 7, 1, 3, 8, 5, 4}
	run(input)
}

//const size = 1_000_000
const size = 10

func run(inputStart []int) {

	input := make([]int, size-1)
	copy(input, inputStart)

	for i := 10; i < size; i++ {
		input[i-1] = i
	}

	fmt.Println("Finished setting up input")

	window := 10

	turns := 2
	for i := 0; i < turns; i++ {

		iw := i % len(input)
		inputExt := make([]int, 0, window)
		copy(inputExt, input[:window])
		cur := input[iw]

		fmt.Printf("Turn %d: %v\n", i+1, input)
		fmt.Println("destination:", cur)
		fmt.Println("pick up:", inputExt[1+iw:4+iw])

		hold := map[int]bool{}
		for _, n := range inputExt[1+iw : 4+iw] {
			hold[n] = true
		}

		num := previous(cur, window-1)
		for i := 0; i < len(hold); i++ {
			if _, ok := hold[num]; !ok {
				break
			}
			num = previous(num, window-1)
		}

		index := 0
		for j, n := range inputExt[iw+4 : window] {
			if n == num {
				index = j
				break
			}
		}

		newInput := make([]int, len(inputExt))
		nj := 0
		for j := 0; j < iw+1; j++ {
			newInput[nj] = inputExt[j]
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
		for _, i := range newInput {
			input[i] = newInput[i]
		}
		window++
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
	min := 1

	cur -= 1
	if cur < min {
		return size
	}
	return cur
}
