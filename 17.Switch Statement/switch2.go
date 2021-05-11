package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Hour(), ":", t.Minute(), ":", t.Second())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() == 12:
		fmt.Println("Good Noon")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	case t.Hour() > 17:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}
}
