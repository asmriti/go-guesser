package main

import (
	"fmt"
	"math/rand"
)

func main() {

	input := 5
	random := rand.Intn(10)

	if input == random {
		fmt.Println("Yaayayy!! Your guess is correct")
	} else {
		fmt.Println("Your guess is wrong.")
	}

	fmt.Println("The generated number is", random)

}