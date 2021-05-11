package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Println("Enter Cost Price : ")
	fmt.Scan(&a)
	fmt.Println("Enter Selling Price : ")
	fmt.Scan(&b)
	percentage(a, b)
}

func percentage(cp, sp float64) {
	var profit, loss, perc float64
	if sp > cp {
		profit = (sp - cp)
		perc = (profit / cp * 100)
		fmt.Println("Profit : ", profit)
		fmt.Println("Percentage : ", perc, "%")
	} else if cp > sp {
		loss = (cp - sp)
		perc = (loss / cp * 100)
		fmt.Println("loss : ", loss)
		fmt.Println("Percentage : ", perc, "%")
	} else {
		fmt.Println("No Profit No Loss")
	}
}
