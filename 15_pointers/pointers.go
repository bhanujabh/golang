package main

import "fmt"

// by value 
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num) // 5
// }

// by reference 
// * -> address of the variable 
func changeNum(num *int){
	// num = 5 // can't change like bcz we received a pointer 
	// first we will have to dereference it 
	*num = 5
	fmt.Println("In changeNum", num) //In changeNum 0x14000106020
	fmt.Println("In changeNum", *num) // 5
}

func main() {
	num := 1

	changeNum(&num)

	// fmt.Println("Memory address", &num) // 0x140000100a0

	// fmt.Println("After changeNum in main", num) // 1 when changed by value 

	fmt.Println("After changeNum in main", num) // 5
}