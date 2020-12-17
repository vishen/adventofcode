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

type state int

const (
	state_fields state = iota
	state_ticket
	state_nearby
	state_nearby_tickets
)

type toFrom struct {
	from, to int
}

func run(data []byte) {
	s := state_fields

	fields := map[string][]toFrom{}

	tickets := []int{}
	nearby := []int{}

	lines := bytes.Split(data, []byte{'\n'})
	for i, line := range lines {
		switch s {
		case state_fields:
			if len(line) == 0 {
				s = state_ticket
				continue
			}

			lineSplit := bytes.SplitN(line, []byte{':'}, 2)
			fmt.Printf("%s\n", lineSplit)

			f1, t1, f2, t2 := 0, 0, 0, 0
			fmt.Sscanf(
				string(lineSplit[1]),
				"%d-%d or %d-%d",
				&f1,
				&t1,
				&f2,
				&t2,
			)
			fields[string(lineSplit[0])] = []toFrom{toFrom{f1, t1}, toFrom{f2, t2}}
		case state_ticket:
			if bytes.HasPrefix(line, []byte("your ticket:")) {
				for _, ticket := range bytes.Split(lines[i+1], []byte{','}) {
					tickets = append(tickets, convertToInt(ticket))
				}
				s = state_nearby
			}
		case state_nearby:
			if bytes.HasPrefix(line, []byte("nearby tickets:")) {
				s = state_nearby_tickets
			}
		case state_nearby_tickets:
			for _, ticket := range bytes.Split(line, []byte{','}) {
				nearby = append(nearby, convertToInt(ticket))
			}
		}
	}

	fmt.Println(fields)
	fmt.Println(tickets)
	// fmt.Println(nearby)

	invalid := 0
	for _, n := range nearby {
		valid := false
		for _, fs := range fields {
			for _, f := range fs {
				if n >= f.from && n <= f.to {
					valid = true
				}
			}
		}
		if !valid {
			// fmt.Println(n)
			invalid += n
		}
	}

	fmt.Println(invalid)
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
