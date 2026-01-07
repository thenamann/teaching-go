package main

import "fmt"


type Talker interface {
	Talk()
}

type Person struct { //write type first then its methods->good practice
	Name string
}

func(p *Person) Talk(){//implement the method
	fmt.Println("Person ", p.Name, " is talking")
}

func StartTalking(t Talker) {
	t.Talk()
}


func main() {
	p1 := Person{
		Name: "naman",
	}

	fmt.Println(p1.Name)
	StartTalking(&p1)
}


