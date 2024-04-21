package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= damageValue
	return damageValue
}

func HealPlayer() int {
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)
	healDiff := PLAYER_HEALTH - currentPlayerHealth

	if healDiff >= healValue { // to avoid player health go over 100
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healDiff
	}
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= damageValue
	return damageValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	/**
	 * we want a random number betwwen "min" & "max"
	 * but problem is Intn function takes only 1 parameter
	 * so we pass "max - min" to functions param
	 * and then subtract "min" from it to have the desired random number
	 */
	return randGenerator.Intn(max-min) + min
}
