package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	p := NewParser(data)
	insts := p.parse()
	// emulator(insts)
	bruteForceFix(insts)
}

type state int

const (
	state_OPCODE state = iota
	state_OPCODE_VALUE
)

type parser struct {
	data []byte
	s    state
}

func NewParser(data []byte) *parser {
	return &parser{
		data: data,
		s:    state_OPCODE,
	}
}

func (p *parser) parse() []inst {
	insts := []inst{}
	start := -1
	var opcode string
	for i, c := range p.data {
		switch p.s {
		case state_OPCODE:
			if c < 'a' || c > 'z' {
				opcode = string(p.data[start:i])
				p.s = state_OPCODE_VALUE
				start = -1
			} else {
				if start == -1 {
					start = i
				}
			}
		case state_OPCODE_VALUE:
			if c == '+' || c == '-' || (c >= '0' && c <= '9') {
				if start == -1 {
					start = i
				}
				// happy case
			} else {
				p.s = state_OPCODE
				insts = append(insts, inst{
					opcode: opcode,
					value:  string(p.data[start:i]),
				})
				start = -1
			}
		}
	}
	return insts
}

type inst struct {
	opcode string
	value  string
}

func emulator(insts []inst) bool {
	instRan := make(map[int]bool, len(insts))
	pc := 0
	acc := 0
	for {
		if pc >= len(insts) {
			break
		}
		i := insts[pc]
		if _, ok := instRan[pc]; ok {
			fmt.Println(acc)
			return false
		}
		instRan[pc] = true
		v, _ := strconv.Atoi(i.value[1:])
		switch i.opcode {
		case "acc":
			if i.value[0] == '-' {
				acc -= v
			} else if i.value[0] == '+' {
				acc += v
			}
		case "jmp":
			if i.value[0] == '-' {
				pc -= v
			} else if i.value[0] == '+' {
				pc += v
			}
			continue
		}
		pc += 1
	}
	return true
}

func bruteForceFix(insts []inst) {
	for i := len(insts) - 1; i >= 0; i-- {
		// Change just the last instruction and run again
		// to see if we can fix the instruction that was wrong
		newInsts := insts[:]
		if newInsts[i].opcode == "acc" {
			// Can ignore accumulator opcodes for now
			continue
		}
		newInsts[i] = fixInst(newInsts[i])

		if emulator(newInsts) {
			fmt.Println("found bad instruction:", insts[i])
			return
		}
	}
	fmt.Println("finished without success")
}

func fixInst(i inst) inst {
	switch i.opcode {
	case "jmp":
		i.opcode = "nop"
	case "nop":
		i.opcode = "jmp"
	}
	return i
}
