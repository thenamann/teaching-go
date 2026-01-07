package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("lets start switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("number is 1")
	case 2:
		fmt.Println("number is 2")
	case 3:
		fmt.Println("number is 3")
		fallthrough
	case 4:
		fmt.Println("number is 4")
		fallthrough
	case 5:
		fmt.Println("number is 5")
	case 6:
		fmt.Println("number is 6")
	}
}
