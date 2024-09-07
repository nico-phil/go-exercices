package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go GenerateNumbers(ch, &wg)
	go ReceiveNumbers(ch, &wg)

	wg.Wait()

	fmt.Println("Done")


}

func GenerateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <  10; i++ {
		ch <- i
	}

	close(ch)
}

func ReceiveNumbers(ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for num := range ch {
		fmt.Println("num", num)
	}
}