package main

import "fmt"

func main() {
	fmt.Println("abc")

	//var arr_name [size]type
	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	//fruitList[2] = "pineapple"
	fruitList[3] = "guava"

	fmt.Println(fruitList)

	var numbers = [5]int{1, 2, 3, 4}

	fmt.Println(numbers)
}
