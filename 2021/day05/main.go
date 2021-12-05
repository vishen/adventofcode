package main

import (
	"fmt"
	"log"
	"strings"

	_ "embed"
)

var (
	//go:embed test.txt
	test string

	//go:embed part1.txt
	p1 string

	//go:embed part2.txt
	p2 string
)

func main() {
	part1(test)
	part1(p1)
	part2(test)
	part2(p2)
}

func part1(data string) {
	m := map[string]int{}

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			break
		}
		x1, y1, x2, y2 := 0, 0, 0, 0
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatalf("Unable to scan line %q: %v", line, err)
		}

		if x1 == x2 {
			low, high := lowHigh(y1, y2)
			for i := low; i <= high; i++ {
				m[key(x1, i)]++
			}
		} else if y1 == y2 {
			low, high := lowHigh(x1, x2)
			for i := low; i <= high; i++ {
				m[key(i, y1)]++
			}
		}
	}

	overlap := 0
	for _, v := range m {
		if v >= 2 {
			overlap++
		}
	}
	fmt.Printf("Part 1: %d overlap more than twice\n", overlap)
}

func part2(data string) {
	m := map[string]int{}

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			break
		}
		x1, y1, x2, y2 := 0, 0, 0, 0
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatalf("Unable to scan line %q: %v", line, err)
		}

		if x1 == x2 {
			low, high := lowHigh(y1, y2)
			for i := low; i <= high; i++ {
				m[key(x1, i)]++
			}
		} else if y1 == y2 {
			low, high := lowHigh(x1, x2)
			for i := low; i <= high; i++ {
				m[key(i, y1)]++
			}
		} else {
			xlow, xhigh, xdir := lowHighDir(x1, x2)
			_, _, ydir := lowHighDir(y1, y2)
			// 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3
			// - xlow=1 xhigh=3, ylow=1 yhigh=3

			// 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9
			// - xlow=7 xhigh=9, ylow=7 yhigh=9

			for i := 0; i <= xhigh-xlow; i++ {
				k := key(x1+(i*xdir), y1+(i*ydir))
				m[k]++
			}
		}
	}

	overlap := 0
	for _, v := range m {
		if v >= 2 {
			overlap++
		}
	}
	fmt.Printf("Part 2: %d overlap more than twice\n", overlap)
}

func lowHigh(v1, v2 int) (int, int) {
	if v1 > v2 {
		return v2, v1
	}
	return v1, v2
}

func lowHighDir(v1, v2 int) (int, int, int) {
	if v1 > v2 {
		return v2, v1, -1
	}
	return v1, v2, 1
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
