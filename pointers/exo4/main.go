package main

import (
	"fmt"
)

// CustomDynamicArray represents a dynamic array using pointers to arrays
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

func main() {
	// Initialize a custom dynamic array with an initial capacity of 2
	dArray := NewDynamicArray(2)

	// Add elements to the dynamic array
	for i := 1; i <= 4; i++ {
		dArray.Add(i)
		fmt.Printf("Added %d: Length = %d, Capacity = %d\n", i, dArray.Len(), dArray.Cap())
	}

	// Retrieve and print elements from the dynamic array
	for i := 0; i < dArray.Len(); i++ {
		value, _ := dArray.Get(i)
		fmt.Printf("Element at index %d: %d\n", i, value)
	}
}
