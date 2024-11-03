package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}

	x := a[:1] //  x = []int{0} l=1, cap=1=4
	fmt.Println("x", x, len(x), cap(x))

	y := a[2:] //  y = []int{2, 3}
	fmt.Println("y", y, len(y), cap(y))

	x = append(x, y...) // x =[]int{0, 2,3} 

	x = append(x, y...)  // x = []int{0,2,3,2,3}
	
	fmt.Println(a, x) //  a = []int{0,1,2,3}, x = []int{0,2,3,2,3}
}