package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)


func main(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	files := []string{"file1.txt", "file.txt", "file3.txt"}
	data, err := processFiles(ctx, files)
	if err != nil {
		log.Fatal("err in main:", err)
	}

	fmt.Println("data:", data)

}

func processFiles(ctx context.Context, filenames []string) (map[string]int, error){
	countResult := map[string]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan error)
	for _, v := range filenames {
		wg.Add(1)
		go func(v string){
			defer wg.Done()
			count, err := process(ctx, v)
			if err != nil {
				fmt.Println("err in go func", err)
				ch <- err
			}

			mu.Lock()
			countResult[v] = count
			mu.Unlock()
		}(v)
		 
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println("error in pass to chanel", v)
	}

	return countResult, nil
}


func process(ctx context.Context, filename string) (int, error) {
	// time.Sleep(time.Second * 1)
	
	// var mu sync.Mutex
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	b := make([]byte, 10)
	for {
		n, err := file.Read(b)
		// fmt.Println("n", n)
		if err != nil {
			if err == io.EOF {
				break
				
			}
		}
		// mu.Lock()
		count += n
		// mu.Unlock()
	}
	return count, nil
}