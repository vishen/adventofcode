package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input1.txt
var input1 string

//go:embed input1.txt
var input2 string

func main() {
	part1()
	part2()
}

func part1() {
	root := parseInput(input1)

	var solve func(*directory) uint64
	solve = func(dir *directory) uint64 {
		total := uint64(0)
		if dir.size <= 100_000 {
			total += dir.size
		}
		for _, d := range dir.dirs {
			total += solve(d)
		}
		return total
	}

	fmt.Println(solve(root))

}

func part2() {
	root := parseInput(input2)

	unused := 70000000 - root.size
	required := 30000000 - unused

	var found []uint64
	var solve func(*directory)
	solve = func(dir *directory) {
		if dir.size >= required {
			found = append(found, dir.size)
		} else {
			return
		}
		for _, d := range dir.dirs {
			solve(d)
		}
	}

	solve(root)
	sort.Slice(found, func(i, j int) bool {
		return found[i] < found[j]
	})
	fmt.Println(found[0])
}

func parseInput(input string) *directory {
	root := &directory{name: "/"}

	cur := root

	addDir := func(name string) *directory {
		dir, ok := cur.dirs[name]
		if !ok {
			dir = &directory{name: name, parent: cur}
			if cur.dirs == nil {
				cur.dirs = make(map[string]*directory)
			}
			cur.dirs[name] = dir
		}
		return dir
	}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		/*
			cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
				cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
				cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
				cd / switches the current directory to the outermost directory, /.

			ls means list. It prints out all of the files and directories immediately contained by the current directory:
				123 abc means that the current directory contains a file named abc with size 123.
				dir xyz means that the current directory contains a directory named xyz.
		*/

		if line[0] == '$' {
			cmds := strings.Split(line[2:], " ")
			switch cmds[0] {
			case "cd":
				switch cmds[1] {
				case "/":
					cur = root
				case "..":
					cur = cur.parent
				default:
					cur = addDir(cmds[1])
				}
			case "ls":
				// Do nothing?
			default:
				log.Fatalf("unexpected command: %v", line)
			}
			continue
		}

		// Not in a $ command
		cmds := strings.Split(line, " ")
		if cmds[0] == "dir" {
			addDir(cmds[1])
			continue
		}
		size, _ := strconv.ParseUint(cmds[0], 10, 64)
		cur.files = append(cur.files, file{cmds[1], size})

		// Add the size to all the parents...
		cur.size += size
		parent := cur.parent
		for {
			if parent == nil {
				break
			}
			parent.size += size
			parent = parent.parent
		}
	}
	return root
}

func printDir(dir *directory, padding int) {
	fmt.Print(">")
	for i := 0; i < padding*4; i++ {
		fmt.Print(" ")
	}
	fmt.Printf(" (dir) name=%s size=%d:\n", dir.name, dir.size)
	for _, f := range dir.files {
		fmt.Print(">")
		for i := 0; i < padding*4; i++ {
			fmt.Print(" ")
		}
		fmt.Printf(" - (file) name=%s size=%d\n", f.name, f.size)
	}
	for _, d := range dir.dirs {
		printDir(d, padding+1)
	}
}

type directory struct {
	parent *directory
	name   string
	files  []file
	dirs   map[string]*directory
	size   uint64
}

type file struct {
	name string
	size uint64
}
