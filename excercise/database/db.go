package database

import "fmt"

type Config struct {
	db string
}

//create a func to create instance of the struct

func (c Config) AccessToDb() {

	fmt.Println("accessing the data from the ", c.db)
}
