package main

import (
	"fmt"
)

func main() {
	var speed float64
	var time float64
	speed = 100
	time = 5.5
	fmt.Println("Speed : ", speed, "time :", time)
	distance := speed * time
	fmt.Println("The distance is : ", distance)
	//fmt.Println("The distance is : ", speed*time)
}
