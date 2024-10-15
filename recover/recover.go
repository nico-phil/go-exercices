package main

import (
	"fmt"
	"time"
)

func doPanic(msg string){
	panic(msg)
}

func div60(i int){

	defer func(){
		v := recover()
		if v != nil {
			fmt.Println("recovered operation:", v)
		}
	}()

	fmt.Println(60 / i)
}


func safeGoroutine(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("recovered from panics in go routine", r)
		}
	}()

	panic("goroutine panicking...")
}

func main(){
	// values := []int{2,4,0,5,10}
	// for _, v := range values {
	// 	div60(v)
	// }

	go safeGoroutine()
	time.Sleep(time.Second)

	fmt.Println("program continue running...")
}