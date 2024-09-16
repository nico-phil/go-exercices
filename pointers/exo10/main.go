package main



type LargeStruct struct {
	filed1 [1000]int
	Field2  [1000]int
	Field3  [1000]int
	Field4  [1000]int
	Field5  [1000]int
	Field6  [1000]int
	Field7  [1000]int
	Field8  [1000]int
	Field9  [1000]int
	Field10 [1000]int
}

// simulate operation
func processValue(ls LargeStruct){
	ls.filed1[0] = 100
}

func processByPointer(ls *LargeStruct){
	ls.filed1[0] = 100
}


