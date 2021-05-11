package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("read.txt")
	if err != nil {
		fmt.Println(err)
	}
	mydata := string(data)
	mydata = strings.Replace(mydata, "bubu", "akash", -1)
	err = ioutil.WriteFile("read.txt", []byte(mydata), 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully replaced")
	}
}
