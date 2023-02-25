package main

import (
	"fmt"
	"sync"
	"time"
)

var cab = 2 // shared resource between goroutines
var wg = &sync.WaitGroup{}

func main() {
	var m = &sync.Mutex{}
	name := []string{"a", "b", "c", "d"}
	for _, n := range name {
		wg.Add(1)
		go bookCab(n, m)
	}

	wg.Wait()
}

func bookCab(name string, m *sync.Mutex) {
	defer wg.Done()
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)

	// when a goroutine acquires a lock then another go routine can't access the critical section until the lock is not released
	//any read , write from other goroutines would not be allowed after lock is acquired
	m.Lock()
	//critical section
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}
	m.Unlock() // release the lock and allow others to enter the critical section
	// ends
	fmt.Println()
}
