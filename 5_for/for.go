package main

import "fmt"

// for -> only construct in go for looping

func main(){
	// while loop implementation through for 

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// infinite loop 
	// for {
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i < 3; i++ {
	// 	if i== 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range
	for i := range 3{
		fmt.Println(i) // prints 0 1 2
	}
}