package main

import "fmt"

type Person struct {
	name string
	age int
}

func Modify( p *Person){
	p.age = 30
}

func(p *Person) Update(){
	p.age = 150
}

func main(){
	p := Person {
		name: "Nico",
		age: 22,
	}

	p.Update()

	fmt.Println(p)
}