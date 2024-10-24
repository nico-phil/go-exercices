package main

import "fmt"


type Comp struct {
	a int
	b string
}

type CompI interface {
	int
}

func main(){
	
	 m := map[int]int{}
	 modifyMap(m)

	 fmt.Println(m)


}

func modifyMap(m map[int]int) {
	m[100] = 100
 }