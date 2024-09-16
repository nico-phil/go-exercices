package main

import "fmt"

type myFunc interface {
	func(int, int) int
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func getOperation(op string) func(int, int) int{
	if op == "add"{
		return add
	}else if op == "multiply"{
		return multiply
	}
	return nil
}

func main(){

	addFunc := getOperation("add")
	fmt.Println("add", addFunc(4, 6))

}