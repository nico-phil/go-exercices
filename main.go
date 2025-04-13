package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fibC(boolChan <-chan bool, resultChan chan<- int) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	var canWrite = make(chan struct{}, 1)

	// Goroutine to control the rate
	go func() {
		for range ticker.C {
			select {
			case canWrite <- struct{}{}:
			default:
				// Don't block if the channel already has a signal
			}
		}
	}()

	for range boolChan {
		select {
		case <-canWrite:
			resultChan <- rand.Int()
		default:
			// Drop the write if it's too soon (rate-limited)
		}
	}
}

func main() {
	boolChan := make(chan bool)
	resultChan := make(chan int)

	go fibC(boolChan, resultChan)

	// Simulate sending signals to boolChan
	go func() {
		for {
			boolChan <- true
			time.Sleep(3 * time.Millisecond)
		}
	}()

	// Receive results
	for val := range resultChan {
		fmt.Println("Received:", val)
	}
}
