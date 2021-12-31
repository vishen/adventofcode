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

var rooms = map[byte]int{
	'A': 3,
	'B': 5,
	'C': 7,
	'D': 9,
}

func main() {
	//part2(test2)
	// part2(test1)
	part2(input1)
}

type point struct {
	y, x int
}

func part2(data []byte) {

	area := map[point]byte{}
	height := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		for i, c := range line {
			area[point{height, i}] = c
		}
		height++
	}

	leastEnergy := find(area)
	fmt.Println("Part 2:", leastEnergy)
}

func move(m map[point]byte, opt, npt point, c byte) map[point]byte {
	n := map[point]byte{}
	for k, v := range m {
		n[k] = v
	}
	n[opt] = '.'
	n[npt] = c
	return n
}

func diff(p1, p2 point) int {
	return absMinus(p1.y, 1) + absMinus(p2.y, 1) + absMinus(p1.x, p2.x)
}

func val(c byte) int {
	switch c {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}
	panic("unknown val")
}

/*
func key(area map[point]byte) string {
	k := ""
	for y := 0; y < 7; y++ {
		for x := 0; x < 13; x++ {
			pt := point{y, x}
			c := area[pt]
			switch c {
			case 'A', 'B', 'C', 'D':
				// expected
			default:
				continue
			}
			k += fmt.Sprintf("%c.%d.%d,", c, pt.y, pt.x)
		}
	}
	return k
}
*/

func key(v version) string {
	k := fmt.Sprintf("%d:", v.energy)
	for y := 0; y < 7; y++ {
		for x := 0; x < 13; x++ {
			pt := point{y, x}
			c := v.area[pt]
			switch c {
			case 'A', 'B', 'C', 'D':
				// expected
			default:
				continue
			}
			k += fmt.Sprintf("%c.%d.%d,", c, pt.y, pt.x)
		}
	}
	return k
}

type version struct {
	area   map[point]byte
	energy int
}

func find(area map[point]byte) int {
	type nextMove struct {
		c     byte
		pt    point
		moves []point
	}

	uniqs := map[string]struct{}{}
	versions := []version{{area, 0}}
	lowest := 1_000_000_000
	for v := 0; ; v++ {
		sort.Slice(versions, func(i, j int) bool {
			return versions[i].energy > versions[j].energy
		})
		if v%10000 == 0 {
			fmt.Println(v, len(versions))
		}
		if len(versions) == 0 {
			return lowest
		}
		v := versions[0]
		versions = versions[1:]

		area := v.area
		k := key(v)
		if _, ok := uniqs[k]; ok {
			// fmt.Println("cache hit")
			continue
		}
		uniqs[k] = struct{}{}

		allMoves := []nextMove{}
		for pt, c := range area {
			if c == '#' || c == '.' {
				continue
			}
			moves := bestMoves(area, pt)
			if len(moves) > 0 {
				allMoves = append(allMoves, nextMove{c, pt, moves})
			}
		}
		sort.Slice(allMoves, func(i, j int) bool {
			return len(allMoves[i].moves) < len(allMoves[j].moves)
		})
		for _, nm := range allMoves {
			for _, npt := range nm.moves {
				narea := move(area, nm.pt, npt, nm.c)
				d := diff(nm.pt, npt)
				cur := (d * val(nm.c))
				spent := cur + v.energy
				if completed(narea) {
					fmt.Println("Completed:", spent)
					if lowest > spent {
						lowest = spent
					}
					break
				}
				versions = append(versions, version{narea, spent})
			}
		}
	}
	return lowest
}

func printArea(area map[point]byte) {
	for y := 0; y < 7; y++ {
		for x := 0; x < 13; x++ {
			fmt.Print(string(area[point{y, x}]))
		}
		fmt.Println()
	}
}

func completed(area map[point]byte) bool {
	for c, x := range rooms {
		for i := 0; i < 4; i++ {
			pt := point{i + 2, x}
			if area[pt] != c {
				return false
			}
		}
	}
	return true
}

func absMinus(x1, x2 int) int {
	if x1 > x2 {
		return x1 - x2
	}
	return x2 - x1
}

func bestMoves(area map[point]byte, pt point) []point {
	c := area[pt]
	// Check to see if the point should be moved at all, ie:
	// if it is already in homw
	allowedRoomX := rooms[c]
	if pt.x == allowedRoomX {
		allowed := true
		inRoom := false
		for i := 0; i < 4; i++ {
			if pt.y == i+2 {
				inRoom = true
			}
			pt := point{i + 2, allowedRoomX}
			if area[pt] != c {
				allowed = false
				break
			}
		}

		if allowed && inRoom {
			return []point{}
		}
	}
	moves := make(map[point]bool, len(area))
	generateMoves(c, area, pt, moves)

	nmoves := make([]point, 0, len(moves))
	for npt, ok := range moves {
		if pt.y == 1 && npt.y == 1 {
			continue
		}
		if ok {
			nmoves = append(nmoves, npt)
		}
	}
	sort.Slice(nmoves, func(i, j int) bool {
		if nmoves[i].y != nmoves[j].y {
			return nmoves[i].y > nmoves[j].y
		}
		return absMinus(pt.x, nmoves[i].x) < absMinus(pt.x, nmoves[j].x)
	})

	// Short-circuit when a pt is already in the best spot or
	// there is a spot in the ideal home room.
	if len(nmoves) > 0 {
		if nmoves[0].x == pt.x {
			return []point{}
		} else if nmoves[0].x == rooms[c] {
			return []point{nmoves[0]}
		}
	}

	xs := map[int]struct{}{}
	nnmoves := make([]point, 0, len(nmoves))
	for _, pt := range nmoves {
		if _, ok := xs[pt.x]; !ok {
			nnmoves = append(nnmoves, pt)
			xs[pt.x] = struct{}{}
		}
	}
	return nnmoves
}

func generateMoves(c byte, area map[point]byte, pt point, moves map[point]bool) {

	dirs := [][2]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}

	for _, dir := range dirs {
		npt := point{pt.y + dir[0], pt.x + dir[1]}
		if _, ok := moves[npt]; ok {
			continue
		}
		valid := isValid(area, npt, c)
		if valid {
			moves[npt] = true
		} else if area[npt] == '.' {
			// do nothnig
			moves[npt] = false
		} else {
			continue
		}
		generateMoves(c, area, npt, moves)
	}
}

func isValid(area map[point]byte, npt point, c byte) bool {
	// Check the new point is possible to move to
	if area[npt] != '.' {
		return false
	}

	// Check to make sure the new points isn't a hallway
	{
		hallways := []point{
			{1, 3},
			{1, 5},
			{1, 7},
			{1, 9},
		}

		for _, hw := range hallways {
			if npt.y == hw.y && npt.x == hw.x {
				return false
			}
		}
	}

	// Check for points in home room
	allowedRoomX := rooms[c]
	if npt.x == allowedRoomX {
		allowed := true
		for i := 0; i < 4; i++ {
			pt := point{i + 2, allowedRoomX}
			if area[pt] == '.' || area[pt] == c {
				// valid case
			} else {
				allowed = false
				break
			}
		}

		if !allowed {
			return false
		}

	} else if npt.y != 1 {
		return false
	}

	return true

}
