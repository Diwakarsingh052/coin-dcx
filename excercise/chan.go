package main

import (
	"fmt"
	"sync"
)

//A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
//channel -> unbuffered chan, buffered chan

// A send on a buffered channel can proceed if there is room in the buffer

var wg = &sync.WaitGroup{}

func main() {
	//channel helps to send signals and data from one go routine boundary to another go routine
	c := make(chan int) // unbuffered chan
	wg.Add(4)

	go addNum(10, 20, c)
	go mult(10, 10, c)
	go sub(100, 90, c)
	go calcAll(c)

	// please create a new channel when we have a new series of task or different task than other go routines
	//go workOnJson()

	wg.Wait()

}

func addNum(a, b int, c chan int) {
	defer wg.Done()

	sum := a + b
	// in case of an unbuffered chan , receiver must be ready otherwise send will block
	// do send operation signal on the channel  // signaling with data
}

func sub(a, b int, c chan int) {
	defer wg.Done()
	sum := a - b
	// do send operation signal on the channel  // signaling with data
}

func mult(a, b int, c chan int) {
	defer wg.Done()
	sum := a * b
	// do send operation signal on the channel  // signaling with data
}

func calcAll(c chan int) {
	defer wg.Done()
	//recv from the channel

	fmt.Println(x, y, z)
}
