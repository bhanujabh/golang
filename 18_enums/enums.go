package main

import "fmt"

// enumerated types

// // custom types
// type MyType string
// we use custom types to make enums

type orderStatus int 

const(
	Received orderStatus = iota // Received's value has become 0
	Confirmed // its  value is now 1 bcz of iota 
	Prepared  // 2
	Delivered // 3
)

// another way 

// type orderStatus string
// const(
// 	Received orderStatus = "received" 
// 	Confirmed = "confirmed"
// 	Prepared  = "prepared"
// 	Delivered = "delivered"
// )

func changeOrderStatus(status orderStatus){
	fmt.Println("changing order status to", status) 
	// changing order status to 3
}

func main(){
	changeOrderStatus(Delivered)
}
