package main

import (
	"fmt"
)

func main() {
	var arry = [2][2]int{{1, 2}, {3, 4}}
	//i						0       1
	//j					   0,1     0,1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("array[%d][%d] = %d\n", i, j, arry[i][j])
		}
	}
}
