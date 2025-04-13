package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)


func main(){
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}

	stautsResponse := checkStatus(done, urls...)

	for r := range stautsResponse {
		if r.Error != nil {
			fmt.Printf("Error:%v\n", r.Error)
			continue
		}

		fmt.Printf("response:%v\n", r.Response.Status)
		
	}

}

type Result struct {
	Error error
	Response *http.Response
}


func checkStatus(done <-chan interface{}, urls... string) <-chan Result{
	responses := make(chan Result)

	go func(){
		defer close(responses)
		for _, s := range urls {
			resp, err := http.Get(s)
			result := Result{Error: err, Response: resp}
	
			select {
			case <-done:
				return
			case responses <- result :
			}
		}
	
	}()


	return responses
}

func randomNum(done chan interface{}) <-chan int {
	random := make(chan int)
	go func(){
		defer fmt.Println("random finished")
		// defer close(random)
		for {
			select {
			case random <- rand.Int():
			case <- done:
				return
			}
			
		}
	}()
	

	return random
}

func doWork(done <-chan any, strings <-chan string) chan any{
	terminated := make(chan any)
	go func(){
		fmt.Println("woWork launched")
		defer fmt.Println("do work exited")
		defer close(terminated)
		for {
			select {
			case s := <- strings:
				fmt.Println(s)
			// case <- done:
			// 	return
			}
		}
	}()
	
	return terminated
}

func chanOwner() <- chan int {
	results := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		// defer close(results)
		defer wg.Done()
		for i:=0; i <= 5; i++{
			results <- i
		}
	}()

	go func(){
		wg.Wait()
		close(results)
	}()

	return results
}

func consume(results <-chan int){
	for v := range results {
		fmt.Println(v)
	}
}

func loopData(handleData chan<- int, data []int){
	defer close(handleData)
	for i :=range data {
		handleData <- data[i]
	}
}


func increment(count *int){
	*count++
}





func removeFromqueue(queue *[]int, delay int, c *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(delay))
	c.L.Lock()
	removedValue := (*queue)[0]
	*queue = (*queue)[1:]
	fmt.Println("remove from queue", removedValue)
	defer c.L.Unlock()

	c.Signal()
}