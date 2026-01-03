package main

import (
	"fmt"
)

// iterating over data structures
func main(){
	// iterating over slice 
	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	sum := 0

	for index, num := range nums{
		fmt.Println(num, index) 
		// 6 0
		// 7 1
		// 8 2
		sum = sum + num
	}
	fmt.Println(sum) // 21 

	// iterating over map
	m := map[string]string{"fname": "john", "lname": "doe"}

	for key, value := range m{
		fmt.Println(key, value)
		// fname john
		// lname doe
	}

	// to get only keys 
	for key := range m{
		fmt.Println(key)
		// fname 
		// lname 
	}

	// iterating over string 
	// here i is not index but starting byte of rune
	// ascii is used for upto 255 characters which can be filled in 1 byte 
	// but if we have any character with value more than 255 
	// it can utilise upto 2 bytes 
	// suppose we have a char w val 300 then it will take up index 0 and 1
	// and after that the 2nd char will have index 2 
	for i, c := range "abc"{
		fmt.Println(i, c) // 0 97 1 98 2 99 -> unicode code point rune 
		fmt.Println(i, string(c)) // 0 a 1 b 2 c
	}
}