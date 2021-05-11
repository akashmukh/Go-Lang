package main

import (
	"fmt"
)

func main() {
	var number int
	var language string
	fmt.Println("WELCOME TO LANGUAGE SECTION")
	fmt.Println("1.English", "2.Bengali", "3.Hindi")
	fmt.Printf("Enter your choice :")
	fmt.Scanln(&number)
	switch number {
	case 1:
		language = "English"
	case 2:
		language = "Bengali"
	case 3:
		language = "Hindi"
	}
	fmt.Println("You Selected :", language)
}
