package main

import "fmt"

func main() {

	week := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// method 1
	for d := 0; d < len(week); d++ {
		fmt.Println(week[d])
	}

	// method 2
	for i := range week {
		fmt.Println(week[i])
	}

	// method 3
	for index, week := range week {
		fmt.Println(index, week)
	}

	for _, days := range week {
		fmt.Printf("today is %v\n", days)
	}
}
