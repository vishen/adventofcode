package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed sample
	sample string
	//go:embed input1
	input1 string
)

func main() {
	part1(sample)
	part1(input1)
	part2(sample)
}

func pprint(diskmap []int) {
	for _, d := range diskmap {
		if d < 0 {
			fmt.Print(".")
		} else {
			fmt.Print(d)
		}
	}
	fmt.Println()
}

func part1(data string) {
	var diskmap []int

	blockFile := true
	curID := 0
	frees := 0
	for _, ch := range data {
		if ch == '\n' {
			continue
		}
		ich := int(ch - '0')
		val := -1
		if blockFile {
			val = curID
			curID++
		} else {
			frees += ich
		}
		for i := 0; i < ich; i++ {
			diskmap = append(diskmap, val)
		}
		blockFile = !blockFile
	}

	moved := 0
	for d1 := len(diskmap) - 1; d1 >= 0; d1-- {
		done := false
		for d2 := 0; d2 < len(diskmap); d2++ {
			if diskmap[d2] == -1 {
				diskmap[d2], diskmap[d1] = diskmap[d1], diskmap[d2]
				moved++
				if moved == frees {
					done = true
				}
				break
			}
		}
		if done {
			break
		}
	}
	checksum := 0
	for di, d := range diskmap {
		if d == -1 {
		} else {
			checksum += di * d
		}
	}
	fmt.Println(checksum)
}

type block struct {
	total int
	used  []int
}

func (b block) free() int {
	return b.total - len(b.used)
}

func part2(data string) {

	curID := 0
	blockFile := true
	blocks := []block{}
	for _, ch := range data {
		if ch == '\n' {
			continue
		}
		ich := int(ch - '0')
		if blockFile {
			var used []int
			for i := 0; i < ich; i++ {
				used = append(used, curID)
			}
			blocks = append(blocks, block{used: used, total: len(used)})
			curID++
		} else {
			blocks[curID-1].total += ich
		}
		blockFile = !blockFile
	}

	for _, b := range blocks {
		for _, u := range b.used {
			fmt.Print(u)
		}
		for f := 0; f < b.free(); f++ {
			fmt.Print(".")
		}
	}
	fmt.Println()

	/*
		for be := len(blocks) - 1; be >= 0; be-- {
			fmt.Println("BE", be, blocks[be], len(blocks[be].used))
			for bs := 0; bs <= be; bs++ {
				fmt.Println("  BS", bs, blocks[bs], blocks[bs].free())

				blocks := []

				if len(blocks[be].used) <= blocks[bs].free() {
					id := blocks[be].used[0]
					for bi := 0; bi < len(blocks[be].used); bi++ {
						blocks[bs].used = append(blocks[bs].used, id)
					}
					blocks[be].used = []int{}
					break
				}
			}
		}
	*/

	for _, b := range blocks {
		for _, u := range b.used {
			fmt.Print(u)
		}
		for f := 0; f < b.free(); f++ {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

/*
type block struct {
	id    int
	total int
	used  []int
}

func (b block) free() int {
	return b.total - len(b.used)
}
func part2(data string) {

	curID := 0
	blockFile := true
	blocks := []block{}
	for _, ch := range data {
		if ch == '\n' {
			continue
		}
		ich := int(ch - '0')
		if blockFile {
			var used []int
			for i := 0; i < ich; i++ {
				used = append(used, curID)
			}
			blocks = append(blocks, block{id: curID, used: used, total: len(used)})
			curID++
		} else {
			blocks[curID-1].total += ich
		}
		blockFile = !blockFile
	}

	for _, b := range blocks {
		for _, u := range b.used {
			fmt.Print(u)
		}
		for f := 0; f < b.free(); f++ {
			fmt.Print(".")
		}
	}
	fmt.Println()

	for be := len(blocks) - 1; be >= 0; be-- {
		fmt.Println("BE", be, blocks[be], len(blocks[be].used))
		for bs := 0; bs <= be; bs++ {
			fmt.Println("  BS", bs, blocks[bs], blocks[bs].free())
			if len(blocks[be].used) <= blocks[bs].free() {
				for bi := 0; bi < len(blocks[be].used); bi++ {
					blocks[bs].used = append(blocks[bs].used, blocks[be].id)
				}
				blocks[be].used = []int{}
				break
			}
		}
	}

	for _, b := range blocks {
		for _, u := range b.used {
			fmt.Print(u)
		}
		for f := 0; f < b.free(); f++ {
			fmt.Print(".")
		}
	}
	fmt.Println()

		for d1 := len(diskmap) - 1; d1 >= 0; {
			pprint(diskmap)
			required := 1
			for d11 := d1 - 1; d11 >= 0; d11-- {
				if diskmap[d11] == diskmap[d1] {
					required++
				} else {
					break
				}

			}
			done := false
			for d2 := 0; d2 < len(diskmap); d2++ {
				if diskmap[d2] == -1 {
					found := 1
					for d22 := d2 + 1; d22 < len(diskmap); d22++ {
						if diskmap[d22] == -1 {
							found++
						} else {
							break
						}
					}
					if required <= found {
						for d22 := d2; d22 < d2+required; d22++ {
							diskmap[d22] = diskmap[d1]
						}
						for d11 := d1; d11 >= d1-required; d11-- {
							diskmap[d11] = -2
						}
					}
					break
				}
			}
			d1 -= required
			if done {
				break
			}
		}
		pprint(diskmap)
		checksum := 0
		for di, d := range diskmap {
			if d == -1 {
			} else {
				checksum += di * d
			}
		}
		fmt.Println(checksum)
}
*/
