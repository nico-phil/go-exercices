package main

import "fmt"


func update(v **int){
	**v = 50
}

func main(){
	v := 10

	pv := &v

	fmt.Println("before", v)
	update(&pv)

	fmt.Println("after", v)
}