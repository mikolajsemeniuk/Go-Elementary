package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("key: ", color, ", value: ", hex)
	}
}

func main() {
	// key and values have to be the same
	colors := map[string]string{
		"red":   "blue",
		"green": "yellow",
	}
	// won't work for map[string]string
	// color["white"] = "black"
	// color["red"] = "red"
	fmt.Printf("%+v\n", colors)

	colors = make(map[string]string)
	// this will work for make(map[string]string)
	colors["black"] = "black"
	colors["white"] = "white"
	colors["yellow"] = "yellow"
	delete(colors, "black")
	fmt.Printf("%+v\n", colors)

	printMap(colors)
}
