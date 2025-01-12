package main

import "fmt"

// index: value
var x = []int{2:5, 6, 0: 7}  

func main() {
	fmt.Println(x) // x = {7,0, 5,6}
}