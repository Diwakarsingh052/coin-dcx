package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[1:4] // index:len

	//b = append(b, 888) //It is going to change the base slice because they are  sharing the same memory because we have enough cap left
	inspectSlice("b", b)

	// we are adding to add 4 elems here, but we don't have enough cap to fit all the values,
	//allocation will happen here
	//x slice would not change as b slice have its own backing array

	b = append(b, 888, 999, 777, 666)
	inspectSlice("b", b)
	inspectSlice("x", x)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()

}
