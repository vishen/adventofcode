package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	p := NewParser(data)
	p.run()
}

type state int

const (
	state_NAME state = iota
	state_CHILD_NUMBER
	state_CHILD_NAME
	state_BAG
)

type bag struct {
	name   string
	number int

	contains []bag
}

type parser struct {
	data []byte
	i    int
	s    state
}

func NewParser(data []byte) *parser {
	return &parser{
		data: data,
		i:    0,
		s:    state_NAME,
	}
}

func (p *parser) run() {
	bags := map[string]bag{}
	childToParents := map[string]map[string]bool{}

	currentBag := bag{}
	childBag := bag{}

	start := -1
	for {
		if p.i >= len(p.data) {
			break
		}
		c := p.data[p.i]
		switch p.s {
		case state_NAME:
			if c >= 'a' && c <= 'z' && start == -1 {
				start = p.i
			}
			if p.isWord(" bags") {
				currentBag.name = string(p.data[start:p.i])
				start = -1
				p.eatWord(" bags contain")
				childBag = bag{}
				p.s = state_CHILD_NUMBER
			}
		case state_CHILD_NUMBER:
			if c == '.' {
				p.s = state_NAME
				bags[currentBag.name] = currentBag
				currentBag = bag{}
				childBag = bag{}
				start = -1
			}
			if c >= '0' && c <= '9' {
				childBag.number = int(c - '0')
				p.s = state_CHILD_NAME
				start = -1
			}
		case state_CHILD_NAME:
			if c >= 'a' && c <= 'z' && start == -1 {
				start = p.i
			}
			if p.isWord(" bag") {
				childBag.name = string(p.data[start:p.i])
				if _, ok := childToParents[childBag.name]; !ok {
					childToParents[childBag.name] = map[string]bool{}
				}
				childToParents[childBag.name][currentBag.name] = true
				start = -1
				currentBag.contains = append(currentBag.contains, childBag)
				childBag = bag{}
				p.s = state_CHILD_NUMBER
			}
		}
		p.i++
	}

	{

		seen := map[string]bool{}
		parents := childToParents["shiny gold"]
		for {
			if len(parents) == 0 {
				break
			}
			for parent, _ := range parents {
				seen[parent] = true
				for pBag, _ := range childToParents[parent] {
					parents[pBag] = true
				}
				delete(parents, parent)
			}
		}
		fmt.Println(len(seen), len(bags))
	}
}

func (p *parser) isWord(word string) bool {
	return string(p.data[p.i:p.i+len(word)]) == word
}

func (p *parser) eatWord(word string) {
	p.i += len(word)
}
