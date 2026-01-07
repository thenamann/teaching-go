package main

import "fmt"

func main() {
	num := 10

	if num < 5 {
		fmt.Println("less than 10")
	} else if num == 10 {
		fmt.Println("equal to 10")
	} else {
		fmt.Println("greater than 10")
	}

	if 9%2 == 0 {
		println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	// here num := 11 is assignment, and num < 10 is the condition
	if num := 11; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("not less than 10")
	}
}
