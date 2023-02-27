package main

import (
	"coin-dcx/excercise/data"
	"coin-dcx/excercise/data/stores/userdb"
)

func main() {
	// call Create method of postgres using Storer interface
	u := data.User{
		Name:  "abc",
		Email: "abc@email.com",
	}
	p := userdb.NewPostgres(nil)
	s := data.NewStore(&p)
	_ = s.Create(nil, u)

}
