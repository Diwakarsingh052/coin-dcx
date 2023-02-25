package main

import "fmt"

//https://go.dev/doc/faq#methods_on_values_or_pointers

type user struct {
	name string
}

// any changes that we would do in update method , then those changes would be reflected back to the caller variable
func (u *user) update(name string) {
	u.name = name // *(dereference) is not required // automatically done
}

func (u *user) show() {
	fmt.Println() // *(dereference) is not required // automatically done
}

func main() {
	u := user{name: "John"}
	u.update("abc") // u would not be copied but its address would get copied in the pointer receiver
	fmt.Println(u)
}
