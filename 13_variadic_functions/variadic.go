package main

import "fmt"

// func sum(nums ...interface{}) int {
// interface is used to take any type of parameters 
func sum(nums ...int) int {
	// in the parameter above we got a slice of nums 
	// using ... we can receive multiple parameters 
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	fmt.Println(1, 2, 3, 4, 5, 86, "hello") // 1 2 3 4 5 86 hello in single line 
	// Println() can receive any number of parameters 
	// hence it is called a variadic function 
	nums := []int{3, 4, 5, 6}
	// result := sum(3, 4, 5, 6)
	result := sum(nums...)// ... spread operator 
	fmt.Println(result) // 18
}