package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Enter age :")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are not adult")
	}
}
