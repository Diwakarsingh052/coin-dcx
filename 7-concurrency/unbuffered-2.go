package main

import "fmt"

// recv is always a blocking call in unbuffered and buffered
func main() {

	c := make(chan int)
	go func() {
		fmt.Println(<-c) // this line blocks
	}()

	c <- 100
	fmt.Println("end of the main")
	// this program needs waitgroup to make sure we wait for recv goroutine to finish
}
