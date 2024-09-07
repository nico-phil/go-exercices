package main

import "fmt"

func main(){
	a := 100
	increment(&a)
	
	fmt.Println(a)
}


func swap(x *int, y *int){
	temp := *x
	*x = *y
	*y = temp
}

func increment(x *int){
	*x= *x + 1
}
