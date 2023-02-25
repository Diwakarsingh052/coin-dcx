package main

import "fmt"

func main() {
	//defer maintains a stack. [2,1]
	defer fmt.Println("1") // when your function is returning defer statements will exec
	defer fmt.Println("2")
	fmt.Println("3")
	panic("testing panic") // defer guarantee exec // if the statements are registered before the panic or return
	//return // it stops the exec of the current func
	fmt.Println("4")

}
