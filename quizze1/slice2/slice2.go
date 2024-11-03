package main

import "fmt"



func main() {
	var x = []string{"A", "B", "C"} // x = A, M, C

	fmt.Println(len(x), cap(x))

	for i, s := range x { //i = 2
		print(i, s, ",") // 0A, 1M, 2C
		x[i+1] = "M" 
		x = append(x, "Z")  // new underline array is created, it does affect the old one [A, Z, C, Z]
		x[i+1] = "Z"
	}
}

//{7,0,5, 6,}