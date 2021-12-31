package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var (
	//go:embed test1.txt
	test1 []byte

	//go:embed input1.txt
	input1 []byte
)

func main() {
	part1(test1)
	part1(input1)
	part2(test1)
	part2(input1)
}

func part1(data []byte) {
	player1 := -1
	player2 := -1
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		vs := bytes.Split(line, []byte{':'})
		if player1 == -1 {
			player1 = convertToInt(vs[1])
		} else if player2 == -1 {
			player2 = convertToInt(vs[1])
		}
	}
	fmt.Printf("Player 1: %d, Player 2: %d\n", player1, player2)

	dice := 1
	rolled := 0
	roll := func() int {
		d := dice
		dice++
		if dice == 101 {
			dice = 1
		}
		rolled++
		return d
	}

	score1 := 0
	score2 := 0

	for i := 0; ; i++ {
		{
			//fmt.Printf("%d) player 1: pos=%d, score=%d\n", i, player1, score1)
			rolled := 0
			for i := 0; i < 3; i++ {
				r := roll()
				rolled += r
			}
			player1 += rolled

			if mod := player1 % 10; mod == 0 {
				player1 = 10
			} else {
				player1 %= 10
				if player1 == 0 {
					player1++
				}
			}
			score1 += player1
			if score1 >= 1000 {
				break
			}
			//fmt.Printf("%d) player 1: pos=%d, score=%d, rolled=%d\n", i, player1, score1, rolled)
		}
		{
			//fmt.Printf("%d) player 2: pos=%d, score=%d\n", i, player2, score2)
			rolled := 0
			for i := 0; i < 3; i++ {
				r := roll()
				rolled += r
			}
			player2 += rolled
			if mod := player2 % 10; mod == 0 {
				player2 = 10
			} else {
				player2 %= 10
				if player2 == 0 {
					player2++
				}
			}
			score2 += player2
			if score2 >= 1000 {
				break
			}
			//fmt.Printf("%d) player 2: pos=%d, score=%d, rolled=%d\n", i, player2, score2, rolled)
		}

	}
	/*
		fmt.Println(rolled)
		fmt.Println(player1, score1)
		fmt.Println(player2, score2)
	*/
	lowest := score1
	if score1 > score2 {
		lowest = score2
	}
	fmt.Printf("Part 1: %d\n", rolled*lowest)
}

type game struct {
	player1, player2 int
	score1, score2   int
}

func (g game) NewGame() game {
	return game{
		player1: g.player1,
		player2: g.player2,
		score1:  g.score1,
		score2:  g.score2,
	}
}

func (g game) key() [4]int {
	return [4]int{
		g.player1,
		g.player2,
		g.score1,
		g.score2,
	}
}

var (
	cache       = map[[4]int][2]int{}
	player1wins uint64
	player2wins uint64
	gamesPlayed uint64
)

func part2(data []byte) {
	player1 := -1
	player2 := -1
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		vs := bytes.Split(line, []byte{':'})
		if player1 == -1 {
			player1 = convertToInt(vs[1])
		} else if player2 == -1 {
			player2 = convertToInt(vs[1])
		}
	}
	fmt.Printf("Player 1: %d, Player 2: %d\n", player1, player2)

	playersWins := play(game{player1: player1, player2: player2})
	fmt.Println(playersWins)
	winner := playersWins[0]
	if playersWins[1] > winner {
		winner = playersWins[1]
	}
	fmt.Printf("Part 2: %d\n", winner)
}

func play(g game) [2]int {
	if g.score1 >= 21 {
		return [2]int{1, 0}
	} else if g.score2 >= 21 {
		return [2]int{0, 1}
	}
	if wins, ok := cache[g.key()]; ok {
		return wins
	}

	wins := [2]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				p1 := g.player1 + i + j + k
				if mod := p1 % 10; mod == 0 {
					p1 = 10
				} else {
					p1 %= 10
					if p1 == 0 {
						p1++
					}
				}
				s1 := g.score1 + p1
				w := play(game{
					player1: g.player2, player2: p1,
					score1: g.score2, score2: s1,
				})
				wins[0] += w[1]
				wins[1] += w[0]
			}
		}
	}
	cache[g.key()] = wins
	return wins
}

func convertToInt(val []byte) int {
	iVal := 0
	startBase := 1
	for i := len(val) - 1; i >= 0; i-- {
		if val[i] < '0' || val[i] > '9' {
			continue
		}
		iVal += startBase * int(val[i]-'0')
		startBase *= 10
	}
	return iVal
}
