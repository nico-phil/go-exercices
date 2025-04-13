package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)


func main(){
	urls := []string{
		"https://example.com",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://golang.org",
		"https://api.github.com",
		"https://go.dev",
		"http://fakeeeee.com",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
		"https://jsonplaceholder.typicode.com/posts/7",
		"https://jsonplaceholder.typicode.com/posts/8",
		"https://jsonplaceholder.typicode.com/posts/9",
		"https://jsonplaceholder.typicode.com/posts/0",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()

	// result, err := fetchUrls3(ctx, urls)
	// if err != nil {
	// 	fmt.Println("err in main:", err)
	// 	return
	// }
	// fmt.Println("result:", result)

	// for r := range fetchUrls3(ctx, urls){
	// 	if r.Error != nil {
	// 		fmt.Println("err in main: result", r)
	// 		return
	// 	}

	// 	fmt.Println("result in main:", r.Status)
	
	// }

	start := time.Now()
	urlchan := generateUrl(ctx, urls)

	r1 := fetchUrlFromChan(ctx, urlchan)
	r2 := fetchUrlFromChan(ctx, urlchan)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		for r := range r1 {
			fmt.Println(r)
		}
	}()

	go func(){
		wg.Done()
		for r := range r2 {
			fmt.Println(r)
		}
	}()

	wg.Wait()

	fmt.Println("time:", time.Since(start))

}

func fetchUrls(ctx context.Context,  urls []string) (map[string]int, error) {
	results := map[string]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error)
	for i, u := range urls {
		wg.Add(1)
		go func(u string, i int){
			defer wg.Done()
			s, err := fetchUrl(ctx, u)
			if err != nil{
				errChan <- err
			}

			// race contition
			mu.Lock()
			results[s] = i + 1
			mu.Unlock()

		}(u, i)
	}

	go func(){
		wg.Wait()
		close(errChan)
	}()

	err := <- errChan 
	if err != nil  {
		return map[string]int{}, err
	}


	return results, nil
}

func fetchUrl(ctx context.Context, url string) (string, error) {

	for {
		select {
		case <- time.After(time.Second * 1):
			return url, fmt.Errorf("there is an error, in url:%s", url)
		case <- ctx.Done():
			return "Done", ctx.Err()
		}
	}
	
}

func fetchUrls2(ctx context.Context, urls []string, )(map[string]string, error){
	results := map[string]string{}
	errChan := make(chan error)
	var mu sync.Mutex	
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			res, err := fetchUrl2(ctx, url)
			if err != nil {
				fmt.Println(err)
				errChan <- err
			}

			mu.Lock()
			defer mu.Unlock()
			results[url] = res

		}(u)
	}

	go func(){
		wg.Wait()
		close(errChan)
	}()

	err := <-errChan
	if err != nil {
		return map[string]string{}, err
	}

	return results, nil

}


func fetchUrl2(ctx context.Context, url string)(string, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}

	// return response.StatusCode, nil

	defer response.Body.Close()
	 
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	return string(data), err

}

type Result struct {
	Status int
	Error error
}

func generateUrl(ctx context.Context, urls []string) <-chan string {
	out := make(chan string)
	go func(){
		defer close(out)
		for _, v := range urls {
			out <- v
		}
	}()

	return out
}

func fetchUrlFromChan(ctx context.Context, urlchan <-chan string)(<-chan Result ){
	out := make(chan Result)
	go func(){
		defer close(out)
		for u := range urlchan {
			select {
			case <-ctx.Done():
			default:
				res, err := fetchUrl3(ctx, u)
				out <- Result{Status: res , Error:err}
			
			}
			
		}

	}()

	return out
}

func fetchUrl3(ctx context.Context, url string)(int, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil

	// defer response.Body.Close()
	 
	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return "", nil
	// }

	// return string(data), err

}