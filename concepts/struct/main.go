package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Roll    int
	Present bool
	cgpa    float32
}

func main() {

	student1 := Student{"naman", 23, 64, true, 7.7}

	fmt.Println(student1.Age)
	fmt.Println(student1.Name)
	fmt.Println(student1.cgpa)

	fmt.Printf("the detais of student1 is %+v\n", student1)
}
