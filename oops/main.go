package main
import "fmt"

type CanEat interface { //interface
	Eat()
}
type CanWalk interface { //interface 2
	Walk()
}
func (a Student) Walk(){ //implement the method
	fmt.Println("A student is walking")
}
func StartWalking(w CanWalk){
	w.Walk()
}

func (s Student) Eat() { // Implement method 2
	fmt.Println("A student is eating")

}

func StartEating (e CanEat ) {
	e.Eat()
}

type Person struct{
	name string
	gender bool
	age int
}

type  Student struct{
	Person
	roll int 
}

func main () {
	fmt.Println("hi ")

	s := Student{
		Person: Person{
			name: "naman",
			gender: true,
			age: 23,
		},
		roll: 64,
	}
	fmt.Println(s.name)
	StartEating(s)
	StartWalking(s)
}