package main

import "fmt"



func main(){
	c := 0
	// pointertoC := &c

	increment := func(){
		for i := 0; i < 10000; i++ {
			c = c + 1
		}
	}
	
	increment()

	fmt.Println("c", c)
}