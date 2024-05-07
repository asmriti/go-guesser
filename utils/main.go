package utils

import (
	"fmt"

	"github.com/asmriti/guesser-game/constants"
)

// This should extract the else if and else logic of our game
func GetErrorMessage(random, input, i int) {
	if IsGameOver(i) {
		PrintGameOverMessage()
		return
	}

	remainingChances := ComputeRemainingChances(i)

	str := "chance"
	if remainingChances > 1 {
		str = str + "s"
	}

	if random > input {
		fmt.Printf("Your guess is lower. You have %d %s remaining.\n", remainingChances, str)
	} else {
		fmt.Printf("Your guess is higher. You have %d %s remaining.\n", remainingChances, str)
	}
}

func GetSuccessMessage(random, input, i int) {
	fmt.Println("***************")
	fmt.Println("Yaayayy!! Your guess is correct.")
	totalPoints := CalculatePoints(i)
	fmt.Printf("You got %d points.\n", totalPoints)
}

func IsGameOver(iteration int) bool {
	return iteration == constants.TOTAL_NO_OF_CHANCES-1
}

func PrintGameOverMessage() {
	fmt.Println("=====================")
	fmt.Printf("You lost. Your %d chances are over. \n", constants.TOTAL_NO_OF_CHANCES)
}

// This function calculates remaining chances
func ComputeRemainingChances(i int) int {
	return constants.TOTAL_NO_OF_CHANCES - i - 1
}

// This function caluclates total points
func CalculatePoints(i int) int {
	return constants.TOTAL_NO_OF_CHANCES - i
}
