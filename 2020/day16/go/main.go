package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input2.txt")
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
	nearbyLen := 0

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
			if len(line) == 0 {
				continue
			}
			tickets := bytes.Split(line, []byte{','})
			if nearbyLen == 0 {
				nearbyLen = len(tickets)
			}
			for _, ticket := range tickets {
				nearby = append(nearby, convertToInt(ticket))
			}
		}
	}

	validNearby := []int{}
	for _, n := range nearby {
		valid := false
		for _, fs := range fields {
			for _, f := range fs {
				if n >= f.from && n <= f.to {
					valid = true
				}
			}
		}
		if valid {
			validNearby = append(validNearby, n)
		}
	}

	possibleField := map[string][]int{}
	for name := range fields {
		for i := 0; i < len(fields); i++ {
			possibleField[name] = append(possibleField[name], 1)
		}
	}
	fmt.Println(possibleField)

	// Add our tickets
	validNearby = append(validNearby, tickets...)

	likelyFields := map[string]int{}
	run := true
	for run {
		run = false
		for i, n := range validNearby {
			index := i % nearbyLen
			for name, fs := range fields {
				v := false
				for _, f := range fs {
					if n >= f.from && n <= f.to {
						v = true
					}
				}
				if !v {
					if _, ok := possibleField[name]; ok {
						possibleField[name][index] = 0
					}
				}
			}
		}

		fmt.Println("--------------------")
		fmt.Println(possibleField)
		for i := 0; i < len(fields); i++ {
			c := 0
			name := ""
			for n, fs := range possibleField {
				if fs[i] == 1 {
					name = n
					c += 1
				}
			}
			fmt.Println(name, c)

			if c == 1 {
				likelyFields[name] = i
				delete(possibleField, name)
				fmt.Println("FOUND")
				run = true
			}
		}

	}

	val := 1
	for name, field := range likelyFields {
		fmt.Println(name, field)
		if strings.HasPrefix("departure ", name) {
			val *= tickets[field]
		}
	}
	fmt.Println(val)
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
