package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u1 user) show() { // func(receiver)methodName(Args)returnTypes {}
	fmt.Println(u1.name)
}
func (u1 user) update(name string) { // value from the main func of the user var would be copied here in the receiver
	u1.name = name

}

func main() {
	u := user{
		name: "John",
		age:  30,
	}

	u.show()
	u.update("abc")
	u.show()
}
