package main

import (
	"bytes"
	"fmt"
	"time"

	_ "embed"
)

var (
	//go:embed input1.txt
	input1 []byte

	ins []inst
)

func main() {
	part1(input1)
	part2(input1)
}

type inst struct {
	op string
	// reg1 byte
	reg1 int

	//reg2 byte
	reg2 int
	val2 int
}

/*
	Largely copied from the amazing Jocelyn: https://www.youtube.com/watch?v=KEUTNCRvXN4
	- https://github.com/emilyskidsister/aoc/blob/main/p2021_24/src/lib.rs
*/

func part1(data []byte) {
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		var i inst
		parts := bytes.Split(line, []byte{' '})
		switch cmd := string(parts[0]); cmd {
		case "inp":
			i.op = cmd
			i.reg1 = regIndex(parts[1][0])
		case "add", "mod", "div", "mul", "eql":
			i.op = cmd
			i.reg1 = regIndex(parts[1][0])
			switch reg := parts[2][0]; reg {
			case 'x', 'y', 'w', 'z':
				i.reg2 = regIndex(reg)
			default:
				i.reg2 = -1
				i.val2 = convertToInt(parts[2])
			}
		}
		ins = append(ins, i)
	}
	t := time.Now()
	ans := run(0, [4]int{}, 1)
	fmt.Printf("Part 1: %d %v\n", ans, time.Since(t))
}

func part2(data []byte) {
	if len(ins) == 0 {
		for _, line := range bytes.Split(data, []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}

			var i inst
			parts := bytes.Split(line, []byte{' '})
			switch cmd := string(parts[0]); cmd {
			case "inp":
				i.op = cmd
				i.reg1 = regIndex(parts[1][0])
			case "add", "mod", "div", "mul", "eql":
				i.op = cmd
				i.reg1 = regIndex(parts[1][0])
				switch reg := parts[2][0]; reg {
				case 'x', 'y', 'w', 'z':
					i.reg2 = regIndex(reg)
				default:
					i.reg2 = -1
					i.val2 = convertToInt(parts[2])
				}
			}
			ins = append(ins, i)
		}
	}
	t := time.Now()
	ans := runSmallest(0, [4]int{}, 1)
	fmt.Printf("Part 2: %d %v\n", ans, time.Since(t))
}

func runSmallest(pc int, regs [4]int, depth int) int64 {

	b := block{pc, regs}
	if val, ok := cache[b]; ok {
		return val
	}

	for i := 1; i <= 9; i++ {
		regs := regs
		pc := pc

		// Handle "inp" instruction
		regs[ins[pc].reg1] = i
		pc++

		for {
			done := false
			if pc >= len(ins) {
				if regs[3] == 0 {
					return int64(i)
				} else if regs[3] > 100_000_000 {
					cache[b] = 0
					return 0
				}
				break
			}
			in := ins[pc]
			switch in.op {
			case "inp":
				if val := runSmallest(pc, regs, depth+1); val > 0 {
					v := int64(i)
					for j := depth; j < 14; j++ {
						v *= 10
					}
					return v + val
				}
				done = true
				break
			default:
				regs = execute(in, regs)
			}
			if done {
				break
			}
			pc++
		}
	}
	cache[b] = 0
	return 0
}

func run(pc int, regs [4]int, depth int) int64 {

	b := block{pc, regs}
	if val, ok := cache[b]; ok {
		return val
	}

	for i := 9; i > 0; i-- {
		regs := regs
		pc := pc

		// Handle "inp" instruction
		regs[ins[pc].reg1] = i
		pc++

		for {
			done := false
			if pc >= len(ins) {
				if regs[3] == 0 {
					return int64(i)
				} else if regs[3] > 100_000_000 {
					cache[b] = 0
					return 0
				}
				break
			}
			in := ins[pc]
			switch in.op {
			case "inp":
				if val := run(pc, regs, depth+1); val > 0 {
					v := int64(i)
					for j := depth; j < 14; j++ {
						v *= 10
					}
					return v + val
				}
				done = true
				break
			default:
				regs = execute(in, regs)
			}
			if done {
				break
			}
			pc++
		}
	}
	cache[b] = 0
	return 0
}

func execute(ins inst, regs [4]int) [4]int {
	regOrVal := func(i inst) int {
		if i.reg2 >= 0 {
			return regs[i.reg2]
		} else {
			return i.val2
		}
	}
	switch cmd := ins.op; cmd {
	case "inp":
		panic("aaaaa")
	case "add":
		r1 := ins.reg1
		val := regOrVal(ins)
		regs[r1] += val
	case "mod":
		r1 := ins.reg1
		val := regOrVal(ins)
		regs[r1] %= val
	case "div":
		r1 := ins.reg1
		val := regOrVal(ins)
		if val != 1 {
			regs[r1] /= val
		}
	case "mul":
		r1 := ins.reg1
		val := regOrVal(ins)
		if val == 0 {
			regs[r1] = 0
		} else {
			regs[r1] *= val
		}
	case "eql":
		r1 := ins.reg1
		val := regOrVal(ins)
		if regs[r1] == val {
			regs[r1] = 1
		} else {
			regs[r1] = 0
		}
	default:
		panic("unknown command: " + cmd)
	}
	return regs
}

func regIndex(c byte) int {
	switch c {
	case 'w':
		return 0
	case 'x':
		return 1
	case 'y':
		return 2
	case 'z':
		return 3
	}
	panic("unknown")
}

type block struct {
	i    int
	regs [4]int
}

var (
	cache = map[block]int64{}
)

func convertToInt(val []byte) int {
	v := 0
	neg := false
	for _, v1 := range val {
		if v1 == '-' {
			neg = true
			continue
		}
		v *= 10
		v += int(v1 - '0')
	}
	if neg {
		v *= -1
	}
	return v
}
