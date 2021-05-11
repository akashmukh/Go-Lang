package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	//pow := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range pow {
		pow[i] = i * i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
		//fmt.Printf("len=%d cap=%d = %v\n", len(pow), cap(pow), value)
	}
}
