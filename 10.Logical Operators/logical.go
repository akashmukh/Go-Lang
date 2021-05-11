package main

import (
	"fmt"
)

func main() {
	a, b, c, d := 10, 15, 20, false
	fmt.Println((a > b) && (b > c))
	fmt.Println((a > b) || (b > c))
	fmt.Println(!d)
}
