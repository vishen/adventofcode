package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	rules = map[string][]byte{}
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(data []byte) {

	toCheck := []string{}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		lineSplit := bytes.Split(line, []byte{':'})
		if len(lineSplit) > 1 {
			rules[string(lineSplit[0])] = append(rules[string(lineSplit[0])], lineSplit[1]...)
		} else {
			toCheck = append(toCheck, string(line))
		}
	}

	m := matches("0")
	matched := 0
	for i, c := range toCheck {
		for _, m1 := range m {
			if c == m1 {
				matched += 1
				fmt.Printf("MATCHED: %d, %q\n", i+1, c)
				break
			}
		}
	}
	fmt.Println("matched", matched)
}

func matches(key string) []string {
	values := rules[key]
	var toReturn []string

	found := []string{}
	for _, c := range bytes.Split(values, []byte{' '}) {
		if len(c) == 0 {
			continue
		}
		if c[0] == '|' {
			toReturn = append(toReturn, found...)
			found = []string{}
		} else if c[0] >= '0' && c[0] <= '9' {
			m := matches(string(c))
			if len(found) == 0 {
				found = m
			} else {
				found = join(found, m)
			}
		} else if c[0] == '"' {
			// This only works since the letters are single characters
			return []string{string(c[1])}
		}
	}

	if len(toReturn) == 0 {
		return found
	} else {
		return append(toReturn, found...)
	}
}

func join(l1, l2 []string) []string {
	l3 := []string{}
	for _, l1 := range l1 {
		for _, l2 := range l2 {
			l3 = append(l3, l1+l2)
		}
	}
	return l3
}
