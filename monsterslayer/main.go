package main

import (
	"github.com/amirhlashgari/monsterslayer/actions"
	"github.com/amirhlashgari/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player", "Monster", ""
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0 // special rounds have some special attacks, and it unlocks every 3 rounds

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealVal int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealVal = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}
	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()
	roundData := interaction.RoundData{
		Action:          userChoice,
		PlayerAttackDmg: playerAttackDmg,
		PlayerHealValue: playerHealVal,
		MonsterAtackDmg: monsterAttackDmg,
		PlayerHealth:    playerHealth,
		MonsterHealth:   monsterHealth,
	}
	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}

/**
 * NOTES:
 * -----------
 * to compile our Go code and convert it to an executable file ---> "go build"
 * in above case there is no need that user have go installed on system and returns a .exe file to use
 */
