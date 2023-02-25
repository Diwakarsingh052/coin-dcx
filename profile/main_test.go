package main

import "testing"

//go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out -gcflags -m=2
//go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
//go tool pprof p.out
//list algOne
//weblist algOne
//go test -run none means don't run any test but only benchmark // you could use anything that don't exist
//-benchmem show the memory
func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algOne()
	}
}
