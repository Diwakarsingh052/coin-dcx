package main

import (
	"errors"
	"fmt"
	"log"
)

type Config struct {
	db string
}

func AddToDb(db string, data string) {

}

func main() {
	c, err := NewConfig("localhost", " :8080", "postgres")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("accessing the db", c.db)
	AddToDb(c.db, "some data")
	//log.New
	//os.OpenFile()
	//errors.New()

}

var ErrConn = errors.New("please provide a valid connection")

// NewConfig func is initializing the Config struct
func NewConfig(host, port, conn string) (*Config, error) {
	if host == "" || port == "" || conn == "" {
		return nil, ErrConn
		// without pointer // Config{},ErrCon

	}

	db := host + port + conn

	c := Config{db: db}

	return &c, nil

}
