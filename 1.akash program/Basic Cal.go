package main

import (
	"fmt"
	"os"
)

func main() {
	var num, num1 float64
	var operator string
	var choice string
calculator:
	fmt.Print("Enter Number : ")
	fmt.Scanln(&num)
	fmt.Print("Enter Operator : ")
	fmt.Scanln(&operator)
	fmt.Print("Enter Another Number : ")
	fmt.Scanln(&num1)
	switch operator {
	case "+":
		fmt.Println(num + num1)
	case "-":
		fmt.Println(num - num1)
	case "*":
		fmt.Println(num * num1)
	case "/":
		fmt.Println(num / num1)
	default:
		fmt.Println("Invalid Operation")
	}
	fmt.Print("Enter choice : ")
	fmt.Scanln(&choice)
	if choice == "no" {
		os.Exit(0)
	}
	goto calculator
}
