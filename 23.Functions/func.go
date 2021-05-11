package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Enter number : ")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	checkvalue(num1, num2)
}
func checkvalue(n1, n2 int) {
	if n1 > n2 {
		fmt.Println(n1, "is greater than", n2)
	} else if n2 < n1 {
		fmt.Println(n2, "is greater than", n1)
	} else {
		fmt.Println(n1, "equal to", n2)
	}
}
