package main

import (
	"fmt"
	"math/rand"

	"github.com/asmriti/guesser-game/constants"
)

func main() {
	random := rand.Intn(100)

	fmt.Println(random)

	// From here
	fmt.Println("You have 3 chances to guess the number.")
	for i := 0; i < 3; i++ {
		var input int
		fmt.Print("Guess number between 0 to 99: ")
		fmt.Scanln(&input)

		if input == random {
			fmt.Println("Yaayayy!! Your guess is correct")
			fmt.Printf("You got %d points.\n", constants.TOTAL_NO_OF_CHANCES-i)
			break
		} else {
			// getFailureMessage
		}

		if random > input {
			fmt.Printf("Your guess is lower. You have %d chances remaining.\n", constants.TOTAL_NO_OF_CHANCES-i-1)
		} else {
			fmt.Printf("Your guess is higher. You have %d chances remaining.\n", constants.TOTAL_NO_OF_CHANCES-i-1)
		}

		if i == 2 {
			fmt.Println("Your 3 chances are over.")
		}
	}

	// to here
}

// IF user guesses on the first attempt, show a message saying he got 3 stars
// if second attempt, give him 2 stars
// if third attempt, give him 1 star

/*

total no of chances, t = 3

i = 0, p = 3
i = 1, p = 2
i = 2, p = 1


p = t - i

*/

/*

total no of chances, t = 3

i = 0, c = 2
i = 1, c = 1
i = 2, c = 0


p = t - i - 1

*/
