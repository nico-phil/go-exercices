package main

import "testing"


func BenchmarkPassByValue(b *testing.B){

	for i:=0; i < b.N; i++ {
		ls := LargeStruct{}
		processValue(ls)
	}
}

func BenchmarkPassByPointer(b *testing.B){

	for i:=0; i < b.N; i++ {
		ls := LargeStruct{}
		processByPointer(&ls)
	}
}
