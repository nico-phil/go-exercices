package main

import "fmt"


func main(){
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i + 1
	}

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

}