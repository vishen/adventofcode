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
	//go:embed sample2
	sample2 string
	//go:embed input
	input string
)

func main() {
	part1(sample)
	part1(input)
	// part2("123", 10)
	part2(sample2, 2000)
	part2(input, 2000)
}

func part1(data string) {
	total := 0
	max_iter := 2000

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		num, _ := strconv.Atoi(line)
		for i := 0; i < max_iter; i++ {
			n1 := num * 64
			num = (num ^ n1) % 16777216
			n2 := num / 32
			num = (num ^ n2) % 16777216
			n3 := num * 2048
			num = (num ^ n3) % 16777216
		}
		total += num
	}
	fmt.Println(total)
}

func part2(data string, max_iter int) {
	total := 0
	totals := map[string]int{}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		num, _ := strconv.Atoi(line)
		diff := make([]string, max_iter)
		secrets := make([]string, max_iter)
		for i := 0; i < max_iter; i++ {
			secrets[i] = strconv.Itoa(lastDigit(num))
			prev := num
			n1 := num * 64
			num = (num ^ n1) % 16777216
			n2 := num / 32
			num = (num ^ n2) % 16777216
			n3 := num * 2048
			num = (num ^ n3) % 16777216
			diff[i] = strconv.Itoa(lastDigit(num) - lastDigit(prev))
		}

		bests := map[string]int{}
		for i := 0; i < max_iter-4; i++ {
			key := strings.Join(diff[i:i+4], ".")
			if _, ok := bests[key]; ok {
				continue
			}
			di, _ := strconv.Atoi(secrets[i+4])
			bests[key] += di
		}

		for key, best := range bests {
			totals[key] += best
			if b := totals[key]; b > total {
				total = b
			}
		}
	}
	fmt.Println(total)
}

func lastDigit(n int) int {
	sn := strconv.Itoa(n)
	v, _ := strconv.Atoi(string(sn[len(sn)-1]))
	return v
}
