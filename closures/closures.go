package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}
func main() {
	increment := counter()

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

// when the function is executed it goes to the call stack 
// when the execution is completed, it is removed from the call stack 
// which means all the variables in the function are deleted from the memory 
// but when we call the function counter and then the inner function is returned 
// the count variable should be deleted but it won't be bcz 
// it is referencing to the variable in the outer function
// if in a function a variable is being used which is of the outer scope (here count)
// then it will be always available in the inner function even after the execution of the inner function