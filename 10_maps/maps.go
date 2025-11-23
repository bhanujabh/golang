package main

import (
	"fmt"
	"maps"
)

func main(){
	// creating map 
	m := make(map[string]string)

	// setting an element 
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element 
	fmt.Println(m["name"], m["area"]) // golang backend 
	fmt.Println(len(m)) // 2
	
	// accessing a key that doesn't exist generates empty value
	fmt.Println(m["phone"]) // empty string 

	m2 := make(map[string]int)
	m2["age"] = 30
	m2["price"] = 50
	fmt.Println(len(m2)) // 2
	fmt.Println(m2["age"]) // 30
	fmt.Println(m2["phone"]) // 0

	delete(m2, "price")
	fmt.Println(m2) // map[age:30] -> price is deleted 

	clear(m2) // to empty the map 
	fmt.Println(m2) // map[]

	// to create map without map function 
	m3 := map[string]int{"price": 50, "phones": 3}
	fmt.Println(m3) // map[phones:3 price:50]

	v, ok := m3["price"]
	fmt.Println(v) // 50 (value)
	// if that key is not present we will get 0 

	if ok {
		fmt.Println("all ok") // all ok
	} else {
		fmt.Println("not ok")
	}

	m4 := map[string]int{"price": 50, "phones": 3}
	m5 := map[string]int{"price": 50, "phones": 3}
	fmt.Println(maps.Equal(m4, m5)) // true

}