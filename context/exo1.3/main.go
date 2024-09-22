package main

import (
	"context"
	"fmt"
	"time"
)


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