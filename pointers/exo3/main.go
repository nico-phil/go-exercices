package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5,6}
	fmt.Println("before", arr)
	modidy(&arr)

	fmt.Println("after", arr)
}

func modidy(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 3
	}
	*arr = append(*arr, 160)
}



func resize(arr *[]int, newSize int){
	if newSize > len(*arr) {
		*arr = append(*arr, make([]int, newSize - len(*arr))...)
	}else {
		*arr = (*arr)[:newSize]
	}
}



type CustomDynamicArray struct {
	data     *[1]int // Pointer to the underlying array
	length   int     // Current number of elements
	capacity int     // Capacity of the array
}

// NewDynamicArray initializes a new CustomDynamicArray with a given capacity
func NewDynamicArray(initialCapacity int) *CustomDynamicArray {
	// Create a new array with the specified initial capacity
	return &CustomDynamicArray{
		data:     new([1]int), // Start with a small array, we'll grow it later
		length:   0,
		capacity: initialCapacity,
	}
}

// Add adds an element to the dynamic array, reallocating memory if needed
func (da *CustomDynamicArray) Add(value int) {
	// If length exceeds capacity, reallocate memory with a new array
	if da.length == da.capacity {
		// Double the capacity
		newCapacity := da.capacity * 2
		// Allocate a new array with the new capacity
		newData := new([1]int)

		// Copy old data to the new array
		for i := 0; i < da.length; i++ {
			newData[i] = da.data[i]
		}

		// Update the pointer to the new array
		da.data = newData
		// Update capacity
		da.capacity = newCapacity
	}

	// Add the new element
	da.data[da.length] = value
	// Increase the length
	da.length++
}

// Get retrieves the value at a specific index in the dynamic array
func (da *CustomDynamicArray) Get(index int) (int, bool) {
	if index >= 0 && index < da.length {
		return da.data[index], true
	}
	return 0, false // Return false if index is out of bounds
}

// Len returns the current length of the dynamic array
func (da *CustomDynamicArray) Len() int {
	return da.length
}

// Cap returns the current capacity of the dynamic array
func (da *CustomDynamicArray) Cap() int {
	return da.capacity
}

