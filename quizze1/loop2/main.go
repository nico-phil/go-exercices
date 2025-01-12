package main

func f() {
	var a = [2]int{5, 7} 
	for i, v := range a { 
		if i == 0 {
			a[1] = 9 // the original array does not change since the array own it value. 
		} else {
			print(v)
		}
	}
}

func g() {
	var a = [2]int{5, 7}
	for i, v := range a[:] {
		if i == 0 {
			a[1] = 9 // {5, 9} // the original array does change because, the slice does not won its value. it point to an underline array
		} else {
			print(v)
		}
	}
}

func main() {
	f()
	g()
}