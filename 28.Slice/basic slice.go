package main

import (
	"fmt"
)

func main() {
	values := [6]int{2, 4, 6, 8, 10, 12}
	var a []int = values[1:5]
	fmt.Println(a)

	fmt.Println(values[:4])
	fmt.Println(values[3:])
}
