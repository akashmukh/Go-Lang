package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Print("Enter Number : ")
	fmt.Scanln(&input)
	if sqrt := math.Sqrt(float64(input)); sqrt < 5 {
		fmt.Println("Less than 5")
	} else {
		fmt.Println("Greater than 5")
	}
}
