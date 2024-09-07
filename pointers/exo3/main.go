package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5,6}
	fmt.Println("before", arr)
	modidy(&arr)

	fmt.Println("after", arr)
}

func modidy(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 3
	}
	*arr = append(*arr, 160)
}



func resize(arr *[]int, newSize int){
	if newSize > len(*arr) {
		*arr = append(*arr, make([]int, newSize - len(*arr))...)
	}else {
		*arr = (*arr)[:newSize]
	}
}