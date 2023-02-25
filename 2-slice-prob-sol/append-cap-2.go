package main

import "fmt"

func main() {
	i := []int{10, 20, 30, 40}
	b := i[:2] // index:len

	i = append(i, 50)
	b = append(b, 999)
	inspectSlice("i", i)
	inspectSlice("b", b)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("%p\n", slice)
	fmt.Println(slice)
	fmt.Println()

}
