package main

import (
	"context"
	"fmt"
	"time"
)

type Message struct {
	content string
}

func main(){
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	go func(){
		for i := 0; i< 10; i++ {
			message := Message {content: "hello from go routine1",}
			ch1 <- message
			time.Sleep(time.Second * 2)
		}

		close(ch1)
	}()

	go func(){
		for i := 0; i< 10; i++ {
			message := Message {content: "Hi from routine2",}
			ch2 <- message
			time.Sleep(time.Second * 1)
		}

		close(ch2)
	}()

	
	// timeout := time.After(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)

	for {
		select {
		case v, ok := <- ch1:
			if !ok {
				ch1 = nil
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
			}
			fmt.Println(v)
		case v := <- ctx.Done():
			fmt.Println("timeout", v)
			cancel()
		}

	}

	


}