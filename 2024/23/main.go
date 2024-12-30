package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

var (
	//go:embed sample
	sample string
	//go:embed input
	input string
)

func main() {
	part1(sample)
	part1(input)
	part2(sample)
	part2(input)
}

func part1(data string) {
	network := map[string]map[string]bool{}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		computers := strings.Split(line, "-")
		if network[computers[0]] == nil {
			network[computers[0]] = map[string]bool{}
		}
		if network[computers[1]] == nil {
			network[computers[1]] = map[string]bool{}
		}
		network[computers[1]][computers[0]] = true
		network[computers[0]][computers[1]] = true
	}

	uniq := map[string]bool{}
	for comp, connected := range network {
		for conn := range connected {
			for conn2 := range network[conn] {
				if _, ok := network[conn2][comp]; ok {
					uniq[key(comp, conn, conn2)] = true
				}
			}
		}
	}
	total := 0
	for u := range uniq {
		for _, comps := range strings.Split(u, ",") {
			if comps[0] == 't' {
				total += 1
				break
			}
		}
	}
	fmt.Println(total)
}

func part2(data string) {
	network := map[string]map[string]bool{}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		computers := strings.Split(line, "-")
		if network[computers[0]] == nil {
			network[computers[0]] = map[string]bool{}
		}
		if network[computers[1]] == nil {
			network[computers[1]] = map[string]bool{}
		}
		network[computers[1]][computers[0]] = true
		network[computers[0]][computers[1]] = true
	}

	// Because all computers are connected to exactly 13 other computers, we can loop through
	// the connected computers and do a count of which of the connected computers
	// are also connected to each other and then we can assume that any cliques
	// that are connected to exactly the same count of computers is a clique.
	highest := 0
	clique := []string{}
	for comp, connected := range network {
		totals := map[string]int{}
		for c := range connected {
			totals[c]++
			for c2 := range network[c] {
				totals[c2]++
			}
		}
		found := map[int][]string{}
		for comp, count := range totals {
			found[count] = append(found[count], comp)
		}
		for count, cliq := range found {
			cliq := cliq
			if count == len(cliq) {
				if count > highest {
					highest = count
					clique = append(cliq, comp)
				}
			}
		}
	}
	sort.Strings(clique)
	fmt.Println(strings.Join(clique, ","))

}

func cpy(m map[string]bool) map[string]bool {
	n := map[string]bool{}
	for k, v := range m {
		n[k] = v
	}
	return n
}

func key(vs ...string) string {
	sort.Slice(vs, func(i, j int) bool {
		return vs[i] < vs[j]
	})
	return strings.Join(vs, ",")
}
