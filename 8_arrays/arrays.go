package main

import "fmt"

// arrays - numbered sequence of specific length
func main(){
	// zeroed values when nothing assigned
	// int -> 0, string -> "", bool -> false
	var nums [4]int

	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums) // [1 0 0 0]

	// array length
	// fmt.Println(len(nums))

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals) // [false false true false]

	var name [4]string
	name[2] = "golang"
	fmt.Println(name) // [  golang ]

	// to declare in sinlge line 
	numArray := [3]int{1, 2, 3}
	fmt.Println(numArray)

	// 2D array 
	numsArray := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(numsArray)

	// - fixed size, that is predictable
	// - Memory optimazation
	// - Contant time access
}