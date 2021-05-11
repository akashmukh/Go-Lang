package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter Number : ")
	fmt.Scan(&num)
	if num%2 != 0 {
		fmt.Printf("%d is a Odd number", num)
	} else {
		fmt.Printf("%d is a Even number", num)
	}
}
