package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

	fmt.Println("Generating matches...")
	m := matches("0")
	fmt.Println("looking for matches")
	matched := 0
	for i, c := range toCheck {
		for _, m1 := range m {
			if match(m1, c) {
				matched += 1
				fmt.Printf("MATCHED: %d, %q -- %q\n", i+1, c, m1)
				break
			}
		}
	}
	fmt.Println("matched", matched)
}

func match(format, value string) bool {
	i := 0
	fi := 0
	//fmt.Println("trying to match:", format, value)
	for {
		//fmt.Println("d8", i, value)
		//fmt.Println("d9", fi, format)
		if len(value) <= i || len(format) <= fi {
			break
		}
		// If we need to loop
		if format[fi] >= '0' && format[fi] <= '9' {
			//fmt.Println("trying to loop")
			l := 0
			j := fi + 1
			// Get number if more than one digit
			for ; ; j++ {
				if j >= len(format) || format[j] < '0' || format[j] > '9' {
					l, _ = strconv.Atoi(format[fi:j])
					break
				}
			}
			repeating := format[j : j+l]
			//fmt.Println("repeating: ", repeating)
			fi = j + l

			// Always has to match at least once
			c := 0
			for {
				if i+len(repeating) > len(value) || value[i:i+len(repeating)] != repeating {
					//		fmt.Println("d4: no repeat found")
					break
				}
				c += 1
				i += len(repeating)
				//fmt.Println("d5: repeat 1 time")
			}
			if c == 0 {
				return false
			}
			continue
		}

		//fmt.Println("d6:", i, value)
		//fmt.Println("d7:", fi, format)

		if format[fi] != value[i] {
			//	fmt.Println("d3: returning")
			return false
		}

		i++
		fi++
	}

	if i < len(value) {
		return false
	}
	// Stupid. This is a hack and returns a false positive in
	// most cases
	// return format[len(format)-1] == value[len(value)-1]
	return true
}

var matchesCache = map[string][]string{}

func matches(key string) []string {
	values := rules[key]
	fmt.Printf("%s=%s\n", key, values)
	var toReturn []string

	found := []string{}
	for _, c := range bytes.Split(values, []byte{' '}) {
		if len(c) == 0 {
			continue
		}
		if c[0] == '|' {
			toReturn = append(toReturn, found...)
			//fmt.Println("toreturn", toReturn)
			found = []string{}
		} else if c[0] >= '0' && c[0] <= '9' {
			// Handle loop
			if string(c) == key {
				l := []string{}
				for _, r := range toReturn {
					l = append(l, fmt.Sprintf("%d%s", len(r), r))
				}
				found = join(found, l...)
				continue
			}
			m, ok := matchesCache[string(c)]
			if !ok {
				m = matches(string(c))
				matchesCache[string(c)] = m
			}
			if len(found) == 0 {
				found = m
			} else {
				//fmt.Println("JOIN", found, m)
				found = join(found, m...)
				//fmt.Println("JOINED", found)
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

func join(l1 []string, l2 ...string) []string {
	fmt.Printf("join: l1=%d l2=%d\n", len(l1), len(l2))
	l3 := []string{}
	for _, l1 := range l1 {
		for _, l2 := range l2 {
			l3 = append(l3, l1+l2)
		}
	}
	fmt.Println("join: done")
	return l3
}
