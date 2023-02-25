package main

import "fmt"

func main() {

	//var i interface{} = 100
	var i any

	i = "hello"

	i = struct{}{}

	s, ok := i.(string) // type assertion
	if !ok {
		fmt.Println("not of string type")
		return
	}
	fmt.Println(s)
}
