package database

import "fmt"

type Config struct {
	db string
}

func (c Config) AccessToDb() {

	fmt.Println("accessing the data from the ", c.db)
}
