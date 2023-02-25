package main

import (
	"fmt"
	"strconv"
)

func main() {
	const job = 3
	ch := make(chan string, job) // buffered channel // size = 10

	for work := 1; work <= job; work++ {
		go func(w int) { // spinning a go routine which doesn't means we are going to run this at this moment
			ch <- "work" + strconv.Itoa(w)
		}(work) // pass in the local copy of the var to the go routine to avoid unexpected results
	}

	for i := 1; i <= job; i++ {
		fmt.Println(<-ch)
	}
}
