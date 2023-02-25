package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not able to find in root directory")

func main() {
	_, err := openFile("any.txt")

	//errors.Is looks inside the chain and check if custom err happened or not
	if errors.Is(err, ErrFileNotFound) {
		fmt.Println("our custom error is found in the chain, lets create the file")
		//create the file
		//retry the operation
		//if still failed then stop this func
		log.Println(err)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

}

func openFile(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)

	if err != nil {
		// %w verb wraps two errors together // go 1.13 introduced wrapping and unwrapping
		return nil, fmt.Errorf("%v %w", err, ErrFileNotFound)
	}
	return f, nil

}
