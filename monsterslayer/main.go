package main

import (
	"github.com/amirhlashgari/monsterslayer/actions"
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
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}
	actions.AttackPlayer()

	playerHealth, monsterHealth = actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {

}
