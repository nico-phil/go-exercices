package main

import "fmt"



type Person struct {
	Name string
	Age int
	Address *Address
}

type Address struct {
	City string
	State string
}

func (p *Person)Shallowcopy() *Person{
	return &Person{
		Name: p.Name,
		Age: p.Age,
		Address: p.Address,
	}
}


func (p *Person) DeepCopy() *Person {
	newAddress := Address {
		City: p.Address.City,
		State: p.Address.State,
	}
	return &Person{
		Name: p.Name,
		Age: p.Age,
		Address: &newAddress,
	}
}

func main(){
	original := &Person{
		Name: "John",
		Age:  30,
		Address: &Address{
			City:  "New York",
			State: "NY",
		},
	}

	copyShallow := original.Shallowcopy()
	deepCopy := original.DeepCopy()

	fmt.Println("copyShallow1", copyShallow.Name, copyShallow.Address.City)
	fmt.Println("DeepCopy1", deepCopy.Name, deepCopy.Address.City)

	original.Address.City = "THE BEST CICY"

	fmt.Println("copyShallow2", copyShallow.Name, copyShallow.Address.City)
	fmt.Println("DeepCopy2", deepCopy.Name, deepCopy.Address.City)
}