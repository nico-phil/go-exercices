package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(3)

	go goRoutineOne(&wg)
	
	go goRoutineTwo(&wg)

	go goRoutinethree(&wg)

	wg.Wait()

	fmt.Println("done")
}

func goRoutineOne(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hello from goroutine one")
}

func goRoutineTwo(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hello from goroutine two")
}

func goRoutinethree(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hello from goroutine two")
}