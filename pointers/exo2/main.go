package main

import "fmt"

func main(){
	nums := [5]int{1,2,3,4,5}
	// nums2 := [5]int{1,2,3,4,4}
	// r := compare(&nums, &nums2)

	fmt.Println("Before",nums)
	
	modify(&nums)

	fmt.Println("",nums)
	
}

func modify(nums *[5]int) {
	for i :=0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}
}


func reverseArray(nums *[5]int){
	start := 0
	end := len(nums) - 1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func compare(arr1 *[5]int, arr2 *[5]int) bool{
	if len(arr1) != len(arr2) {
		return false
	}

	start := 0
	for start < len(arr1){
		if arr1[start] != arr2[start]{
			return false
		}

		start++
	}

	return true
}


