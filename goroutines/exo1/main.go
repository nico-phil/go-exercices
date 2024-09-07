package main

import (
	"fmt"
	"time"
)

func main(){
	go printNumbers()
	 
	time.Sleep(time.Second * 11)
	fmt.Println("Main func completed")

}

func printNumbers(){
	for i:=0; i < 11; i++ {
		fmt.Println("i=",i)
		time.Sleep(time.Second)
	}
}

