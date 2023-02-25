package main

import "io"

//	type ReadWriter interface {
//		Reader //embedding
//		Writer
//	}

type user struct {
	name string
}

func (u user) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (u user) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	var rw io.ReadWriter
	var r io.Reader
	var w io.Writer
	_, _, w = rw, r, w
	u := user{name: "abc"}

	rw = u
	r = u
	w = u

	anything(u)

}

func anything(r io.ReadWriter) {

}
