package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(data []byte) {
	sum := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		// Add paratheses around addition to force precendence
		newLine := make([]byte, len(line))
		copy(newLine, line)

		start := 0
		var num byte
		var prev byte
		for i, c := range line {
			if c >= '0' && c <= '9' {

				if prev == '+' {
					newLine = append(newLine[:start], '(', num, prev, c, ')', newLine[i:]...)
				}

				num = c
				start = i
			}
			if c != ' ' {
				prev = c
			}
		}

		total := (&parser{data: newLine}).evaluate()
		fmt.Printf("%s = %d\n", newLine, total)
		sum += total
		return
	}
	fmt.Println(sum)
}

type parser struct {
	data []byte
	i    int
}

func (p *parser) evaluate() int {
	total := 0
	var operator byte = '+'

	for p.i < len(p.data) {
		b := p.data[p.i]
		p.i++

		switch b {
		case '+', '*':
			operator = b
		case '(':
			by := p.evaluate()
			total = value(total, operator, by)
		case ')':
			return total
		case ' ':
			// Ignore space
		default:
			// assume a number
			total = value(total, operator, int(b-'0'))
		}
	}
	return total
}

func value(total int, operator byte, by int) int {
	switch operator {
	case '+':
		return total + by
	case '*':
		return total * by
	}
	return total
}
