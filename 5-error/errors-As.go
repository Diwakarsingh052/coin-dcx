package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	var nErr *strconv.NumError
	_, err := strconv.Atoi("efg")
	if errors.As(err, &nErr) {
		fmt.Println(nErr.Err)
	}

}
