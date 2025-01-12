package main

// When a function is pushed into the deferred call queue, the function value and all the argument are
// evaluated. The evaluated values will be used when the call is executed later

func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()
	
	var f func()
	defer f() //
	f = func() {
		r += 2
	}

	return 1
}

func main() {
	println(bar())
}