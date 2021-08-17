package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	// other go types:
	//	* bool
	//	* int
	//	* float64
	//	* string
	//	* time
	//  * nil == null
	var card string = "there" // this statement is equal to card := "there"
	card = "there"
	hi := newString()

	fmt.Printf("hi: %v\n", hi)
	fmt.Printf("hi %v\n", card)

	one, two := returnTwoVariables()
	fmt.Println(one, two)

	createSliceOfStringsAndShowThem()

	// initialize new type
	decks := deck{"hi", "there"}
	// call method assign to this type
	decks.print()

	// convert string to binary array
	fmt.Println("bytes: ", []byte("Hi there"))
}

func newString() string {
	return "hehs"
}

func returnTwoVariables() (string, string) {
	return "one", "two"
}

// GO has array => fixed size
// 	  and slice => array which could shirk or grow (List in other languages)
//    both array and slice all elements has to be the same type
func createSliceOfStringsAndShowThem() {
	slices := []string{"hey", "there"}
	// add new element to array
	slices = append(slices, "new element")
	fmt.Println(slices)
	// or iterate to show them
	// use `_` instead of `i` if you don't want to use it
	for i, s := range slices {
		fmt.Println(i, s)
	}
	// if you want to slice use syntax below
	fmt.Println(slices[0:])
}
