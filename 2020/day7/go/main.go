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
	name string

	parents  []*bag
	children []*bag
	numbers  []int
}

func newBag(name string) *bag {
	return &bag{
		name: name,
	}
}

type bags map[string]*bag

func (b bags) Add(bagName string) {
	if _, ok := b[bagName]; !ok {
		b[bagName] = newBag(bagName)
	}
}

func (b bags) AddChild(bagName string, childBagName string, childBagNumber int) {
	parent := b[bagName]

	child, ok := b[childBagName]
	if !ok {
		child = newBag(childBagName)
		b[child.name] = child
	}
	child.parents = append(child.parents, parent)

	parent.children = append(parent.children, child)
	parent.numbers = append(parent.numbers, childBagNumber)
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
	bags := bags{}

	var currentBagName, childBagName string
	var childBagNumber int

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
				currentBagName = string(p.data[start:p.i])
				bags.Add(currentBagName)
				start = -1
				p.eatWord(" bags contain")
				p.s = state_CHILD_NUMBER
			}
		case state_CHILD_NUMBER:
			if c == '.' {
				p.s = state_NAME
				start = -1
			}
			if c >= '0' && c <= '9' {
				childBagNumber = int(c - '0')
				p.s = state_CHILD_NAME
				start = -1
			}
		case state_CHILD_NAME:
			if c >= 'a' && c <= 'z' && start == -1 {
				start = p.i
			}
			if p.isWord(" bag") {
				childBagName = string(p.data[start:p.i])
				start = -1
				bags.AddChild(currentBagName, childBagName, childBagNumber)
				p.s = state_CHILD_NUMBER
			}
		}
		p.i++
	}

	// Find all parents of "shiny gold"
	fmt.Println(numberOfParents(bags["shiny gold"]))

	// Find number of child bags of "shiny gold"
	fmt.Println(totalForBag(bags["shiny gold"]))
}

func (p *parser) isWord(word string) bool {
	return string(p.data[p.i:p.i+len(word)]) == word
}

func (p *parser) eatWord(word string) {
	p.i += len(word)
}

func numberOfParents(b *bag) int {
	seen := map[string]bool{}
	parents := b.parents
	for {
		if len(parents) == 0 {
			break
		}
		nextParents := []*bag{}
		for _, p := range parents {
			seen[p.name] = true
			nextParents = append(nextParents, p.parents...)
		}
		parents = nextParents
	}
	return len(seen)
}

func totalForBag(b *bag) int {
	if len(b.children) == 0 {
		return 0
	}
	total := 0
	for i := range b.children {
		num := b.numbers[i]
		t := totalForBag(b.children[i]) * num
		total += t + num
	}
	return total
}
