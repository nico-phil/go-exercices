package main

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