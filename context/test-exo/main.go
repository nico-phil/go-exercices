package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx, cancel  := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	r := countTo(ctx, 10)
	for i := range r {
		fmt.Println(i)
	}
}

func countTo(ctx context.Context, max int) chan int {
	ch := make(chan int)
	go func(){
		defer close(ch)
		for i:=0; i < max; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}