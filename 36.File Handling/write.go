package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	text := "All the best"
	mydata := []byte(text)
	err := ioutil.WriteFile("write.txt", mydata, 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Write Successfully")
	}
}
