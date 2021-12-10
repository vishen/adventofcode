package main

import (
	"bytes"
	"fmt"
	"sort"

	_ "embed"
)

var (
	//go:embed test.txt
	test []byte

	//go:embed test1.txt
	test1 []byte

	//go:embed test2.txt
	test2 []byte

	//go:embed part1.txt
	p1 []byte
)

func main() {
	part1(test)
	part1(p1)
	part2(test)
	part2(p1)
}

func part1(data []byte) {
	width, height := 0, 0
	var heatmap [][]byte

	for i, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		heatmap = append(heatmap, line)

		if i == 0 {
			width = len(line)
		}
		height++
	}

	lowPoints := []byte{}
	for ri, row := range heatmap {
		for pi, point := range row {

			adjacentLowPoints := 0

			// Check up
			if ri == 0 || heatmap[ri-1][pi] > point {
				adjacentLowPoints++
			}

			// Check down
			if ri == height-1 || heatmap[ri+1][pi] > point {
				adjacentLowPoints++
			}

			// Check left
			if pi == 0 || heatmap[ri][pi-1] > point {
				adjacentLowPoints++
			}

			// Check right
			if pi == width-1 || heatmap[ri][pi+1] > point {
				adjacentLowPoints++
			}

			if adjacentLowPoints == 4 {
				lowPoints = append(lowPoints, point)
				// fmt.Printf("LOW POINT: ri=%d pi=%d point=%c\n", ri, pi, point)
			}
		}
	}

	total := 0
	for _, p := range lowPoints {
		total += int((p - '0')) + 1
	}

	fmt.Printf("Part 1: %d adjacent points for total of %d\n", len(lowPoints), total)

}

type basin struct {
	clumps map[int][][2]int
}

type clump struct {
	pos        [2]int
	neighbours []*clump

	visited bool
}

func part2(data []byte) {
	found := map[int][]*clump{}
	lines := 0
	for li, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		start := -1
		for i, point := range line {
			if start != -1 && point == '9' {
				found[li] = append(found[li], &clump{pos: [2]int{start, i - 1}})
				start = -1
			} else if start == -1 && point != '9' {
				start = i
			}
		}
		lines++
		if start != -1 {
			found[li] = append(found[li], &clump{pos: [2]int{start, len(line) - 1}})
			start = -1
		}
	}

	totals := []int{}
	for row, clumps := range found {
		for _, clump := range clumps {
			if clump.visited == true {
				continue
			}
			total := visit(found, row, clump)
			totals = append(totals, total)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	total := 1
	for _, t := range totals[:3] {
		total *= t
	}
	fmt.Printf("Part 2: largest=%v sum=%d\n", totals[:3], total)
}

func visit(found map[int][]*clump, row int, clump *clump) int {
	if clump.visited {
		return 0
	}
	total := (clump.pos[1] + 1) - clump.pos[0]
	clump.visited = true
	for _, dir := range []int{1, -1} {
		for _, c := range found[row+dir] {
			if inClump(clump.pos, c.pos) {
				total += visit(found, row+dir, c)
			}
		}
	}
	return total
}

func inClump(a, b [2]int) bool {
	// Dumb, but can't work out elegant solution
	for ai := a[0]; ai <= a[1]; ai++ {
		for bi := b[0]; bi <= b[1]; bi++ {
			if ai == bi {
				return true
			}
		}
	}
	return false
}
