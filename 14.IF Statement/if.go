package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("Enter age :")
	fmt.Scanln(&age)
	//check the boolean condition using if
	if age > 18 {
		/*if the condition is true then print the following*/
		fmt.Println("You are adult")
	}
}
