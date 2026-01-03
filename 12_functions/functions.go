package main

import "fmt"

// if same type of parameters then
// func add(a, b int) all parameters before it become int
func add(a int, b int) int {
	return a + b
}

// function can return multiple values 
// we also return error here 
func getLanguages()(string, string, bool){
	return "golang", "javascript", true
}

func main(){
	fmt.Println(getLanguages()) // golang javascript true in one line
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3) // golang javascript true in one line
	// suppose we dont wanna use one of the values
	val1, val2, _ := getLanguages()
	fmt.Println(val1, val2) // golang javascript
	result := add(3, 5) // 8 
	fmt.Println(result)
}