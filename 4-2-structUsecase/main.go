package main

import (
	"coin-dcx/4-2-structUsecase/user"
	"log"
	"os"
)

func main() {

	l := log.New(os.Stdin, "sales:", log.LstdFlags)
	c, err := user.NewService("localhost", " :8080", "postgres", l)
	if err != nil {
		l.Fatalln(err)
	}

	c.AddToDb("john", 33)
	c.Update("abc", 28)
}
