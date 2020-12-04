package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

type state int

const (
	state_KEY state = iota
	state_VALUE
)

func run(data []byte) {
	validPassports := 0

	current := map[string]string{}
	var currentKey string
	startKey := 0
	startValue := 0
	s := state_KEY
	for i, c := range data {
		switch s {
		case state_KEY:
			switch c {
			case ':':
				currentKey = string(data[startKey:i])
				s = state_VALUE
				startValue = i + 1
			case '\n':
				if isValidPassport(current) {
					validPassports += 1
				}
				current = map[string]string{}
			default:
				if c >= 'a' && c <= 'z' {
					if startKey == -1 {
						startKey = i
					}
				}
			}
		case state_VALUE:
			if c == ' ' || c == '\n' {
				s = state_KEY
				startKey = -1
				current[currentKey] = string(data[startValue:i])
			}
		}
	}

	if len(current) > 0 && isValidPassport(current) {
		validPassports += 1
	}

	fmt.Printf("%d valid passports\n", validPassports)
}

func isValidPassport(passport map[string]string) bool {

	requiredFieldsSet := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}
	for k, v := range passport {

		/*
		   byr (Birth Year) - four digits; at least 1920 and at most 2002.
		   iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		   eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		   hgt (Height) - a number followed by either cm or in:
		   If cm, the number must be at least 150 and at most 193.
		   If in, the number must be at least 59 and at most 76.
		   hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		   ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		   pid (Passport ID) - a nine-digit number, including leading zeroes.
		   cid (Country ID) - ignored, missing or not.
		*/

		// TODO: should only do this when needed, but easier for now.
		vInt, _ := strconv.Atoi(v)
		switch k {
		case "byr":
			if len(v) != 4 || vInt < 1920 || vInt > 2002 {
				return false
			}
		case "iyr":
			if len(v) != 4 || vInt < 2010 || vInt > 2020 {
				return false
			}
		case "eyr":
			if len(v) != 4 || vInt < 2020 || vInt > 2030 {
				return false
			}
		case "ecl":
			switch v {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				// expected
			default:
				return false
			}
		case "pid":
			if len(v) != 9 || vInt < 0 {
				return false
			}
		case "hcl":
			// a # followed by exactly six characters 0-9 or a-f.
			if len(v) != 7 || v[0] != '#' || isAlphaNumericString(v[1:]) {
				return false
			}
		case "hgt":
			/*
				hgt (Height) - a number followed by either cm or in:
				If cm, the number must be at least 150 and at most 193.
				If in, the number must be at least 59 and at most 76.
			*/
			if len(v) < 4 || len(v) > 5 {
				return false
			}
			switch v[len(v)-2:] {
			case "cm":
				i, err := strconv.Atoi(v[:len(v)-2])
				if err != nil {
					return false
				}
				if i < 150 || i > 193 {
					return false
				}
			case "in":
				i, err := strconv.Atoi(v[:len(v)-2])
				if err != nil {
					return false
				}
				if i < 59 || i > 79 {
					return false
				}
			default:
				return false
			}
		}
		requiredFieldsSet[k] = true
	}

	for _, r := range requiredFieldsSet {
		if !r {
			return false
		}
	}
	return true
}

func isAlphaNumericString(val string) bool {
	for _, c := range val {
		if !isAlphaNumeric(c) {
			return false
		}
	}
	return false
}

func isAlphaNumeric(c rune) bool {
	return c >= 'a' && c <= 'z' ||
		// c >= 'A' && c <= 'Z' ||
		c >= '0' && c <= '9'
}
