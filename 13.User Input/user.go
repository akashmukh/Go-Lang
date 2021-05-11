package main

import (
	"fmt"
)

func main() {
	var name, surname string
	var age int
	fmt.Println("Enter name :")
	//take user input and add it to name
	fmt.Scanln(&name, &surname)
	fmt.Println("Enter age :")
	//take user input and add it to age
	fmt.Scanln(&age)
	fmt.Println("Your name : ", name, surname)
	fmt.Println("Your age : ", age)
}
