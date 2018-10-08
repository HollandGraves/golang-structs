package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	// This is one way to use a struct and add values to it, but it is not recommended because the values are assigned to the struct
	// 		properties based upon the order, so if you ever have to swap the order of the values for any reason e.g. lastname firstanem
	// 		then you are in a world of hurt with this method
	//
	// alex := person{"Alex", "Anderson"}

	alex := person{
		firstName: "Alex",
		lastName: "Anderson"
	}
}
