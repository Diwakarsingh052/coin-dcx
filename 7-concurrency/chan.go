package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go doWork(c) // go routine
	recvWork(c)  // this line block the main func as well
	fmt.Println("hey")
}

func doWork(c chan int) {

	c <- 10 // send

}
func recvWork(c chan int) {

	x := <-c // recv // blocking call until we don't recv data or signal
	fmt.Println(x)
}
