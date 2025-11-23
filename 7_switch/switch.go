package main

import (
	"fmt"
)

func main(){
	// simple switch 
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// 	// no need to write break, it is working automatically in the bg 
	// case 2: 
	//     fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default: // optional 
	//     fmt.Println("other")
	// }

	// multiple condition switch 
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday: 
	//     fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// type switch
	// function is first class citizen in go, 
	// fn can be passed in another fn as parameter, 
	// can return from a fn, 
	// can be stored in a variable
	// in go we have to use double quotes, single quotes don't work for strings  
	whoAmI := func(i interface{}){ // any type of value can be passed in i using interface{}
		switch t  := i.(type){ // i.(type) will return type of i
		//In a type switch, t gets assigned the value stored in the interface i, not the type itself.
		case int: 
		    fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean value")
		default:
			fmt.Println("other", t)
		} 
	} 

	whoAmI(50.9)

}