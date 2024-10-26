package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	const numWorkers = 3
	const numTasks =  15

	tasks := make(chan int, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for i :=0; i < numTasks; i++ {
		tasks <- i
	}

	close(tasks)

	wg.Wait()
	fmt.Println("all go routine are done")


}

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("process id: %d task %d \n",id,task)
		time.Sleep(time.Second)
	}
}

