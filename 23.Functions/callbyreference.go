package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Enter values : ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("Before Swap : ")
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println("After Swap : ")
	fmt.Println(a, b)
}
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
