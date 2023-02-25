package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
/*
	append func working

	i := []int{10, 20, 30, 40, 50 } // len = 5 , cap =5
	append(i,60) // not enough cap so allocation is going to happen

//  sufficiently large underlying array.
	underlying array -> [10 20 30 40 50,60,{},{}] len =6 cap = 8

append(i,70,90,300) // allocation would happen as we don't have enough cap to fit three values
	underlying array -> [10 20 30 40 50,60,70,80,90, , , , ] len =9 cap = 13

	If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values.
    Otherwise, append re-uses the underlying array.
*/
func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	inspectSlice("x", x)
	x = append(x, 80)
	inspectSlice("x", x)
	x = append(x, 90)
	inspectSlice("x", x)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("%p\n", slice)
	fmt.Println(slice)
	fmt.Println()

}
