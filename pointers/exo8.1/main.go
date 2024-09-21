package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	id int
	name string
}


func generateUser(i int, userChanel chan<- User, wg *sync.WaitGroup){
	defer wg.Done()
	user := User{id: i, name: fmt.Sprintf("user%d", i)}

	userChanel <- user
}


func main(){

	userChanel := make(chan User)
	var wg sync.WaitGroup

	start := time.Now()
	
	for i:=0; i < 100000; i++ {
		wg.Add(1)
		go generateUser(i, userChanel, &wg)
	}

	
	go func(){
		wg.Wait()
		close(userChanel)
	}()

	for u := range userChanel {
		fmt.Printf("name:%s, id:%d \n", u.name, u.id)
	}

	elapedTime := time.Since(start)

	fmt.Println("total time", elapedTime)
}

// 452 ms pointer
// 480 ms copy

/**

Pointers (Sharing Data):

Memory Efficiency: Passing pointers is memory efficient because all goroutines modify the same data. No additional memory is allocated for each goroutine.
Synchronization Overhead: Since all goroutines share the same data, we need to use synchronization mechanisms like sync.Mutex to avoid race conditions. This introduces some overhead.
Performance: When synchronization is done carefully, passing pointers tends to be faster for large data structures because it avoids copying large amounts of data.

Copying Data:
Memory Usage: Passing copies of data creates separate instances for each goroutine, leading to higher memory usage if the data is large.
No Synchronization Needed: Since each goroutine has its own copy of the data, no synchronization (mutexes) is required, which can simplify the code.
Performance: For small data structures, copying might be faster because it avoids the overhead of synchronization. However, for larger data structures, the cost of copying can outweigh the benefits.

**/

