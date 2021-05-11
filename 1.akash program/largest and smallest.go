package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("Enter Three Number : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a > b && a > c {
		fmt.Println("The Greatest Number is : ", a)
	} else if b > a && b > c {
		fmt.Println("The Greatest Number is : ", b)
	} else {
		fmt.Println("The Greatest Number is : ", c)
	}
	if a < b && a < c {
		fmt.Println("The Smallest Number is : ", a)
	} else if b < a && b < c {
		fmt.Println("The Smallest Number is : ", b)
	} else {
		fmt.Println("The Smallest Number is : ", c)
	}
}
