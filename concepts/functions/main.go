package main

import "fmt"

func main() {

	prosum := proAdder(1, 2, 3, 4)
	fmt.Println(prosum)

	sum1 := adder(100, 200)
	println(sum1)
}

func proAdder(nums ...int) int {

	total := 0

	for _, val := range nums {
		total += val
	}

	return total
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}
