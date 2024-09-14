package main

import (
	"fmt"
)


func receive(prt *int){
	*prt = 150
}



func main(){
	prt := new(int)
	fmt.Println("before", *prt)
	receive(prt)
	fmt.Println("after", *prt)

	intSlice := make([]int, 5)
	fmt.Println("before (make)",intSlice)

	for i:= 0; i < 5 ; i++ {
		intSlice[i] = i * 3
	}

	fmt.Println("after (make)", intSlice)
}