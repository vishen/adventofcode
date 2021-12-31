package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed test2.txt
	test2 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	// part1(input1)
	part22(input1)
	// part21(input1)
}

func part21(data []byte) {

	type step struct {
		cube [6]int
		on   bool
	}

	xs := []int{}
	ys := []int{}
	zs := []int{}
	steps := []step{}

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		on := line[0] == 'o' && line[1] == 'n'

		var xyz [6]int
		pos := 0
		for i := 0; i < 6; i++ {
			val, npos := readNextInt(line[pos:])
			pos += npos
			xyz[i] = val
		}
		xyz[1]++
		xyz[3]++
		xyz[5]++
		xs = append(xs, xyz[0], xyz[1])
		ys = append(ys, xyz[2], xyz[3])
		zs = append(zs, xyz[4], xyz[5])
		steps = append(steps, step{xyz, on})
	}

	deduplicate := func(arr []int) []int {
		narr := make([]int, 0, len(arr))
		used := map[int]struct{}{}
		for _, v := range arr {
			if _, ok := used[v]; ok {
				continue
			}
			narr = append(narr, v)
			used[v] = struct{}{}
		}
		return narr
	}

	xs = deduplicate(xs)
	ys = deduplicate(ys)
	zs = deduplicate(zs)

	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)

	find := func(val int, arr []int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		panic("should not get here")
	}

	grid := map[[3]int]bool{}
	for i, s := range steps {
		fmt.Println(len(steps), i)
		x1 := find(s.cube[0], xs)
		x2 := find(s.cube[1], xs)
		y1 := find(s.cube[2], ys)
		y2 := find(s.cube[3], ys)
		z1 := find(s.cube[4], zs)
		z2 := find(s.cube[5], zs)
		for xi := x1; xi < x2; xi++ {
			for yi := y1; yi < y2; yi++ {
				for zi := z1; zi < z2; zi++ {
					grid[[3]int{xi, yi, zi}] = s.on
				}
			}
		}
	}

	total := uint64(0)
	for xi := 0; xi < len(xs); xi++ {
		for yi := 0; yi < len(ys); yi++ {
			for zi := 0; zi < len(zs); zi++ {
				if on := grid[[3]int{xi, yi, zi}]; on {
					total += uint64((xs[xi+1] - xs[xi]) * (ys[yi+1] - ys[yi]) * (zs[zi+1] - zs[zi]))
				}
			}
		}
	}

	fmt.Printf("Part 2: total=%d\n", total)
}

func part22(data []byte) {

	cubes := [][7]int{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		var xyz [7]int
		pos := 0
		for i := 0; i < 6; i++ {
			val, npos := readNextInt(line[pos:])
			pos += npos
			xyz[i] = val
		}

		for _, c := range cubes {
			in, ok := intersect(c, xyz)
			if !ok {
				continue
			}
			if in[6] == 0 {
				in[6] = 1
			} else {
				in[6] = 0
			}
			cubes = append(cubes, in)
		}

		if on := line[0] == 'o' && line[1] == 'n'; on {
			xyz[6] = 1
			cubes = append(cubes, xyz)
		}
	}

	total := 0
	for _, c := range cubes {
		if c[6] == 1 {
			total += volume(c)
		} else {
			total -= volume(c)
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}

func intersect(c1, c2 [7]int) ([7]int, bool) {
	x1 := max(c1[0], c2[0])
	x2 := min(c1[1], c2[1])
	y1 := max(c1[2], c2[2])
	y2 := min(c1[3], c2[3])
	z1 := max(c1[4], c2[4])
	z2 := min(c1[5], c2[5])

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return [7]int{}, false
	}
	return [7]int{x1, x2, y1, y2, z1, z2, c1[6]}, true
}

func volume(c [7]int) int {
	return (c[1] - c[0] + 1) * (c[3] - c[2] + 1) * (c[5] - c[4] + 1)
}

func part1(data []byte) {
	type cube [3]int
	cubes := map[cube]bool{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		on := line[0] == 'o' && line[1] == 'n'
		var xyz [6]int
		pos := 0
		for i := 0; i < 6; i++ {
			val, npos := readNextInt(line[pos:])
			pos += npos
			xyz[i] = val
		}
		for x := max(xyz[0], -50); x <= min(xyz[1], 50); x++ {
			for y := max(xyz[2], -50); y <= min(xyz[3], 50); y++ {
				for z := max(xyz[4], -50); z <= min(xyz[5], 50); z++ {
					cubes[cube{x, y, z}] = on
				}
			}
		}
	}

	total := 0
	for _, on := range cubes {
		if on {
			total += 1
		}
	}
	fmt.Printf("Part 1 cubes on %d\n", total)
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func readNextInt(val []byte) (int, int) {
	iVal := 0
	i := 0
	found := false
	neg := false
	for ; i < len(val); i++ {
		if val[i] == '-' {
			neg = true
		}
		if val[i] < '0' || val[i] > '9' {
			if found {
				break
			}
			continue
		}
		iVal *= 10
		iVal += int(val[i] - '0')
		found = true
	}
	if neg {
		iVal *= -1
	}
	return iVal, i
}
