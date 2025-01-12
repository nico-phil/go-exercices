package main


func main() {
	m := make(map[string]int, 3) //{}
	x := len(m) // 0
	m["Go"] = m["Go"] // {"GO": 0}
	y := len(m) // 1
	println(x, y) // 0, 1 
}
