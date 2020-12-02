package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const ()

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

func run(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})

	validPasswords := 0
	for _, line := range lines {
		if isValidPart2(line) {
			fmt.Printf("VALID LINE: %s\n", line)
			validPasswords += 1
		} else {
			fmt.Printf("INVALID LINE: %s\n", line)
		}
	}
	fmt.Printf("valid passwords: %d\n", validPasswords)
}

type state int

const (
	state_FirstNumber state = iota
	state_Dash
	state_SecondNumber
	state_Letter
	state_Word
)

func isValidPart2(line []byte) bool {

	var existPosition, notExistPosition int
	var letter byte

	s := state_FirstNumber

	start := 0
	for i, c := range line {
		switch s {
		case state_FirstNumber:
			if c == '-' {
				s = state_Dash
				existPosition = convertToInt(line[start:i])
			}
		case state_Dash:
			if c >= '0' && c <= '9' {
				s = state_SecondNumber
				start = i
			}
		case state_SecondNumber:
			if c == ' ' {
				s = state_Letter
				notExistPosition = convertToInt(line[start:i])
			}
		case state_Letter:
			if c >= 'a' && c <= 'z' {
				s = state_Word
				letter = c
			}
		case state_Word:
			if c >= 'a' && c <= 'z' {
				// validate password
				password := line[i:]
				if len(password) < notExistPosition {
					return false
				}

				if password[existPosition-1] == letter {
					return password[notExistPosition-1] != letter
				} else {
					return password[notExistPosition-1] == letter
				}
			}
		}
	}
	return false
}

func isValidPart1(line []byte) bool {

	var min, max int
	var letter byte

	s := state_FirstNumber

	start := 0
	for i, c := range line {
		switch s {
		case state_FirstNumber:
			if c == '-' {
				s = state_Dash
				min = convertToInt(line[start:i])
			}
		case state_Dash:
			if c >= '0' && c <= '9' {
				s = state_SecondNumber
				start = i
			}
		case state_SecondNumber:
			if c == ' ' {
				s = state_Letter
				max = convertToInt(line[start:i])
			}
		case state_Letter:
			if c >= 'a' && c <= 'z' {
				s = state_Word
				letter = c
			}
		case state_Word:
			if c >= 'a' && c <= 'z' {
				// validate password
				count := 0
				if c == letter {
					count += 1
				}
				for _, c := range line[i+1:] {
					if c == letter {
						count += 1
					}
				}
				return min <= count && count <= max
			}
		}
	}
	return false
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
