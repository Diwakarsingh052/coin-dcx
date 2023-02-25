package main

import "fmt"

func main() {
	a := 10
	var p *int // this var can store address of var of int type // default is nil
	p = &a
	// any changes made by p would be reflected back to the var a because p stores the memory address of a var and it is directly manipulating it

	*p = 20 // * is a dereference operator to access the memory stored in the pointer
	fmt.Println(a)
	*p++
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(&p) // address of p where it is storing address of the var
}
