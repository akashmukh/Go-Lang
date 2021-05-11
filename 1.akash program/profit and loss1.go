package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Print("Enter Cost Price : ")
	fmt.Scan(&a)
	fmt.Print("Enter Selling Price : ")
	fmt.Scan(&b)
	percentage(a, b)
}
func percentage(cp, sp float64) {
	var profit, loss float64
	if sp > cp {
		profit = ((sp - cp) / cp * 100)
		fmt.Println("Profit percentage : ", profit, "%")
	} else if cp > sp {
		loss = ((cp - sp) / cp * 100)
		fmt.Println("Loss Percentage : ", loss, "%")
	} else {
		fmt.Println("No Profit No Loss")
	}
}
