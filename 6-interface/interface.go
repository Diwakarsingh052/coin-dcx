package main

import (
	"fmt"
)

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike
// Bigger the interface weaker the abstraction // Rob Pike

// interfaces are abstract type // it does not store anything of their own
type reader interface {
	read(b []byte) (int, error)
	//hello() // all the methods should be impl over the struct to use the interface
}

type file struct {
	name string
}

//func (f file) hello() {
//	//TODO implement me
//	panic("implement me")
//}

func (f file) read(b []byte) (int, error) {
	fmt.Println("inside file read")
	s := "hello go devs"
	copy(b, s)
	return len(b), nil
}
func (f file) write() {

}

type jsonObject struct {
	data string
}

func (j jsonObject) read(b []byte) (int, error) {
	s := `{name:"abc"}`
	fmt.Println("inside json read")
	copy(b, s)
	return len(s), nil
}

//func fetch(f file) {
//	data := make([]byte, 50)
//	_, _ = f.read(data)
//	fmt.Println(string(data))
//}

// fetch is a polymorphic func
// fetch() will accept any type of value which implements reader interface
func fetch(r reader) {

	fmt.Printf("%T\n", r)
	data := make([]byte, 50)
	_, _ = r.read(data)
	//r.write () // we can only call methods part of the interface signature

	fmt.Println(string(data))
	fmt.Println()
}

func main() {

	f := file{name: "abc.txt"} // concrete data
	j := jsonObject{data: "any json"}

	_, _ = f, j

	var r reader // nil is the default value of the interface
	r = f        //passing in file var as it impls the interface
	r.read(nil)

	//fetch(f) // we can pass if only when file struct impls the interface
	//fetch(j)

}
