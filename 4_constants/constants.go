package main

import "fmt"

// can declare const and var outside function too
const age = 30
// var name string = "golang"
// shorthand is not allowed outside function
// name := "golang"

func main(){
	// const name string = "golang"

	// fmt.Println(age )

	const(
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}