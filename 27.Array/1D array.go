package main

import (
	"fmt"
)

func main() {
	var a = [6]int{1, 4, 6, 12, 23, 34}
	for i := 0; i < 6; i++ {
		fmt.Print(a[i], " ")
	}
}
