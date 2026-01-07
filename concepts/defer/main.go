//easy thing.

// multiple defer executed LIFO.

package main

import "fmt"

func main() {

	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")

	defer fmt.Println("d")
	defer fmt.Println("e")

	myDefer()
	defer fmt.Println("f")

	fmt.Println("this will come after abc")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
