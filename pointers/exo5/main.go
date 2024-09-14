package main

import "fmt"


func ModifyMap(m map[string]int){
	for k, v := range m {
		ptr := &v

		*ptr = *ptr * 2

		m[k] = *ptr
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
	
	v := 4
	
	fmt.Println("before", m)


	update(m, "four", &v)

	// ModifyMap(m)

	fmt.Println("after", m)
}

