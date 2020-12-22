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

	player1 := []int{}
	player2 := []int{}

	p1 := true

	for _, line := range bytes.Split(data, []byte{'\n'})[1:] {
		if len(line) == 0 {
			continue
		}

		if bytes.Equal(line, []byte("Player 2:")) {
			p1 = false
			continue
		}

		if p1 {
			player1 = append(player1, convertToInt(line))
		} else {
			player2 = append(player2, convertToInt(line))
		}
	}

	max := 10000000

	for i := 0; i < max; i++ {
		fmt.Println("Player1: ", player1)
		fmt.Println("Player2: ", player2)
		fmt.Println()
		player1, player2 = turn(player1, player2)

		if len(player1) == 0 {
			fmt.Println("Player 2 won")
			fmt.Println(score(player2))
			break
		} else if len(player2) == 0 {
			fmt.Println("Player 1 won")
			fmt.Println(score(player1))
			break
		}
	}
}

func turn(p1, p2 []int) ([]int, []int) {
	if p1[0] > p2[0] {
		p1 = append(p1[1:], p1[0], p2[0])
		p2 = p2[1:]
	} else {
		p2 = append(p2[1:], p2[0], p1[0])
		p1 = p1[1:]
	}
	return p1, p2
}

func score(p []int) int {
	total := 0
	for i, c := range p {
		total += c * (len(p) - i)
	}
	return total
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
