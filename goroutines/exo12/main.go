package main

import (
	"fmt"
	"sync"
)

func main(){

	nums := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)
	go gerenate(nums, &wg)
	go square(nums, squares, &wg)
	go printSquared(squares, &wg)

	
	wg.Wait()

	fmt.Println("DONE")

}

func gerenate(nums chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i< 10; i++ {
		nums <- i
	}

	close(nums)
}

func square(nums <-chan int, squareNumbers chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for num := range nums {
		squareNumbers <- num * num
	}

	close(squareNumbers)
}

func printSquared(squareNumbers <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for s := range squareNumbers {
		fmt.Println("square", s)
	}
}

