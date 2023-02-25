package main

import "fmt"

func main() {
	var a []int
	b := []int{89, 69, 798}
	a = append(a, 10, 20, 30)
	a = append(a, b...) // we are unpacking the b slice // append func does not accept slice as an argument
	fmt.Println(a)

}
