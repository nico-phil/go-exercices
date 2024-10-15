package main

import "fmt"

func main(){
	s := "Nico"

	fmt.Println(s)
	bs := []byte(s)
	bs[0] = uint8(65)
	fmt.Println(string(bs), s)
	


}