package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/asmriti/guesser-game/constants"
	"github.com/asmriti/guesser-game/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(constants.MAX_NUMBER)

	fmt.Println(random)

	fmt.Printf("You have %d chances to guess the number. \n", constants.TOTAL_NO_OF_CHANCES)
	for i := 0; i < constants.TOTAL_NO_OF_CHANCES; i++ {
		var input int
		fmt.Printf("Guess number between 0 to %d: ", constants.MAX_NUMBER-1)
		fmt.Scanln(&input)

		if input == random {
			utils.GetSuccessMessage(random, input, i)
			break
		} else {
			utils.GetErrorMessage(random, input, i)
		}
	}
}
