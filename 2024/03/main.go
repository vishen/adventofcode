package main

import (
	"fmt"
	"strings"

	_ "embed"
)

var (
	//go:embed sample
	sample string

	//go:embed sample2
	sample2 string

	//go:embed input1
	input1 string
)

func main() {
	//part1(sample)
	part1(input1)
	// part2(sample2)
	part2(input1)
}

func eatNumber(data string) (int, int) {
	val := 0
	cur := 0
	for i, d := range data {
		if d >= '0' && d <= '9' {
			val *= 10
			val += int(d - '0')
		} else {
			cur = i
			break
		}
	}
	return val, cur
}

func part1(data string) {
	total := 0
	muls := strings.Split(data, "mul")
	for _, m := range muls {
		cur := 0
		if m[0] != '(' {
			continue
		}
		cur++
		num1, ncur := eatNumber(m[cur:])
		cur += ncur
		if m[cur] != ',' {
			continue
		}
		cur++
		num2, ncur := eatNumber(m[cur:])
		cur += ncur
		if m[cur] != ')' {
			continue
		}
		total += num1 * num2
	}
	fmt.Println(total)
}

func part2(data string) {

	total := 0
	muls := strings.Split(data, "mul")
	enabled := true
	for i, m := range muls {
		if i > 0 {
			for mi, mo := range muls[i-1] {
				if mo != 'd' {
					continue
				}

				if muls[i-1][mi:mi+4] == "do()" {
					enabled = true
				} else if mi+7 <= len(muls[i-1]) && muls[i-1][mi:mi+7] == "don't()" {
					enabled = false
				}
			}
		}

		cur := 0
		if m[0] != '(' {
			continue
		}
		cur++
		num1, ncur := eatNumber(m[cur:])
		cur += ncur
		if m[cur] != ',' {
			continue
		}
		cur++
		num2, ncur := eatNumber(m[cur:])
		cur += ncur
		if m[cur] != ')' {
			continue
		}
		if enabled {
			total += num1 * num2
		}
	}
	fmt.Println(total)
}
