package main

import "fmt"

type user struct {
	name string
}

func main() {
	//u := data{name: "abc"}
	CheckType(true)
}

func CheckType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("it is an int", v)
	case string:
		fmt.Println("it is a string", v)
	case user:
		fmt.Println("it is data type", v)
	default:
		fmt.Println("nothing matches")

	}
}

//x := "1"
//
//	switch x {
//	case "1":
//		fmt.Println("x is 1")
//		fallthrough
//		// break is by default
//	case "2":
//		fmt.Println("x is 2")
//		fallthrough
//
//	case "3":
//		fmt.Println("x is 2")
//	default:
//		fmt.Println("x is neither 1 nor 2")
//	}
