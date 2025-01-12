package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3} 

	x := a[:1] //  x = []int{0} l=1, cap=4

	y := a[2:] //  y = []int{2, 3} l=1, cap=4
	

	x = append(x, y...) // x =[]int{0, 2,3}, a = {0,2,3,3}, y={3,3}

	x = append(x, y...)  // x = []int{0,2,3,3,3}, a={0,2,3,3}
	
	fmt.Println(a, x) //  a ={0,2,3,3} x={0,2,3,3,3}
}