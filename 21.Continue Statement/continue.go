package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i += 1 {
		if i < 5 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
