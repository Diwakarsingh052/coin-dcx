package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

func main() {
	data, err := FetchRecord(100)
	if err != nil {
		log.Println(err) // log means you have handled your errors, don't do duplicate logging and once an error is handled then don't propagate that error to further chain
		return
	}
	fmt.Println(data)
}

// prefix your err variables with Err word

var ErrRecordNotFound = errors.New("not found data") //Error string should not be capitalized or end with punctuation mark

func FetchRecord(id int) (string, error) {
	email, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	//os.ErrClosed
	//sql.ErrNoRows

	return email, nil
}
