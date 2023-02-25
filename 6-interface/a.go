package main

// interfaces are abstract type // it does not store anything of their own
type worker interface {
	work(b []byte) (int, error)
	hello() // all the methods should be impl over the struct to use the interface
}

type student struct {
	name string
}

func (s student) hello() {
	//TODO implement me
	panic("implement me")
}

func (s student) work(b []byte) (int, error) {
	return 0, nil
}

func doWork(w worker) {

}

func main() {
	var s student
	doWork(s) // yes or no
}
