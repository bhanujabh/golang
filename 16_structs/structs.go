package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition / inheritance -> read more 
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer // struct embedding 
}

// constructor 
// we add new word as convention 
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	} 
	// we return pointer to the struct 
	return &myOrder
}

// receiver type -> in the braces after func keyword 
// by convention we use first letter of struct for receiver type 
// connecting the function with a struct
func (o *order) changeStatus(status string) { 
	// we don't have to dereference it bcz go does it automatically 
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// if you don't set a field then default value is zero value 
	// int => 0 float => 0 string => "" bool => false 
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }

	// one way of making an instance
	// var order order = 
	// second way
	// don't need to assign all the fields 
	newOrder2 := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
		// customer: newCustomer -> for inline 
	}

	newOrder2.customer.name = "robin"
	fmt.Println(newOrder2)

	myOrder := order {
		id: "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.changeStatus("confirmed")

	myOrder.createdAt = time.Now()

	fmt.Println("myOrder", myOrder)
	fmt.Println(myOrder.status)
	fmt.Println(myOrder.getAmount())

	myOrder2 := newOrder("2", 30.50, "received")
	fmt.Println(myOrder2)
	// &{2 30.5 received {0 0 <nil>} { }}
	// it returns a pointer but struct dereferences it automatically 
	fmt.Println(myOrder2.amount)

	// inline struct 
	// when we need the struct only a few times and don't need to make multiple instances 
	language := struct {
		name   string
		isGood bool
	}{"golang", true} // values should be assigned in the same order 

	fmt.Println(language) // {golang true}

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)
}