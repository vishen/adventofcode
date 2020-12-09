package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers = append(numbers, convertToInt(line))
	}

	invalidNum := 0

	preamble := 25
	for i := preamble; i < len(numbers); i++ {
		if !valid(numbers[i-preamble:i], numbers[i]) {
			invalidNum = numbers[i]
			fmt.Println("Found invalid", i, numbers[i])
			break
		}
	}

	for i, n1 := range numbers {
		sum := n1
		min := n1
		max := n1
		for j, n2 := range numbers[i+1:] {
			sum += n2
			if n2 > max {
				max = n2
			}
			if n2 < min {
				min = n2
			}
			if sum == invalidNum {
				fmt.Println("Found sequence", i, i+j, sum, invalidNum)
				fmt.Println("Encryption weakness", min+max, min, max)
				return
			}
		}
	}
}

func valid(numbers []int, val int) bool {
	for i, n1 := range numbers {
		for _, n2 := range numbers[i:] {
			if n1+n2 == val {
				return true
			}
		}
	}

	return false
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
