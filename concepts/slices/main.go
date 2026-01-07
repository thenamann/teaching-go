package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Email  string
	Status bool
	pvt    string
}

func main() {

	naman := Person{"naman", 23, "naman@gmail.com", true, "cricket"}
	naman.GetStatus()
}

// GetStatus is initialised with capital letter so it can exported.
func (p1 Person) GetStatus() {
	fmt.Println("abc", p1.Status)
	p1.Age = 25
	fmt.Println(p1.Age)
}

