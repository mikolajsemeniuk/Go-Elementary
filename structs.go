package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type mail struct {
	email string
	// also we coudl use only person to declare property of type person and name of person
	person person
}

func (m mail) print() {
	fmt.Printf("\n%+v", m)
}

// won't work
// func (this mail) setMail(s string) {
// 	this.email = s
// }

// will work
func (this *mail) setMail(s string) {
	(*this).email = s
}

// value types, 					| reference type
// (need pointer to modify value)   | (no need to pointer to modify)
// ---------------------------------|------------------------------
// int 								| slices
// float							| maps
// string							| channels
// bool								| pointers
// structs							| functions

func main() {
	alex := person{"hi", "there"}
	mike := person{firstName: "", lastName: ""}

	mike.firstName = "heh"
	mike.lastName = "there"

	fmt.Printf("alex: %v\nmike: %v\n", alex, mike)

	var john person
	// show properties names
	fmt.Printf("john: %+v", john)

	var email = mail{email: "email", person: person{firstName: "mike", lastName: "lastName"}}
	email.setMail("mike@mock")
	email.print()

	emailPointer := &email
	emailPointer.setMail("mike@pointer")
	email.print()
}
