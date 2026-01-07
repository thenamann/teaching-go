package main

import "fmt"

func main() {

	fmt.Println(".")

	lang := make(map[string]string)

	lang["js"] = "javascript"
	lang["py"] = "python"
	lang["go"] = "golang"

	for key, value := range lang {
		fmt.Printf("for key %v, the value is %v\n", key, value)
	}
}
