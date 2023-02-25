package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func algOne() {
	// l := &Logger{out: out, prefix: prefix, flag: flag}
	//if out == io.Discard {
	//	l.isDiscard = 1
	//}
	//return l

	l := log.New(os.Stdout, "hello", log.LstdFlags)
	s := fmt.Sprintf("hey %v", l)

	_ = s
	str := "diwakar"
	nr := bytes.NewReader([]byte(str))

	_, _ = str, nr

}
