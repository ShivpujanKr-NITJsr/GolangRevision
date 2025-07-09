package main

import (
	"fmt"
	"slices"
)

// import "fmt"

func main() {
	// slices are dynamic arrays, can grow and shrink in size
	// slices are reference types, they point to the underlying array
	// slices are more flexible than arrays, can be resized, appended, and sliced

	// var slice []int // nil slice ,unintialized
	// fmt.Println(slice==nil)
	// println(len(slice)) // length of the slice
	// slice := []int{1, 2, 3} // slice literal
	// slice := make([]int, 5) // create a slice with length 5 and capacity 5
	// fmt.Println(cap(slice))
	// slice := make([]int, 2, 5) // create a slice with length 2 and capacity 5

	// // append to a slice
	// slice = append(slice, 4)
	// fmt.Println(slice) // prints [0 0 4]

	// // capacity doubles

	// nums := []int {}
	// nums=append(nums,1,2,3,4,5) // append multiple elements
	// nums[2]=10 // modify an element in the slice
	// fmt.Println(nums) // prints [1 2 3 4 5]

	// copy function

	// var nums = make([]int, 0, 5)

	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums) // copy the contents of nums to nums2

	// fmt.Println(nums, nums2)

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[1:3]) // prints [2 3] ,slice from index 1 to 3 (exclusive)
	// fmt.Println(nums[:3])  // prints [1 2 3] ,slice from the beginning to index 3 (exclusive)
	// fmt.Println(nums[2:])  // prints [3 4 5] ,slice from index 2 to the end
	// fmt.Println(nums[:])   // prints [1 2 3 4 5]

	// slices
	var nums1 = []int{1, 2}
	var num2 = []int{1, 2}
	fmt.Println(slices.Equal(nums1, num2))


	var nums =[][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(nums) // prints true,
}
