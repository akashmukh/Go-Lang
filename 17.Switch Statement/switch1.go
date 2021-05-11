package main

import (
	"fmt"
)

func main() {
	var code int
	var subject string
	fmt.Println("1.English", "2.Bengali", "3.Math", "4.History", "5.Geography")
	fmt.Println("Choose your Subject :")
	fmt.Scanln(&code)
	switch code {
	case 1:
		subject = "English"
	case 2:
		subject = "Bengali"
	case 3:
		subject = "Math"
	case 4:
		subject = "History"
	case 5:
		subject = "Geography"
	default:
		subject = "Your not Eligilbe to select any Subject"

	}
	fmt.Println("You Choose :", subject)
}
