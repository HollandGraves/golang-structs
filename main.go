package main

import "fmt"

// STRUCTS AND DATA TYPES BEGIN HERE

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// MAIN BEGINS HERE

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contactInfo: contactInfo{
			email:   "cdxx00042@gmail.com",
			zipCode: 80501,
		},
	}

	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Bob", "Bobby")
	jimPointer.print()
}

// DEFINED FUNCTIONS BEGIN HERE

func (pointerToPerson *person) updateName(newfirstname string, newlastname string) {
	(*pointerToPerson).firstName = newfirstname
	(*pointerToPerson).lastName = newlastname
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
