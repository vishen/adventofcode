package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed sample
	sample string
	//go:embed sample2
	sample2 string
	//go:embed sample3
	sample3 string

	//go:embed input
	input string
)

func main() {
	part1(sample)
	part1(sample2)
	part1(input)
	part2(input)
}

func part1(data string) {
	parts := strings.Split(strings.TrimSpace(data), "\n\n")

	values := map[string]bool{}
	for _, line := range strings.Split(parts[0], "\n") {
		vs := strings.Split(line, ": ")
		v := false
		if vs[1] == "1" {
			v = true
		}
		values[vs[0]] = v
	}

	var steps []step
	var finals []string
	for _, line := range strings.Split(parts[1], "\n") {
		inputsOutput := strings.Split(line, " -> ")
		inputs := strings.Split(inputsOutput[0], " ")
		output := inputsOutput[1]
		steps = append(steps, step{inputs, output, false})
		if output[0] == 'z' {
			finals = append(finals, output)
		}
	}

	results := run(steps, values, finals)

	final := 0
	for i, r := range results {
		if r {
			final |= 1 << i
		}
	}
	fmt.Println(final)
}

type step struct {
	inputs    []string
	output    string
	completed bool
}

func run(steps []step, values map[string]bool, finals []string) []bool {
	for {
		completed := 0
		for _, s := range steps {
			if s.completed {
				completed += 1
				continue
			}

			in1, ok := values[s.inputs[0]]
			if !ok {
				continue
			}
			in2, ok := values[s.inputs[2]]
			if !ok {
				continue
			}

			v := false
			switch s.inputs[1] {
			case "OR": // TODO: Doesn't need to be dependant
				v = in1 == true || in2 == true
			case "AND":
				v = in1 == true && in2 == true
			case "XOR":
				v = in1 != in2
			}
			values[s.output] = v
			s.completed = true
			completed += 1
		}
		if completed == len(steps) {
			break
		}
	}
	sort.Slice(finals, func(i, j int) bool {
		return finals[i] < finals[j]
	})
	var result []bool
	for _, f := range finals {
		result = append(result, values[f])
	}
	return result
}

func part2(data string) {
	parts := strings.Split(strings.TrimSpace(data), "\n\n")

	values := map[string]bool{}
	var xs []string
	for _, line := range strings.Split(parts[0], "\n") {
		vs := strings.Split(line, ": ")
		v := false
		if vs[1] == "1" {
			v = true
		}
		values[vs[0]] = v
		if vs[0][0] == 'x' {
			xs = append(xs, vs[0])
		}
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	x := 0
	y := 0
	for i, xx := range xs {
		if values[xx] {
			x |= 1 << i
		}
		if values[fmt.Sprintf("y%s", string(xx[1:]))] {
			y |= 1 << i
		}
	}
	fmt.Println(strconv.FormatInt(int64(x+y), 2))

	deps := map[string]step{}
	var finals []string
	var steps []step

	for _, line := range strings.Split(parts[1], "\n") {
		inputsOutput := strings.Split(line, " -> ")
		inputs := strings.Split(inputsOutput[0], " ")
		output := inputsOutput[1]
		s := step{inputs, output, false}
		steps = append(steps, s)
		deps[output] = s
		if output[0] == 'z' {
			finals = append(finals, output)
		}
	}

	/*
		results := run(steps, values, finals)
		z := 0
		for i, r := range results {
			if r {
				z |= 1 << i
			}
		}
		fmt.Println(strconv.FormatInt(int64(z), 2))
	*/

	var run func(string) []step
	run = func(cur string) []step {
		s := deps[cur]

		if len(s.inputs) > 0 {
			switch s.inputs[0][0] {
			case 'x', 'y':
				return []step{s}
			}
		}
		var found []step
		for _, i := range s.inputs {
			found = append(found, run(i)...)
		}
		return found
	}

	sort.Slice(finals, func(i, j int) bool {
		return finals[i] < finals[j]
	})

	seen := map[string]bool{}
	_ = seen
	for _, f := range finals[:7] {
		inputs := run(f)
		// fmt.Println(f, x, y)
		sort.Slice(inputs, func(i, j int) bool {
			v1, _ := strconv.Atoi(string(inputs[i].inputs[0][1:]))
			v2, _ := strconv.Atoi(string(inputs[j].inputs[0][1:]))
			if v1 == v2 {
				return inputs[i].inputs[1] < inputs[j].inputs[1]
			}
			return v1 > v2
		})
		// fmt.Println(f, inputs)
		fmt.Println("---------------------")
		fmt.Println(f)
		for _, i := range inputs {
			if seen[i.output] {
				continue
			}
			fmt.Printf("%v || ", i)
			seen[i.output] = true
		}
		fmt.Println()
		fmt.Println(len(inputs), inputs)
		/*
			for ii, i := range inputs {
				if ii > 1 {
					break
				}
				fmt.Println("> ", i)
			}
		*/
	}
}
