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
	if err != nil {
		log.Println(err)
		return
	}
	//err := fmt.Errorf("file not found %v", ErrFileNotFound) // fmt.Errorf returns an error msg

	// problem we will face with merging of the errors
	//if err contains ErrFileNotFound { // not possible with %v verb as it will merge the error
	//	//do something
	//}

}

func openFile(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)

	if err != nil {
		// %v verb merges two errors together  // go 1.13 this was replaced with wrapping
		// fmt.Errorf returns an error msg, we are trying to add some extra info to our error value
		return nil, fmt.Errorf("%v %v", err, ErrFileNotFound)
	}
	return f, nil

}
