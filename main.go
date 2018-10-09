package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// This is one way to use a struct and add values to it, but it is not recommended because the values are assigned to the struct
	// 		properties based upon the order, so if you ever have to swap the order of the values for any reason e.g. lastname firstname
	// 		then you are in a world of hurt with this method
	//
	// alex := person{"Alex", "Anderson"}

	// This is the second way of using a struct and adding values to the structure. It is more recommend than the 1st method.
	// 		The only difference is that you're adding the value directly to the specific struc property you'd like the value assigned to
	//
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson"}

	var alex person

	fmt.Println(alex)
}
