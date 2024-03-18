package interaction

import "fmt"

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
