package interaction

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure" // to load third-party modules use --> "go mod tidy" | also can use "go get ..." to install package
)

type RoundData struct {
	Action          string
	PlayerAttackDmg int
	PlayerHealValue int
	MonsterAtackDmg int
	PlayerHealth    int
	MonsterHealth   int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()
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
	asciiFigure := figure.NewColorFigure("GAME OVER", "", "red", true)
	asciiFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won.\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("gamelog.txt")

	if err != nil {
		fmt.Println("Saving log file failed. Exiting...")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAtackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing log into a file failed. Exiting...")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")
}
