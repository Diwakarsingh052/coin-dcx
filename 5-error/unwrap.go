package main

import (
	"errors"
	"fmt"
)

var ErrFeesNotSubmitted = errors.New("fees not submitted")
var ErrAdmissionCancelled = errors.New("admission cancelled")
var ErrDocumentNotSubmitted = errors.New("document not submitted")

// u -> admission -> fees - > documents

func main() {
	err := admission()
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

}

func admission() error {
	err := fees()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrAdmissionCancelled)
	}
	return nil
}

func fees() error {
	err := documents()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrFeesNotSubmitted)
	}
	return nil
}

func documents() error {
	//assuming the student didn't submit the documents
	return fmt.Errorf("%w", ErrDocumentNotSubmitted)
}
