package main

import (
	_ "embed"
	"fmt"
	"math/bits"
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

/*
NOTE: This doesn't solve it automatically, but only gives the place the error occurs.

The main thing to look for is that we're expecting an:

	<z_output> = <input> XOR <input>
	  <output> = <x_input_same_number_as_z_output> XOR <y_input_same_number_as_z_output>
	  <output> = <input> OR <input>
	    <output> = <x_input_for_previous_z> AND <y_input_for_previous_z>

eg:

	  z29 = qdw XOR mhh
		qdw = jbf OR bmh
		  jbf = jsd AND kbc
		  bmh = y28 AND x28
		mhh = x29 XOR y29
*/
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
	expected := x + y

	var finals []string
	var steps []step
	deps := map[string]step{}

	swaps := map[string]string{
		"z05": "hdt",
		"z09": "gbf",
		"jgt": "mht",
		"z30": "nbf",
	}

	for _, line := range strings.Split(parts[1], "\n") {
		inputsOutput := strings.Split(line, " -> ")
		inputs := strings.Split(inputsOutput[0], " ")
		output := inputsOutput[1]

		for s1, s2 := range swaps {
			switch output {
			case s1:
				output = s2
			case s2:
				output = s1
			}
		}

		s := step{inputs, output, false}
		steps = append(steps, s)
		deps[output] = s
		if output[0] == 'z' {
			finals = append(finals, output)
		}
	}

	seen := map[string]bool{}
	var eval func(string, int)
	results := run(steps, values, finals)
	z := 0
	for i, r := range results {
		if r {
			z |= 1 << i
		}
	}

	if expected == z {
		var codes []string
		for s1, s2 := range swaps {
			codes = append(codes, s1, s2)
		}
		sort.Strings(codes)
		fmt.Println(strings.Join(codes, ","))
		return
	}

	diff := z ^ expected
	fmt.Println("EXPECTED", strconv.FormatInt(int64(expected), 2))
	fmt.Println("ACTUAL  ", strconv.FormatInt(int64(z), 2))
	fmt.Println("DIFF    ", strconv.FormatInt(int64(diff), 2))

	eval = func(cur string, depth int) {
		s := deps[cur]
		ns := s
		if s.output == "" {
			return
		}
		line := fmt.Sprintf("%s = %s", s.output, strings.Join(s.inputs, " "))
		if _, ok := seen[line]; !ok {
			for d := 0; d < depth; d++ {
				fmt.Print(" ")
			}
			fmt.Println(line)
		}
		// sort.Strings(ns.inputs)
		seen[line] = true
		for _, i := range ns.inputs {
			eval(i, depth+1)
		}
	}

	pos := bits.TrailingZeros(uint(diff))
	for i := pos - 3; i < pos+3; i++ {
		zs := "z"
		if i < 10 {
			zs += "0" + strconv.Itoa(i)
		} else {
			zs += strconv.Itoa(i)
		}
		if i == pos {
			fmt.Println()
			fmt.Println("####-------WRONG-------------####")
			fmt.Println()
		}
		eval(zs, 0)
		if i == pos {
			fmt.Println()
			fmt.Println("################################")
			fmt.Println()
		} else {
			fmt.Println("----------------")
		}
	}

}
