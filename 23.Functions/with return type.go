package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter Number : ")
	fmt.Scan(&num)
	result := sqr(num)
	fmt.Println("Square : ", result)
}
func sqr(n int) int {
	result := n * n
	return result
}
