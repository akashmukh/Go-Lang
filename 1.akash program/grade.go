package main

import (
	"fmt"
)

func main() {
	var s string
	var n int
	fmt.Println("Type Your Subject and Number to know your Grade")
	fmt.Print("Enter Subject : ")
	fmt.Scan(&s)
	fmt.Print("Enter No :")
	fmt.Scan(&n)
	if n > 10 && n <= 30 {
		fmt.Println("Grade D")
		if n >= 30 && n <= 50 {
			fmt.Println("Grade C")
		}
	} else if n >= 50 && n <= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade A")
	}
}
