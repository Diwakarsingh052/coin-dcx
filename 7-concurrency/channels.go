package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
//A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
//channel -> unbuffered chan, buffered chan

// A send on a buffered channel can proceed if there is room in the buffer

var wg = &sync.WaitGroup{}

func main() {
	//channel helps to send signals and data from one go routine boundary to another go routine
	c := make(chan int) // unbuffered chan
	//wg.Add(1)
	go addNum(10, 20, c)

	//time.Sleep(3 * time.Second)
	//recv from the channel //it blocks the go routine where it is being used
	x := <-c // recv will block until addNum sends the data
	fmt.Println(x)
	//wg.Wait()

}

func addNum(a, b int, c chan int) {
	//defer wg.Done()
	fmt.Println("addNum started")
	sum := a + b
	c <- sum // send operation signal on the channel  // signaling with data
	fmt.Println("add num finished")
}
