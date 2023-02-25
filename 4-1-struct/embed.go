package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type time int
type admin struct {
	user // embedding // anonymous field because we are not giving any name to it // a
	// anonymous field name would be assigned according to the type name
	//time  // any custom type can be embedded in the struct

	role []string
}

func (a admin) show() {
	fmt.Printf("%+v\n", a)
}

func main() {

	a := admin{
		user: user{
			name:  "abc",
			email: "abc@email.com",
		},
		role: []string{"admin", "dev"},
	}

	fmt.Println(a.name)
	a.updateEmail("bob@email.com")
	a.show()

	//accept(u) // it will not work

}

func accept(a admin) {

}
