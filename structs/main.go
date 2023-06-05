package main

import "fmt"

type simplePerson struct {
	firstName string
	lastName  string
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo   //'contact' var could also be emitted so that type name 'contactInfo' becomes the var name as well
}

type contactInfo struct {
	email string
	zipcode int
}

func main() {
	
	//Declarations

	//1st way of declaration of structs
	john := simplePerson{"john", "thomas"}  //This type of declaration completely depends on the order of the variables in struct, hence not that safe
	fmt.Println(john.lastName)

	//2nd way of declaration
	jake := simplePerson{firstName: "jake", lastName: "thomas"}
	fmt.Println(jake.lastName)

	//3rd wat of declartion
	var alex simplePerson
	alex.firstName = "alex"
	alex.lastName = "hales"
	fmt.Printf("%+v", alex)  //prints all the properties along with the corresponding values



	//Embedding structs

	andre := person{
		firstName: "andre",
		lastName: "lotham",
		contact: contactInfo{
			email: "andre18@gmail.com",
			zipcode: 1234,
		}, //Comma is intentional
	}
	fmt.Printf("%+v", andre)



	//Calling the receiver functions with struct
	andre = andre.updateFirstName("andreas")

	//using pointers to update the name
	andrePointer := &andre
	andrePointer.updateFirstNameByPointer("andreas")

	//other way of using the pointer
	andre.updateFirstNameByPointer("andreas")
	andre.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(firstName string) person{		//pass by value hence had to return the struct as well
	p.firstName = firstName
	return p
}

func (pointerToPerson *person) updateFirstNameByPointer(firstName string) {
	(*pointerToPerson).firstName = firstName
}