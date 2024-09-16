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

