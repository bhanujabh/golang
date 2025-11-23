package main

import "fmt"

func main() {
	// age := 15

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermissions = true

	// if(role == "admin" && hasPermissions==true){
	// 	fmt.Println("Yes")
	// }

	// we can declare variable inside if construct 
	if age:= 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go doesn't have ternary operator, you will have to use normal if-else 
}