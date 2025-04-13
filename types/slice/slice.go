package main

import "fmt"

func main(){
	queue := []int{5,3,7,9,1}
	 modifySlice(&queue)
	fmt.Println("in main", queue)
}

func modifySlice(queue *[]int){
	*queue = (*queue)[1:]
	fmt.Println("in func",queue)
}