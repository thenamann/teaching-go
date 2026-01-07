package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "these are the content which should go in the file"

	file, err := os.Create("./mylcogofile.txt") //creates the file

	if err != nil {
		panic(err)
	} else {
		fmt.Println("file created")
	}

	length, err := io.WriteString(file, content) //write the file
	checkNilError(err)

	fmt.Println("the length is: ", length)

	readFile("./mylcogofile.txt")

	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Printf("text data in binary is %v\n", databyte)
	fmt.Printf("text data is:  %v\n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
