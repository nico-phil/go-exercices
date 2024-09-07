package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=1; i < 10; i++ {
			ch <- i
		}

		close(ch)
	}()

	go func(){
		defer wg.Done()
		for v := range ch {
			fmt.Println("v=", v)
		}
	}()


	wg.Wait()

}

