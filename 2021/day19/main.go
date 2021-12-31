package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte

	allBeacons  = map[[3]int]bool{}
	scannersPos = [][3]int{}
)

func main() {
	// part12(test1)
	part12(input1)
}

type beacon [3]int

type scanner struct {
	name    string
	beacons []beacon
}

func part12(data []byte) {
	var scanners []scanner
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		if line[0] == '-' && line[1] == '-' {
			scanners = append(scanners, scanner{name: fmt.Sprintf("Scanner %d", len(scanners))})
			continue
		}

		vs := bytes.Split(line, []byte{','})
		v1, _ := readNum(vs[0])
		v2, _ := readNum(vs[1])
		v3, _ := readNum(vs[2])
		b := beacon{v1, v2, v3}
		scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, b)

		if len(scanners) == 1 {
			allBeacons[b] = true
		}
	}

	scanners = scanners[1:]
	for {
		if len(scanners) == 0 {
			break
		}

		newScanners := []scanner{}
		for _, s := range scanners {
			if !compare(s) {
				newScanners = append(newScanners, s)
			}
		}
		scanners = newScanners
	}
	fmt.Printf("Part 1: %d beacons\n", len(allBeacons))

	largest := 0
	for i, s := range scannersPos {
		for _, s1 := range scannersPos[i+1:] {
			md := manhatten(s, s1)
			val := md[0] + md[1] + md[2]
			if val > largest {
				largest = val
			}
		}
	}
	fmt.Printf("Part 2: Largest manhatten distance %d\n", largest)
}

func compare(s scanner) bool {
	type possiblePerm struct {
		perm int
		s    [3]int
	}

	possiblePerms := map[possiblePerm]int{}
	for _, sb := range s.beacons {
		for ab := range allBeacons {
			for i, perm := range perms(sb) {
				d1 := diff(ab, perm)
				possiblePerms[possiblePerm{i, d1}]++
			}
		}
	}

	for pp, count := range possiblePerms {
		if count >= 12 {
			scannersPos = append(scannersPos, pp.s)
			for _, pb := range s.beacons {
				rot := perms(pb)[pp.perm]
				d := add(rot, pp.s)
				allBeacons[d] = true
			}
			return true
		}
	}
	return false
}

func manhatten(s1, s2 [3]int) [3]int {
	return [3]int{
		abs(s1[0] - s2[0]),
		abs(s1[1] - s2[1]),
		abs(s1[2] - s2[2]),
	}
}

func diff(b1, b2 beacon) [3]int {
	return [3]int{
		b1[0] - b2[0],
		b1[1] - b2[1],
		b1[2] - b2[2],
	}
}

func add(b1, b2 beacon) [3]int {
	return [3]int{
		b1[0] + b2[0],
		b1[1] + b2[1],
		b1[2] + b2[2],
	}
}

func equals(d1, d2 [3]int) bool {
	used := map[int]struct{}{}
	found := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if _, ok := used[j]; ok {
				continue
			}
			if d1[i] == d2[j] {
				found++
				used[j] = struct{}{}
				break
			}
		}
	}
	return found == 3
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func readNum(data []byte) (int, int) {
	neg := false
	val := 0
	end := 0
	for i, c := range data {
		if c == '-' {
			neg = true
		} else if c >= '0' && c <= '9' {
			val *= 10
			val += int(c - '0')
		} else {
			end = i
			break
		}
	}
	if neg {
		val *= -1
	}
	return val, end
}

func perms(xyz [3]int) [24][3]int {
	allPerms := [24][3]int{}

	// https://imgur.com/Ff1vGT9
	allPerms[0] = [3]int{xyz[0], xyz[1], xyz[2]}   //x,y,z
	allPerms[1] = [3]int{xyz[0], xyz[2], -xyz[1]}  //x,z,-y
	allPerms[2] = [3]int{xyz[0], -xyz[1], -xyz[2]} //x,-y,-z
	allPerms[3] = [3]int{xyz[0], -xyz[2], xyz[1]}  //x,-z,y

	allPerms[4] = [3]int{-xyz[0], -xyz[1], xyz[2]}  //-x,-y,z
	allPerms[5] = [3]int{-xyz[0], xyz[2], xyz[1]}   //-x,z,y
	allPerms[6] = [3]int{-xyz[0], xyz[1], -xyz[2]}  //-x,y,-z
	allPerms[7] = [3]int{-xyz[0], -xyz[2], -xyz[1]} //-x,-z,-y

	allPerms[8] = [3]int{xyz[1], xyz[2], xyz[0]}    //y,z,x
	allPerms[9] = [3]int{xyz[1], xyz[0], -xyz[2]}   //y,x,-z
	allPerms[10] = [3]int{xyz[1], -xyz[2], -xyz[0]} //y,-z,-x
	allPerms[11] = [3]int{xyz[1], -xyz[0], xyz[2]}  //y,-x,z

	allPerms[12] = [3]int{-xyz[1], -xyz[2], xyz[0]}  //-y,-z,x
	allPerms[13] = [3]int{-xyz[1], xyz[0], xyz[2]}   //-y,x,z
	allPerms[14] = [3]int{-xyz[1], xyz[2], -xyz[0]}  //-y,z,-x
	allPerms[15] = [3]int{-xyz[1], -xyz[0], -xyz[2]} //-y,-x,-z

	allPerms[16] = [3]int{xyz[2], xyz[0], xyz[1]}   //z,x,y
	allPerms[17] = [3]int{xyz[2], xyz[1], -xyz[0]}  //z,y,-x
	allPerms[18] = [3]int{xyz[2], -xyz[0], -xyz[1]} //z,-x,-y
	allPerms[19] = [3]int{xyz[2], -xyz[1], xyz[0]}  //z,-y,x

	allPerms[20] = [3]int{-xyz[2], -xyz[0], xyz[1]}  //-z,-x,y
	allPerms[21] = [3]int{-xyz[2], xyz[1], xyz[0]}   //-z,y,x
	allPerms[22] = [3]int{-xyz[2], xyz[0], -xyz[1]}  //-z,x,-y
	allPerms[23] = [3]int{-xyz[2], -xyz[1], -xyz[0]} //-z,-y,-x
	return allPerms
}
