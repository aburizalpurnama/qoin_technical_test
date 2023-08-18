package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestPlayDiceGame(t *testing.T) {
	PlayDiceGame(4, 6)
}

func PlayDiceGame(countPlayer, countDice int) {
	fmt.Printf("Pemain = %d, Dadu = %d\n", countPlayer, countDice)
	fmt.Println("==================")

	players := generatePlayers(countPlayer, countDice)
	round := 1
	for getCountPlayersWithDice(players) > 1 {
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		playDicesOfPlayers(players)

		printStatusOfPlayers(players)

		evaluateDicesOfPlayers(players)

		fmt.Println("Setelah evaluasi:")
		printStatusOfPlayers(players)

		fmt.Println("==================")
		round++
	}

	printResult(findHigestPointIndex(players), findLastPlayerIndex(players))
}

type player struct {
	point        int
	countOfDices int
	dices        []int
}

func newPlayer(countDices int) (p player) {
	p.countOfDices = countDices
	return p
}

func (player *player) evaluateDices() (movedDice []int) {
	movedDice = nil
	var updatedDices []int

	for _, v := range player.dices {
		if v == 1 {
			movedDice = append(movedDice, v)
			player.countOfDices--
		} else if v == 6 {
			player.point++
			player.countOfDices--
		} else {
			updatedDices = append(updatedDices, v)
		}
	}

	player.dices = updatedDices

	return
}

func (player *player) playDices() {
	var dices []int
	maxNumberOfDice := 6
	minNumberOfDice := 1
	randomInterval := 10 * time.Microsecond
	for i := 0; i < player.countOfDices; i++ {
		rand.Seed(time.Now().UnixNano())
		dices = append(dices, rand.Intn(maxNumberOfDice)+minNumberOfDice)
		time.Sleep(randomInterval)
	}
	player.dices = dices
}

func generatePlayers(countPlayers, countDices int) (players []player) {
	for i := 0; i < countPlayers; i++ {
		players = append(players, newPlayer(countDices))
	}
	return
}

func evaluateDicesOfPlayers(players []player) {
	movedDicesMap := make(map[int][]int, 0)

	for i := 0; i < len(players); i++ {
		if i < len(players)-1 {
			movedDices := players[i].evaluateDices()
			if len(movedDices) > 0 {
				movedDicesMap[i+1] = movedDices
			}
		} else {
			movedDices := players[i].evaluateDices()
			if len(movedDices) > 0 {
				movedDicesMap[0] = movedDices
			}
		}
	}

	for k, v := range movedDicesMap {
		if players[k].countOfDices > 0 {
			players[k].dices = append(players[k].dices, v...)
			players[k].countOfDices += len(v)
		}
	}
}

func printStatusOfPlayers(players []player) {
	for i, v := range players {
		fmt.Printf("\tPemain #%d (%d): %v\n", i+1, v.point, v.dices)
	}
}

func playDicesOfPlayers(players []player) {
	for i := 0; i < len(players); i++ {
		if players[i].countOfDices > 0 {
			players[i].playDices()
		}
	}
}

func findHigestPointIndex(players []player) []int {
	var winnerIndexes []int
	higestPoint := 0
	for _, v := range players {
		if v.point > higestPoint {
			higestPoint = v.point
		}
	}

	for i, v := range players {
		if v.point == higestPoint {
			winnerIndexes = append(winnerIndexes, i)
		}
	}
	return winnerIndexes
}

func findLastPlayerIndex(player []player) int {
	lastPlayerIndex := -1
	for i, v := range player {
		if v.countOfDices > 0 {
			lastPlayerIndex = i
		}
	}

	return lastPlayerIndex
}

func getCountPlayersWithDice(players []player) (result int) {
	for _, v := range players {
		if v.countOfDices > 0 {
			result++
		}
	}
	return
}

func printResult(highestPlayerIndexes []int, lastPlayerIndex int) {
	if lastPlayerIndex != -1 {
		fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu\n", lastPlayerIndex+1)
	} else {
		fmt.Print("Game berakhir karena semua pemain tidak memiliki dadu\n")
	}

	if len(highestPlayerIndexes) == 1 {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", highestPlayerIndexes[0]+1)
	} else {
		builder := strings.Builder{}
		builder.WriteString("Game dimenangkan oleh pemain ")
		for i, v := range highestPlayerIndexes {
			if i < len(highestPlayerIndexes)-1 {
				builder.WriteString(fmt.Sprintf("#%d, ", v+1))
			} else {
				builder.WriteString(fmt.Sprintf("#%d ", v+1))
			}
		}
		builder.WriteString("karena memiliki poin lebih banyak dari pemain lainnya.\n")
		fmt.Print(builder.String())
	}
}
