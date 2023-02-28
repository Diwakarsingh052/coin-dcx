package main

import (
	"fmt"
	"sync"
	"time"
)

// A send on a buffered channel can proceed if there is room in the buffer
// s -> [, 20, 40,]
func main() {
	var wg = &sync.WaitGroup{}

	ch := make(chan int, 2) //make(chan type,bufferSize)

	wg.Add(1)
	go func(c chan int) { // putting exact value in the loop condition is not the best idea
		defer wg.Done()
		for i := 1; i <= 2; i++ {
			//fmt.Println(<-ch) // recv will take the value out of the buffer, and it will empty the space taken up by the value so we can push more new values
			time.Sleep(3 * time.Second)
			fmt.Println(<-ch)
		}

	}(ch)

	//forgotten sender // case where you forgot to recv values
	//when creating a buffered channel make sure at the same time, where the receiver would be
	ch <- 100 // A send on a buffered channel can proceed if there is room in the buffer
	fmt.Println("100 sent")
	ch <- 200
	fmt.Println("200 sent")

	wg.Wait()
}
