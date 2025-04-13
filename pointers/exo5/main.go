package main

import "fmt"


func ModifyMap(m map[string]int){
	for k, v := range m {
		ptr := &v

		*ptr = *ptr * 2

		m[k] = *ptr
	}
}

func ModifyMap1(m map[string]int){
	for k, v := range m {
		m[k] = v * 5
 	}
}


func update(m map[string]int,  key string, v *int){
		if _, ok := m[key]; !ok {
			m[key] = *v
		}
}


func main(){
	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println("before", m)
	ModifyMap(m)

	fmt.Println("After", m)
}

