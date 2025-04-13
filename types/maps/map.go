package main

import (
	"fmt"
	"hash/maphash"
)


type Comp struct {
	a int
	b string
}

type CompI interface {
	int
}

type Square struct {
	x int
	y string
}

func main(){
	
	//  m := map[int]int{}
	//  modifyMap(m)

	//  fmt.Println(m)

	//  n := map[Square]int{}
	//  fmt.Println(n) 

	var h maphash.Hash
	h.WriteString("hello")
	fmt.Printf("%d\n",  h.Sum64())

}

func modifyMap(m map[int]int) {
	m[100] = 148888
 }