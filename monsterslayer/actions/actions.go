package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth = currentMonsterHealth - damageValue
}

func HealPlayer() {
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)
	healDiff := PLAYER_HEALTH - currentPlayerHealth

	if healDiff >= healValue { // to avoid player health go over 100
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= damageValue
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
