package main

import (
	"context"
	"fmt"
	"time"
)

/**
	3. Context in Concurrent Goroutines
		3.1: Context Propagation in Goroutines
		Start a goroutine that runs a long task.
		Pass a context.Context to the goroutine and make it responsive to the context's cancellation signal.
		Cancel the context after a certain time and ensure the goroutine exits gracefully.
**/


func longRunningTask(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			fmt.Println("context cancelled:", ctx.Err())
			return
		case <- time.After(time.Second * 1):
			fmt.Println("running...")

		}
	}
}

func main(){

	cxt, cancel := context.WithCancel(context.Background())

	go longRunningTask(cxt)

	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second)
	fmt.Println("main func exit")

}