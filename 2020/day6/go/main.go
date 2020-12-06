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

	run(data)
}

type state int

const (
	state_LETTER state = iota
	state_SEPERATOR
)

func run(data []byte) {
	total := 0

	form := make(map[byte]int, 26)
	people := 0
	s := state_LETTER

	for _, c := range data {
		switch s {
		case state_LETTER:
			switch {
			case c >= 'a' && c <= 'z':
				form[c] += 1
			case c == '\n':
				s = state_SEPERATOR
				people += 1
			}
		case state_SEPERATOR:
			switch {
			case c >= 'a' && c <= 'z':
				form[c] += 1
				s = state_LETTER
			case c == '\n':
				total += allForForm(people, form)
				form = map[byte]int{}
				people = 0
				s = state_LETTER
			}
		}
	}

	if len(form) > 0 {
		total += allForForm(people, form)
	}

	fmt.Println(total)
}

func allForForm(num int, form map[byte]int) int {
	total := 0
	for _, v := range form {
		if v == num {
			total++
		}
	}
	return total
}
