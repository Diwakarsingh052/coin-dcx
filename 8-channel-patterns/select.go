package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}

	//wgWorker keep track of if the go routine work is finished or not and we wil close the channel when work is done
	var wgWorker = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan bool, 1)

	wgWorker.Add(3)
	go func(a string) {
		defer wgWorker.Done()

		//v, err := strconv.Atoi(a)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//c1 <- v
		time.Sleep(3 * time.Second)

		// doing some work
		//err can happen
		//return

		c1 <- "one" // send

	}("101")

	go func() {
		defer wgWorker.Done()

		time.Sleep(1 * time.Second)
		c2 <- "two"

	}()

	go func() {
		defer wgWorker.Done()

		c3 <- "three"

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//waiting for go routines to finish
		wgWorker.Wait()
		fmt.Println("closing")
		//closing the channel
		close(done)
	}()

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			case x := <-c1: // recv over the channel
				fmt.Println("send the result in the further pipeline", x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)
			case <-done: // this case will exec when channel is closed
				fmt.Println("it is closed")
				return
			}
		}
	}()

	fmt.Println("end of the main")
	wg.Wait()
}
