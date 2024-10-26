package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func main(){
	urls := []string{"url1", "url2", "url3", "url4"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()

	result, err := fetchUrls(ctx, urls)
	if err != nil {
		fmt.Println("err in main:", err)
		return
	}

	fmt.Println("result:", result)

}

func fetchUrls(ctx context.Context,  urls []string) (map[string]int, error) {
	results := map[string]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i, u := range urls {
		wg.Add(1)
		go func(u string, i int){
			defer wg.Done()
			s, err := fetchUrl(ctx, u)
			if err != nil{
				fmt.Println("err in fetch", err)
			}

			// race contition
			mu.Lock()
			results[s] = i + 1
			mu.Unlock()

		}(u, i)
	}

	wg.Wait()

	return results, nil
}

func fetchUrl(ctx context.Context, url string) (string, error) {

	for {
		select {
		case <- time.After(time.Second * 1):
			return url, nil
		case <- ctx.Done():
			return "Done", ctx.Err()
		}
	}
	
}