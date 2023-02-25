package main

import "fmt"

func main() {

	// func without name is an anonmous func
	func(a, b int) {
		// do your work
		fmt.Println(a + b)
	}(10, 20) // calling of the function
}
