package main

import (
	"github.com/amirhlashgari/monsterslayer/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player", "Monster", ""
	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0 // special rounds have some special attacks, and it unlocks every 3 rounds

	interaction.ShowAvailableActions(isSpecialRound)
}

func endGame() {

}
