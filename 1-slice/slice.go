package main

import "fmt"

func main() {
	// slice points to an array (backing , underlying) in the memory
	var i []int // nil slice
	//i[10] = 100 // this will cause panic as length is not available to store the value

	fmt.Println(i)
	fmt.Printf("%#v\n", i)

	//nil = not pointing to backing array in memory
	//empty = not storing data
	if i == nil {
		fmt.Println("it is a nil slice")
	}

	b := []int{}
	if b == nil {
		fmt.Println("b is a nil slice")
	} else {
		fmt.Println("b is just empty")
	}

}
