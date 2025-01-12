package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		defer panic(1)
		recover()
	}()
	defer recover()
	panic(2)
}

// d1, d2(d3,d4) d5, 