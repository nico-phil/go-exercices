package main

import (
	"fmt"
	"sync"
)

func increment(c *int, wg *sync.WaitGroup, mu *sync.Mutex){
	defer wg.Done()
	mu.Lock()
	*c = *c + 1
	mu.Unlock()
}

func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0

	for i:=1; i < 100001; i++ {
		wg.Add(1)
		go increment(&count, &wg, &mu)
	}

	wg.Wait()

	fmt.Println(count)
}