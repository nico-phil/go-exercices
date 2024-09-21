package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processOne(ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("processOne: listen to context signals...")

	
	<- ctx.Done() 

	fmt.Println("processOne: context canceled")
}

func processtwo(ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("processTwo: listen to context signals...")

	<- ctx.Done() 

	fmt.Println("processTwo: context canceled")
}

func process(ctx context.Context, name string, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("%s: cancelled: %s \n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s:  running... \n",name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main(){

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	
	wg.Add(2)
	go process(ctx, "Process 1",&wg)
	go process(ctx, "Processn 2", &wg)

	time.Sleep(time.Second)
	fmt.Println("cancelling context...")
	cancel()

	wg.Wait()

}