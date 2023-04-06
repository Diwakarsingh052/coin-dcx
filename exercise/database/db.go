package database

import (
	"errors"
	"fmt"
)

type Config struct {
	db string // unexported
}

var ErrConn = errors.New("please provide a valid connection")

func NewConfig(host, port, conn string) (*Config, error) {
	if host == "" || port == "" || conn == "" {

		// without pointer // Config{},ErrCon

	}
	db := host + port + conn
	c := Config{db: db}

}

func (c Config) AccessToDb() {

	fmt.Println("accessing the data from the ", c.db)
}
