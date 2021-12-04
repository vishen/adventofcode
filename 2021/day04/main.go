package main

import (
	"bytes"
	"fmt"

	_ "embed"
)

const (
	SIZE = 5
)

var (
	//go:embed test.txt
	test []byte

	//go:embed part1.txt
	p1 []byte

	//go:embed part2.txt
	p2 []byte
)

func main() {
	part1(test)
	part1(p1)

	part2(test)
	part2(p2)
}

func part1(data []byte) {
	var (
		numbersToCall []int
		boards        [][]int
		curBoard      []int
	)

	for i, line := range bytes.Split(data, []byte{'\n'}) {
		if i == 0 {
			numbersToCall = numbers(bytes.Split(line, []byte{','}))
			continue
		}
		if len(line) == 0 {
			if len(curBoard) > 0 {
				boards = append(boards, curBoard)
				curBoard = []int{}
			}
			continue
		}
		curBoard = append(curBoard, numbers(bytes.Split(line, []byte{' '}))...)
	}

	if len(curBoard) > 0 {
		boards = append(boards, curBoard)
		curBoard = []int{}
	}

	for _, n := range numbersToCall {
		for bi, board := range boards {
			for i, b := range board {
				if b == n {
					board[i] = -1
				}
			}
			if total, ok := checkWin(board); ok {
				fmt.Printf("Part 1: board %d wins, total=%d * n=%d == %d\n", bi, total, n, total*n)
				return
			}
		}
	}
}

func part2(data []byte) {
	var (
		numbersToCall []int
		boards        [][]int
		curBoard      []int
	)

	for i, line := range bytes.Split(data, []byte{'\n'}) {
		if i == 0 {
			numbersToCall = numbers(bytes.Split(line, []byte{','}))
			continue
		}
		if len(line) == 0 {
			if len(curBoard) > 0 {
				boards = append(boards, curBoard)
				curBoard = []int{}
			}
			continue
		}
		curBoard = append(curBoard, numbers(bytes.Split(line, []byte{' '}))...)
	}

	if len(curBoard) > 0 {
		boards = append(boards, curBoard)
		curBoard = []int{}
	}

	won := make(map[int]struct{})
	for _, n := range numbersToCall {
		for bi, board := range boards {
			for i, b := range board {
				if b == n {
					board[i] = -1
				}
			}
			if total, ok := checkWin(board); ok {
				won[bi] = struct{}{}
				if len(won) == len(boards) {
					fmt.Printf("Part 2: board %d wins last, total=%d * n=%d == %d\n", bi, total, n, total*n)
					return
				}
			}
		}
	}
}

func checkWin(board []int) (int, bool) {
	for i := 0; i < SIZE; i++ {
		markedRow := 0
		markedColumn := 0
		for j := 0; j < SIZE; j++ {
			if board[(i*SIZE)+j] == -1 {
				markedRow++
			}
			if board[i+(SIZE*j)] == -1 {
				markedColumn++
			}
		}
		if markedColumn == SIZE || markedRow == SIZE {
			total := 0
			for _, b := range board {
				if b == -1 {
					continue
				}
				total += b
			}
			return total, true
		}
	}
	return 0, false
}

func printBoard(board []int) {
	for i, b := range board {
		fmt.Printf("%d ", b)
		if i%SIZE == 4 {
			fmt.Println()
		}
	}
}

func numbers(data [][]byte) []int {
	var result []int
	for _, d := range data {
		if len(d) == 0 {
			continue
		}
		result = append(result, ints(d))
	}
	return result
}

func ints(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
