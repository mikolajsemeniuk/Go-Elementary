package main

import "fmt"

type deck []string

// assign method to type `class` in other languages
func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}
