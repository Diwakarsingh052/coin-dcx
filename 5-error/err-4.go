package main

import (
	"errors"
	"fmt"
)

//Notice how strconv package changes is input field and func name according to statements exec
//strconv.Atoi: parsing "abc": invalid syntax
//strconv.Atoi: parsing "xyz": invalid syntax
//strconv.ParseInt: parsing "efg": invalid syntax

//_, err := strconv.Atoi("abc")
//	fmt.Println(err)
//	_, err = strconv.Atoi("xyz")
//	fmt.Println(err)
//	_, err = strconv.ParseInt("efg", 10, 64)
//	fmt.Println(err)

//type error interface {
//	Error() string
//}

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

// QueryError is dedicated to work with error handling, it should not be used to work with normal data
type QueryError struct {
	Func  string
	Input string
	Err   error
}

// this method is compulsory to format your err msg //
func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + ": " + q.Err.Error()
}

func main() {
	err := SearchSomething("data")

	var q *QueryError       // nil
	if errors.As(err, &q) { // reference imp // it checks whether struct is inside the error chain or not
		fmt.Println("true", q.Func)
		return
	}
	fmt.Println("not")

	//fmt.Println(err)
	//err = SearchName("John")
	//fmt.Println(err)
}

func SearchSomething(s string) error {

	//do searching and if that is not found then return the error below
	// QueryError struct can be returned as an error value because error method is implemented over it.
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   ErrNotFound,
	}

}

func SearchName(name string) error {
	//do searching and if that is not found then return the error below
	return &QueryError{
		Func:  "SearchName",
		Input: name,
		Err:   ErrMismatch,
	}
}
