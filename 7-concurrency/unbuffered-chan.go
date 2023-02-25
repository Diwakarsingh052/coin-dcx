package main

import "fmt"

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
func main() {

	c := make(chan int)

	// anonymous go routine
	go func(c chan int) {
		c <- 100 //send // block until no receiver
	}(c)

	x := <-c // recv
	fmt.Println("end of main", x)

}
