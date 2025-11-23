package main

import (
	"fmt"
	"slices"
)

// slices are used when we don't know the size of the array -> dynamic array
// ost used construct in go
// useful methods
func main(){
	// uninitialized lice is nil 
	var nums []int // size is not mentioned
	fmt.Println(nums) // [] 
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // 0

	var numSlice = make([]int, 2) // 2 is initial size 
	// third argument in make can be capacity 
	// capacity -> max numbers of elements can fit 
	fmt.Println(numSlice) // [0  0]
	fmt.Println(numSlice == nil) // false
	fmt.Println(len(numSlice)) // 2
	fmt.Println(cap(numSlice)) // 2

	var numSlice2 = make([]int, 2, 5)
	numSlice2 = append(numSlice2, 1)
	numSlice2 = append(numSlice2, 2)
	fmt.Println(numSlice2) // [0 0 1 2]
	fmt.Println(cap(numSlice2)) // 5
	numSlice2 = append(numSlice2, 3)
	numSlice2 = append(numSlice2, 4)
	fmt.Println(numSlice2) // [0 0 1 2 3 4]
	fmt.Println(cap(numSlice2)) // 10 -> resizes automatically to double the current capacity 

	var numSlice3 = make([]int, 0, 5)
	numSlice3 = append(numSlice3, 1)
	numSlice3 = append(numSlice3, 2)
	fmt.Println(numSlice3) // [1 2] 

	numSlice4 := []int{1, 2, 3}
	numSlice4 = append(numSlice4, 4)
	fmt.Println(numSlice4) // [1 2 3 4]
	fmt.Println(len(numSlice4)) // 4
	fmt.Println(cap(numSlice4)) // 6

	var numSlice5 = make([]int, 2, 5)
	numSlice5[1] = 3
	fmt.Println(numSlice5) // [0 3]
	fmt.Println(len(numSlice5)) // 2
	fmt.Println(cap(numSlice5))// 5

	var nums2 = make([]int, 0, 5)
	nums2 = append(nums2, 2)

	var nums3 = make([]int, len(nums2))

	// copy function 
	copy(nums3, nums2) // copy(dest, src)
	fmt.Println(nums2, nums3) // [2] [2]

	// slice operator
	var nums4 = []int{1, 2, 3}
	fmt.Println(nums4[0:2]) // [1 2]
	fmt.Println(nums4[:1])// [1]
	fmt.Println(nums4[1:]) // [2 3]  

	// compare if 2 slices are equal 
	var nums5 = []int{1, 2}
	var nums6 = []int{1, 2}
	fmt.Println(slices.Equal(nums5, nums6)) // true
	// various methods in slices  

	var nums7 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums7) // [[1 2 3] [4 5 6]]
}