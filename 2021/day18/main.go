package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
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
	part1(input1)
	part2(input1)
}

func part1(data []byte) {
	var cur []token
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		if cur == nil {
			cur = tokens(line)
		} else {
			toks := tokens(line)
			cur = add(cur, toks)
		}
		cur = reduce(cur)
	}

	root := toNum(cur)
	m := magnitude(root)
	fmt.Printf("Part 1: magnitude %d\n", m)
}

func part2(data []byte) {
	numbers := [][]token{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		numbers = append(numbers, tokens(line))
	}

	mag := 0
	for i, toks := range numbers {
		for _, otoks := range numbers[i:] {
			{
				ncur := reduce(add(otoks, toks))
				m := magnitude(toNum(ncur))
				if m > mag {
					mag = m
				}
			}
			{
				ncur := reduce(add(toks, otoks))
				m := magnitude(toNum(ncur))
				if m > mag {
					mag = m
				}
			}
		}
	}

	fmt.Printf("Part 2: largest magnitude %d\n", mag)
}

func add(t1, t2 []token) []token {
	// [1,2] + [[3,4],5] becomes [[1,2],[[3,4],5]]
	n := []token{{typ: "["}}
	n = append(n, t1...)
	n = append(n, token{typ: ","})
	n = append(n, t2...)
	return append(n, token{typ: "]"})
}

func magnitude(root *num) int {
	// [9,1] is 3*9 + 2*1 = 29
	// [1,9] is 3*1 + 2*9 = 21
	// [[9,1],[1,9]] is 3*29 + 2*21 = 129
	// [[4, [9,1]],[1,9]]
	total := 0
	mul := 3
	for _, p := range root.pairs {
		val := p.val
		if p.child != nil {
			val = magnitude(p.child)
		}
		if root.parent == nil {
			return val
		}
		total += val * mul
		mul = 2
	}
	return total
}

type pair struct {
	val   int
	child *num
}

type num struct {
	parent *num
	pairs  []pair
	depth  int
}

func (n *num) addValue(val int) {
	n.pairs = append(n.pairs, pair{val: val})
}

func (n *num) addChild(c *num) {
	n.pairs = append(n.pairs, pair{child: c})
}

func toNum(toks []token) *num {
	root := &num{}
	cur := root

	for _, tok := range toks {
		switch tok.typ {
		case "[":
			n := &num{}
			n.parent = cur
			cur.addChild(n)
			cur = n
		case "]":
			cur = cur.parent
		case "LIT":
			cur.addValue(tok.val)
		}
	}
	return root
}

func printNum(n *num, depth int) []byte {
	for d := 0; d < depth; d++ {
		fmt.Print("*")
	}
	fmt.Printf(" %p) depth=%d parent=%p, pairs=%v\n", n, depth, n.parent, n.pairs)
	line := []byte{'['}
	for i, p := range n.pairs {
		if p.child != nil {
			line = append(line, printNum(p.child, depth+1)...)
		} else {
			line = append(line, []byte(strconv.Itoa(p.val))...)
		}
		if i == 0 {
			line = append(line, ',')
		}
	}
	line = append(line, ']')
	if depth == 0 {
		fmt.Printf("LINE: %s\n", line)
	}
	return line
}

func ptokens(toks []token) {
	for _, tok := range toks {
		switch t := tok.typ; t {
		case "LIT":
			fmt.Print(tok.val)
		default:
			fmt.Print(t)
		}
	}
	fmt.Println()
}

func reduce(toks []token) []token {
	for r := 0; ; r++ {
		p := len(toks)
		toks = reduce_(toks)
		if p == len(toks) {
			break
		}
	}
	return toks
}

func reduce_(toks []token) []token {
	depth := 0
	found10 := 0
	for i, tok := range toks {
		switch tok.typ {
		case "[":
			depth++
			continue
		case "]":
			depth--
			continue
		case "LIT":
			if tok.val >= 10 && found10 == 0 {
				found10 = i
			}
		}
		if depth == 5 {
			// explode
			// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4]
			// [7,[6,[5,[4,[3,2]]]]] becomes [7,[6,[5,[7,0]]]]
			// [[6,[5,[4,[3,2]]]],1] becomes [[6,[5,[7,0]]],3]
			ntoks := make([]token, len(toks[:i-1]))
			copy(ntoks, toks[:i-1])

			ntoks = append(ntoks, token{typ: "LIT", val: 0})
			v := len(ntoks)
			ntoks = append(ntoks, toks[i+4:]...)

			for j := i - 2; j >= 0; j-- {
				if ntoks[j].typ == "LIT" {
					ntoks[j].val += tok.val
					break
				}
			}

			assert(toks[i+1].typ == ",")
			assert(toks[i+2].typ == "LIT")

			for j := v; j < len(ntoks); j++ {
				if ntoks[j].typ == "LIT" {
					ntoks[j].val += toks[i+2].val
					break
				}
			}
			return ntoks
		}
	}

	if found10 > 0 {
		assert(toks[found10].typ == "LIT")
		v := toks[found10].val
		assert(v >= 10)

		v1 := v / 2
		v2 := v1
		if v%2 == 1 {
			v2 += 1
		}

		ctoks := make([]token, len(toks))
		copy(ctoks, toks)
		ntoks := append(
			ctoks[:found10],
			[]token{
				token{typ: "["},
				token{typ: "LIT", val: v1},
				token{typ: ","},
				token{typ: "LIT", val: v2},
				token{typ: "]"},
			}...,
		)

		toks = append(ntoks, toks[found10+1:]...)
	}

	return toks
}

type token struct {
	typ string
	val int
}

func tokens(line []byte) []token {
	var toks []token
	i := 0
	for {
		if i >= len(line) {
			break
		}
		c := line[i]
		if c >= '0' && c <= '9' {
			val, ni := readNum(line[i:])
			toks = append(toks, token{typ: "LIT", val: val})
			i += ni
			continue
		}

		switch c {
		case '[', ']', ',':
			toks = append(toks, token{typ: string(c)})
		}
		i++
	}
	return toks
}

func readNum(data []byte) (int, int) {
	val := 0
	end := 0
	for i, c := range data {
		if c >= '0' && c <= '9' {
			val *= 10
			val += int(c - '0')
		} else {
			end = i
			break
		}
	}
	return val, end
}

func assert(t bool) { //, msg string, args ...interface{}) {
	if !t {
		// panic(fmt.Sprintf(msg, args...))
		panic("ASSERTION FAILED")
	}
}
