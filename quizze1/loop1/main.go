package main

func main() {
	x := []int{7, 8, 9}
	y := [3]*int{}
	for i, v := range x { 
		defer func() {
			print(v) // 7, 8, 9
		}()
		y[i] = &i // y={0,1,2}
	}
	print(*y[0], *y[1], *y[2], " ") //0,1,2 "" 9,8,7
}