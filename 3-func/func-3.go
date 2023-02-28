package main

import (
	"fmt"
)

func main() {
	add := func(a, b int) {
		fmt.Println(a + b)
	}

	sub := func(a, b int) {
		fmt.Println(a - b)
	}
	//sumString := func(a, b string) { // func with a different signature
	//	fmt.Println(a + b)
	//}

	//add(10, 20) // calling the func directly using the var created
	calc(add) //add var can be passed to the calc func as the type of func matches
	calc(sub)
	//calc(sumString) // type should be same
	//http.HandleFunc()
}

func calc(sum func(x, y int)) {

	sum(10, 40)

}
