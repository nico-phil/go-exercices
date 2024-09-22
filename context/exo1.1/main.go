package main

/**
	1. Basics of Context
		1.1: Simple Timeout Context
		Write a function that simulates a long-running process (e.g., time.Sleep for 3 seconds).
		Use context.WithTimeout to enforce a timeout of 2 seconds.
		Print whether the operation completed or was canceled due to the timeout.

**/

import (
	"context"
	"fmt"
	"time"
)

func longruningProcess(ctx context.Context) error {
	select{
	case <- time.After(time.Second * 2):
		fmt.Println("Process complete successfully")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}

func longRuningProcees2(){
	fmt.Println("start process....")
	time.Sleep(time.Second * 5)
	fmt.Println("end process...")
}

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	done := make(chan bool)

	go func(){
		longRuningProcees2()
		done <- true
	}()

	select {
	case <- done:
		fmt.Println("process completed within the timeout")
		return
	case <- ctx.Done():
		fmt.Println("timeout exceded")
		return
	}
}

// when you have a long running function, you can put it in goroutine
// make a chanel, to tell when the long running function complete 
// you can use context to controle the timeout of the long running function
// use select to listen to the ctx.Done chanel and the other chanel to know if the function complete
// within the timeout