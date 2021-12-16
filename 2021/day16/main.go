package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte

	decoding = map[byte]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
)

func main() {

	// part1(test1)
	part1(input1)
	/*
		for _, s := range []string{
			"C200B40A82",
			"04005AC33890",
			"880086C3E88112",
			"CE00C43D881120",
			"D8005AC2A8F0",
			"F600BC2D8F",
			"9C005AC2F8F0",
			"9C0141080250320F1802104A08",
		} {
			part2([]byte(s))
		}*/
	part2(input1)
}

func part1(data []byte) {
	decoded := ""
	for _, c := range data {
		if c == '\n' {
			break
		}
		decoded += decoding[c]
	}

	p := newParser(decoded)
	p.parse()

	fmt.Printf("Part 1: versions total %d\n", p.versions)
}

func part2(data []byte) {
	decoded := ""
	for _, c := range data {
		if c == '\n' {
			break
		}
		decoded += decoding[c]
	}

	p := newParser(decoded)
	total := p.parse()
	fmt.Printf("Part 2: %d\n", total)
}

func newParser(decoded string) *parser {
	return &parser{
		i:       0,
		decoded: decoded,
	}
}

type parser struct {
	i       int
	decoded string

	versions int
}

func (p *parser) eat(l int) string {
	val := p.decoded[p.i : p.i+l]
	p.i += l
	return val
}

func (p *parser) parse() int {
	end := len(p.decoded)
	for {
		if p.i >= len(p.decoded) {
			break
		}
		V := p.eat(3)
		T := p.eat(3)

		p.versions += toInt(V)
		if T == "100" { // 4
			// Literal value packet
			val := ""
			for {
				if p.i >= end {
					break
				}
				if p.eat(1) == "1" {
					val += p.eat(4)
				} else {
					val += p.eat(4)
					break
				}
			}
			return toInt(val)
		}

		vals := []int{}

		// Operator packet
		if I := p.eat(1); I == "0" {
			i := toInt(p.eat(15))
			oi := p.i
			for {
				vals = append(vals, p.parse())
				if p.i >= oi+i {
					break
				}
			}
		} else {
			i := toInt(p.eat(11))
			for x := 0; x < i; x++ {
				vals = append(vals, p.parse())
			}
		}
		switch T {
		case "000":
			t := 0
			for _, v := range vals {
				t += v
			}
			return t
		case "001":
			t := 1
			for _, v := range vals {
				t *= v
			}
			return t
		case "010":
			t := 0
			for i, v := range vals {
				if i == 0 || t > v {
					t = v
				}
			}
			return t
		case "011":
			t := 0
			for i, v := range vals {
				if i == 0 || t < v {
					t = v
				}
			}
			return t
		case "101":
			if vals[0] > vals[1] {
				return 1
			}
			return 0
		case "110":
			if vals[0] < vals[1] {
				return 1
			}
			return 0
		case "111":
			if vals[0] == vals[1] {
				return 1
			}
			return 0
		}
	}
	panic("should not get here")
	return -1
}

func toInt(val string) int {
	iv := 0
	for i := len(val) - 1; i >= 0; i-- {
		if val[i] == '1' {
			iv |= 1 << (len(val) - 1 - i)
		}
	}
	return iv
}
