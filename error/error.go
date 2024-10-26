package main

import "fmt"

func main(){
	data, err := LoginAndGetData("12345nico", "4444", "data.txt")
	if err != nil {
		fmt.Println("error in main:", err)
		return
	}

	fmt.Println("data", string(data))

}