package main

import (
	"fmt"
	"time"
)


func main(){

	ch := make(chan string)

	go GenerateNumbers(ch)

	for {
		select {
		case v, ok := <- ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <- time.After(time.Second):
			fmt.Println("timeout, no value received")
		}
	}

}

func GenerateNumbers(ch chan<- string){
	for i:=0; i < 10; i++ {
		ch <- fmt.Sprintf("message %d", i)
		time.Sleep(time.Second )
	}

	close(ch)
}