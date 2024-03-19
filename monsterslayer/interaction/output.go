package interaction

import "fmt"

type RoundData struct {
	Action          string
	PlayerAttackDmg int
	PlayerHealValue int
	MonsterAtackDmg int
	PlayerHealth    int
	MonsterHealth   int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER GAME")
	fmt.Println("Starting a new game...")
	fmt.Println("========================")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose yout action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Moster")
	fmt.Println("(2) Heal Yourself")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player special attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAtackDmg)
	fmt.Printf("Player Health: %v & Monster Health: %v\n", roundData.PlayerHealth, roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("-------------------------")
	fmt.Printf("%v won.\n", winner)
}
