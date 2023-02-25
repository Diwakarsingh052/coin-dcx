package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40}
	a[2] = 100
	//a[5] = 800 // update only

	fmt.Println(a)

}
