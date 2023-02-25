package main

import "fmt"

func main() {

	var i []int // nil

	//make([]int,len, cap)
	i = make([]int, 0, 1000) //pre allocating the underlying array

	i = append(i, 15)
	inspectSlice("name", i)

	//i = make([]int, 10) // len = cap
	//for x:= range i {
	//	i[x] = somevalue
	//}
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
