package main

import (
	"fmt"
	"sync"
)

//var wg = sync.WaitGroup{}

func main() {
	wg := &sync.WaitGroup{}
	//wg.Add(10) // counter to add number of goroutines that we are starting or spinning up
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg) // 10 goroutines // each func call creates a goroutine
	}

	//wg.Add(1) // it is a deadlock as we are running just 10 goroutines not 11
	fmt.Println("task 1 of main going on")
	fmt.Println("task 2 of main going on")

	wg.Add(1)
	go calc(10, 20, wg)
	wg.Wait() // it will wait until counter resets to zero
	fmt.Println("end of main")

}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the counter by one // defer make sure counter is always decremented even in case of panic
	//time.Sleep(2 * time.Second)
	//panic("work") // panic reveals go routines id
	fmt.Println("I am doing some work", i)
}

func calc(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := add(a, b)
	fmt.Println("calc", s)
}

func add(a, b int) int {
	return a + b
}
