package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Enter two numbers : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	sum, sub := operation(a, b)
	fmt.Println("Sum = ", sum, "Sub = ", sub)
}
func operation(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
