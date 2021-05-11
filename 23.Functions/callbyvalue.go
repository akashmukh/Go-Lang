package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Enter Numbers: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println("Before Swap : ")
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println("After Swap : ")
	fmt.Println(a, b)
}
func swap(x, y int) (int, int) {
	return y, x
}
