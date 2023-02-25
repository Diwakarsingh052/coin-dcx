package main

import "fmt"

func main() {
	a := []int{900, 670, 610, 950, 870, 200}
	//b := make([]int, len(a)) // creating a new memory with len according to a slice
	b := make([]int, 3, 100) // only three elem max can be copied here as that is the available len only
	//copy(b, a)
	copy(b, a[1:4])
	b[0] = 8989
	fmt.Println(a)
	fmt.Println(b)

}
