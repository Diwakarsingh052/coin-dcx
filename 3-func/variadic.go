package main

import "fmt"

func main() {
	// variadic param can accept any number of values
	show("hello", false, 10, 20, 30, 40, 70, 80) // variadic param are optional
	//b := []int{10, 20, 30, 40}
	//abc(b)

}

func show(s string, b bool, i ...int) { // variadic param should be the last in the func signature
	fmt.Printf("%T\n", i)
	fmt.Println(i)

}
func abc(i []int) {

}
