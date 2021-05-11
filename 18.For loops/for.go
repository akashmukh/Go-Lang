package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter number : ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Loop", i)
	}
}
