package main

import (
	"context"
	"fmt"
	"time"
)

/**
	1.3: Passing Context through Function Calls
		Implement a chain of function calls where each function receives a context.Context.
		Simulate an operation in each function that checks for cancellation or timeout and returns accordingly.
		Ensure that if the context is canceled or times out at any point, the rest of the function chain halts execution.

**/

func One(ctx context.Context){
	time.Sleep(time.Millisecond * 500)
	if ctx.Err() != nil {
		fmt.Println("One: cancelled")
		return
	}
	
	fmt.Println("funcOne: running")

	two(ctx)
}

func two(ctx context.Context){
	time.Sleep(time.Second)
	if ctx.Err() != nil {
		fmt.Println("two: cancelled")
		return
	}
	
	fmt.Println("funcTwo: running")

	three(ctx)
}

func three(ctx context.Context){
	time.Sleep(time.Second * 2)
	if ctx.Err() != nil {
		fmt.Println("three: cancelled")
		return
	}
	
	fmt.Println("funThree: running")
}



func main(){
	ctx, cancel :=  context.WithCancel(context.Background())

	go One(ctx)

	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second * 2)
}