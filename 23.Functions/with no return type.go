package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter number : ")
	fmt.Scan(&num)
	sqr(num)
}
func sqr(n int) {
	result := n * n
	fmt.Println("Square : ", result)
}
