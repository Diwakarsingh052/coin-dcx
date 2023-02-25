package main

import "fmt"

func main() {
	y := 10 // x809
	z := &y //v = x809 , p= x840
	fmt.Println(&z)
	update(z)
	fmt.Println(y)
}

//go is pass by copy, which means value would be copied to the function call,
//in this case an address is copied to p // p have its own address

func update(p *int) { //v =x809 p = x850  // y address would be copied here
	fmt.Println(&p)
	*p = 100

}
