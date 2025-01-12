package main



func main() {
	var x = []string{"A", "B", "C"} // {"A", "M", "C"}

	for i, s := range x { 
		print(i, s, ",")  // 0A,1M, 2C
		x[i+1] = "M" // x[0 + 1] = M=> 
		x = append(x, "Z")  
		x[i+1] = "Z"
	}
}





