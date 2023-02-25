package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	u user // simple field

	role []string
}

func (a admin) show() {
	fmt.Printf("%+v\n", a)
}

func main() {

	a := admin{
		u: user{
			name:  "abc",
			email: "abc@email.com",
		},
		role: []string{"admin", "dev"},
	}

	fmt.Println(a.u.name)
	a.u.updateEmail("bob@email.com")
	a.show()

	//accept(u) // it will not work

}

func accept(a admin) {

}
