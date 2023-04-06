package main

import (
	"coin-dcx/excercise/database"
	"log"
)

func main() {
	c, err := database.NewConfig("localhost", " :8080 ", "postgres")
	if err != nil {
		log.Println(err)
		return
	}

	c.AccessToDb()

}
