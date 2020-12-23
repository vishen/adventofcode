package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
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

	_, player := game(1, player1, player2)
	fmt.Println("score:", score(player))

}

func game(number int, player1, player2 []int) (int, []int) {

	moves := map[string]bool{}

	for i := 0; ; i++ {
		fmt.Println("Game", number, "round", i+1)
		fmt.Println("Player1: ", player1)
		fmt.Println("Player2: ", player2)
		fmt.Println()

		if len(player1) == 1 || len(player2) == 1 {
			if player1[0] > player2[0] {
				player1, player2 = update(player1, player2)
				fmt.Println("Player 1 won (only one card left; top card > p2 top card)")
				return 1, player1
			} else {
				player2, player1 = update(player2, player1)
				fmt.Println("Player 2 won (only one card left; top card > p1 top card)")
				return 2, player2
			}
		}

		if _, ok := moves[h(player1[1:], player2[1:])]; ok {
			fmt.Println("Player 1 won (seen game before)")
			player1, player2 = update(player1, player2)
			return 1, player1
		}
		if _, ok := moves[h(player2[1:], player1[1:])]; ok {
			fmt.Println("Player 1 won (seen game before)")
			player1, player2 = update(player1, player2)
			return 1, player1
		}

		if rule1(player1, player2) {
			fmt.Println("playing another game...")
			p1Max := len(player1)
			if player1[0] < len(player1) {
				p1Max = player1[0] + 1
			}
			p2Max := len(player2)
			if player2[0] < len(player2) {
				p2Max = player2[0] + 1
			}
			player1Copy := make([]int, len(player1))
			copy(player1Copy, player1)
			player2Copy := make([]int, len(player2))
			copy(player2Copy, player2)
			number, _ := game(number+1, player1Copy[1:p1Max], player2Copy[1:p2Max])
			if number == 1 {
				player1, player2 = update(player1, player2)
			} else {
				player2, player1 = update(player2, player1)
			}
		} else {
			if player1[0] > player2[0] {
				player1, player2 = update(player1, player2)
			} else {
				player2, player1 = update(player2, player1)
			}
		}

		if len(player1) == 0 {
			fmt.Println("Player 2 won (no cards for player1)")
			return 2, player2
		} else if len(player2) == 0 {
			fmt.Println("Player 1 won (no cards for player2)")
			return 1, player1
		}

		moves[h(player1, player2)] = true
		moves[h(player2, player1)] = true
	}
	// return 1, player1
}

func h(p1, p2 []int) string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%v -- %v", p1, p2))
	s := fmt.Sprintf("%x", h.Sum(nil))
	return s
}

func rule1(p1, p2 []int) bool {
	return len(p1)-1 >= p1[0] && len(p2)-1 >= p2[0]
}

/*
func turn(p1, p2 []int) ([]int, []int) {
	if p1[0] > p2[0] {
		p1 = append(p1[1:], p1[0], p2[0])
		p2 = p2[1:]
	} else {
		p2 = append(p2[1:], p2[0], p1[0])
		p1 = p1[1:]
	}
	return p1, p2
}*/

func update(winner, loser []int) ([]int, []int) {
	winner = append(winner[1:], winner[0], loser[0])
	loser = loser[1:]
	return winner, loser
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
