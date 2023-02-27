package main

import "log"

type user struct {
	name  string
	email string
}

func main() {
	u := user{
		name:  "john",
		email: "john@email.com",
	}

	l := log.New(u, "custom log: ", log.LstdFlags) // make this line work
}
