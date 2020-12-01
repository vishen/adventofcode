package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

const (
	lookingFor = 2020
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	// part1(data)
	part2(data)
}

func part2(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	reportMap := make(map[int]struct{}, len(lines))
	reportData := make([]int, len(lines))

	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		val := convertToInt(l)
		reportMap[val] = struct{}{}
		reportData[i] = val
	}

	// Sort the data in descending orders since it will be more likely that one of the larger numbers
	// is apart of the 2020.
	sort.Slice(reportData, func(i, j int) bool {
		return reportData[i] > reportData[j]
	})

	for i, number1 := range reportData {
		found := false
		for _, number2 := range reportData[i:] {
			want := lookingFor - (number1 - number2)
			if want <= 0 {
				continue
			}

			if _, ok := reportMap[want]; !ok {
				continue
			}

			fmt.Printf("found numbers %d, %d and %d == 2020. Multiplied = %d\n", number1, number2, want, number1*number2*want)
			found = true
			break
		}
		if found {
			break
		}
	}

}

func part1(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	reportMap := make(map[int]struct{}, len(lines))
	reportData := make([]int, len(lines))

	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		val := convertToInt(l)
		reportMap[val] = struct{}{}
		reportData[i] = val
	}

	// Sort the data in descending orders since it will be more likely that one of the larger numbers
	// is apart of the 2020.
	sort.Slice(reportData, func(i, j int) bool {
		return reportData[i] > reportData[j]
	})

	for _, number := range reportData {
		want := lookingFor - number
		if want < 0 {
			continue
		}

		if _, ok := reportMap[want]; !ok {
			continue
		}

		fmt.Printf("found numbers %d and %d == 2020. Multiplied = %d\n", number, want, number*want)
		break
	}

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
