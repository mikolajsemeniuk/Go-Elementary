// rountine stands for asynchronous
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	// create brand new channel of type string
	// to prevent main routine finishing execution before child routines
	c := make(chan string)

	for _, l := range links {
		// make our function asynchronous using `go` => `go routine` == `one single process`
		// `go` == create a new thread go routine
		// go execute main routine when progam starts and start new child routine when we use `go` keyword
		go checkLink(l, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println("might be down")
		// send data to channel
		c <- "might be down"
	}

	fmt.Println(l, "up and running")
	// send data to channel
	c <- "up and running"
}
