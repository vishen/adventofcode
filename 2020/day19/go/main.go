package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	rules = map[string]*rule{}
)

type rule struct {
	id        string
	children1 []string
	children2 []string

	repeat bool

	c string

	cached []string
}

func main() {
	data, err := ioutil.ReadFile("./input8.txt")
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
			id := string(lineSplit[0])
			r := &rule{id: id}

			if bytes.HasPrefix(lineSplit[1], []byte{' ', '"'}) {
				r.c = string(lineSplit[1][2])
			} else {
				for i, children := range bytes.Split(lineSplit[1], []byte{'|'}) {
					if len(children) == 0 {
						continue
					}
					for _, c := range bytes.Split(children, []byte{' '}) {
						if len(c) == 0 {
							continue
						}
						if i == 0 {
							r.children1 = append(r.children1, string(c))
						} else {
							r.children2 = append(r.children2, string(c))
						}
					}
				}
			}
			rules[id] = r
		} else {
			toCheck = append(toCheck, string(line))
		}
	}

	for id, rule := range rules {
		fmt.Println(id, rule)
	}

	total := 0
	for _, check := range toCheck {
		matched := match(check)
		if matched {
			total += 1
		}
		fmt.Println(check, matched)
	}
	fmt.Println("Matched:", total)
}

func match(input string) bool {
	i := 0
	r := rules["0"]

	var step func(rid string, children []string) bool
	step = func(rid string, children []string) bool {
		if len(children) == 0 {
			return false
		}
		matched := true
		for ci, c := range children {
			rc := rules[c]
			if rc.c != "" {
				if i >= len(input) {
					break
				}
				if rc.c != string(input[i]) {
					matched = false
					break
				}
				i++
				if i >= len(input) {
					break
				}
			} else if c == rid {
				// handle loop
				count := 0
				restMatch := false
				for {
					pi := i
					if !step(c, rc.children1) {
						i = pi
						break
					}

					{
						// Check to see if the last repeat match is the same as
						// the next in children, if we have children remaining
						if len(children) > ci {
							ppi := i
							i = pi
							restMatch = step(rid, children[ci+1:])
							i = ppi
							fmt.Println("TRYING REPEAT", restMatch, pi)
						}
					}

					count += 1
					if i >= len(input) {
						break
					}
				}
				if count == 0 {
					matched = false
					break
				}
				if restMatch {
					fmt.Println("AAAAAAAA")
					break
				}

			} else {
				pi := i
				matched1 := step(c, rc.children1)
				pi1 := i
				i = pi
				matched2 := step(c, rc.children2)
				if !matched2 && matched1 {
					i = pi1
				}
				matched = matched1 || matched2
				if !matched {
					break
				}
			}
		}
		return matched
	}

	return step(r.id, r.children1) && i == len(input)
}

/*

nstate := []string{"0"}

	for _, char := range []byte(message) {
		var next []string

		var step func(string)
		step = func(state string) {
			if state == "" {
				return
			}

			head, rest := state, ""
			if i := strings.Index(head, " "); i >= 0 {
				head, rest = head[:i], head[i:]
			}

			if head[0] == '"' {
				if head[1] == char {
					next = append(next, strings.TrimPrefix(rest, " "))
				}
			} else if rule, ok := rules[head]; ok {
				for _, alt := range strings.Split(rule, " | ") {
					step(alt + rest)
				}
			} else {
				panic(head)
			}
		}

		for _, state := range nstate {
			step(state)
		}
		nstate = next
	}

	for _, state := range nstate {
		if len(state) == 0 {
			return true
		}
	}
*/

/*

type engine struct {
	input    string
	i        int
	inRepeat bool
}

func (e *engine) ended() bool {
	fmt.Println("ended", e.i, len(e.input))
	return e.i == len(e.input)
}

func (e *engine) checkRules(r *rule, input string) bool {
	if r.c != "" {
		return string(input[0]) == r.c
	}

	found := true
	for _, c := range r.children1 {
		rc := rules[c]
		if !e.checkRules(rc) {
			found = false
			break
		}
		e.i++
	}
	if found {
		return true
	}
	return true
}

*/

/*
func (e *engine) checkRules(r *rule) bool {
	fmt.Printf("d1: rid=%s, i=%d\n", r.id, e.i)
	if r.c != "" {
		if e.i >= len(e.input) {
			return false
			// panic("shouldn't get here")
		}
		fmt.Printf("d2: rid=%s, equal=%t\n", r.id, r.c == string(e.input[e.i]))
		equal := r.c == string(e.input[e.i])
		e.i++
		return equal
	}

	pi := e.i
	found := false
	if !e.inRepeat && len(r.children2) > 0 {
		found = true
		for _, c := range r.children2 {
			fmt.Printf("d4: rid=%s c2=%s\n", r.id, c)
			if r.id == c {
				fmt.Println("repeating")
				e.inRepeat = true
				count := 0
				for {
					if !e.checkRules(r) {
						break
					}
					count += 1
				}
				e.inRepeat = false
				if count == 0 {
					found = false
					break
				}
			} else {
				rc := rules[c]
				if !e.checkRules(rc) {
					found = false
					break
				}
				fmt.Printf("d6: rid=%s c2=%s e.i=%d\n", r.id, c, e.i)
				if e.i >= len(e.input) {
					break
				}
			}
		}
	}
	if found {
		return true
	}

	e.i = pi

	found = true
	for _, c := range r.children1 {
		fmt.Printf("d3: rid=%s c1=%s\n", r.id, c)
		rc := rules[c]
		if !e.checkRules(rc) {
			found = false
			break
		}
		fmt.Printf("d5: rid=%s c1=%s e.i=%d\n", r.id, c, e.i)
		if e.i >= len(e.input) {
			break
		}
	}

	if found {
		return true
	}

	e.i = pi
	return false
}
*/

func buildCache(r *rule) []string {
	if r.cached != nil {
		return r.cached
	}

	if r.c != "" {
		return []string{r.c}
	}

	children1 := []string{}
	for _, c := range r.children1 {
		if c == r.id {
			r.repeat = true
			continue
		}
		values := buildCache(rules[c])
		if len(children1) == 0 {
			children1 = append(children1, values...)
		} else {
			children1 = join(children1, values...)
		}
	}
	results := make([]string, len(children1))
	copy(results, children1)

	children2 := []string{}
	for _, c := range r.children2 {
		var values []string
		if c == r.id {
			r.repeat = true
			continue
			// values = join([]string{"R"}, children1...)
		} else {
			values = buildCache(rules[c])
		}
		if len(children2) == 0 {
			children2 = append(children2, values...)
		} else {
			children2 = join(children2, values...)
		}
	}

	results = append(results, children2...)
	r.cached = results
	return results
}

func join(l1 []string, l2 ...string) []string {
	l3 := []string{}
	for _, l1 := range l1 {
		for _, l2 := range l2 {
			l3 = append(l3, l1+l2)
		}
	}
	return l3
}

/*
func buildCache(r *rule) []string {
	if r.cached != nil {
		return r.cached
	}

	if r.c != "" {
		return []string{r.c}
	}

	children1 := []string{}
	for _, c := range r.children1 {
		if c == r.id {
			r.repeat = true
			continue
		}
		values := buildCache(rules[c])
		if len(children1) == 0 {
			children1 = append(children1, values...)
		} else {
			children1 = join(children1, values...)
		}
	}
	results := make([]string, len(children1))
	copy(results, children1)

	children2 := []string{}
	for _, c := range r.children2 {
		var values []string
		if c == r.id {
			r.repeat = true
			continue
			// values = join([]string{"R"}, children1...)
		} else {
			values = buildCache(rules[c])
		}
		if len(children2) == 0 {
			children2 = append(children2, values...)
		} else {
			children2 = join(children2, values...)
		}
	}

	results = append(results, children2...)
	r.cached = results
	return results
}

	fmt.Println("Generating matches...")
	m := matches("0")
	fmt.Println("looking for matches")
	return
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
*/
