package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// How to create an instace of struct
	// #Method-1 (Relying on the position): alex := person {"Alex", "Anderson"}
	// #Method-2(Prefer) (Relying on the key)
	// alex := person {firstName: "Alex", lastName: "Anderson"}
	
	// Declare empty struct alex person and updating it's value
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contact: contactInfo{
			email: "alex@gmail.com",
			zipCode: 70000,
		},
	}
	
	// alex.updateFirstName("Jimmy")

	// *
	// Get the address of variable and pass into method
	// So the method can modified directly on the value of that variable

	// Method-1 for (*), we get manually the address of variable and pass to the method
	// pointerAlex := &alex
	// pointerAlex.updateFirstName("Jimmy")

	// Method-2 for (*), we let "go" support extract the address the variable for us and pass into the method
	alex.updateFirstName("Jimmy")
	alex.print()
}

func (pointerToPerson *person) updateFirstName(updatedFirstName string) {
	(*pointerToPerson).firstName = updatedFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
