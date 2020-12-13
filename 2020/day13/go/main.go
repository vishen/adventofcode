package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	run(data)
}

func run(data []byte) {
	dataSplit := bytes.Split(data, []byte{'\n'})

	startTime := convertToInt(dataSplit[0])

	buses := []int{}
	for _, b := range bytes.Split(dataSplit[1], []byte{','}) {
		if b[0] == 'x' {
			continue
		}

		buses = append(buses, convertToInt(b))
	}

	fmt.Println(startTime, buses)

	nextBus := buses[0]
	min := buses[0] - (startTime % buses[0])
	for _, b := range buses[1:] {
		diff := b - (startTime % b)
		if diff < min {
			nextBus = b
			min = diff
		}
	}

	fmt.Println(nextBus, min, nextBus*min)
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
