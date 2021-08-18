package main

import "fmt"

// interfaces in go allows us to add type `bot`
// to all types which has function called `getName`
// and return `string` to enable write function with
// same name but different arguments
type bot interface {
	getName() string
}
type englishbot struct{}
type spanishbot struct{}

func (e englishbot) getName() string {
	return "hola"
}

func (s spanishbot) getName() string {
	return "hello"
}

// we cannot assign this function to interface
// like: func (b bot) printName() {}
// like we did earlier: func (e englishbot) printName()
// and we cannot create instance of interface type but
// we could only assign it to existing type
func printName(b bot) string {
	return b.getName()
}

func main() {
	e := englishbot{}
	s := spanishbot{}
	fmt.Println(printName(e), printName(s))
}
