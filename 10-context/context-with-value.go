package main

import (
	"context"
	"fmt"
	"log"
)

// Users of WithValue should define their own types for keys

type Key string

func main() {
	// context can store timeouts and values,
	//  values are stored as key value pair like a map
	ctx := context.Background()

	const k Key = "anyKey"

	ctx = context.WithValue(ctx, k, "abc") // putting value in the context associated with a key
	FetchValue2(ctx, k)
}

func FetchValue(ctx context.Context, k Key) {
	v := ctx.Value(k) // fetch the value on the basis of the key
	if v == nil {
		log.Println("value is not here")
		return
	}
	if v == "abc" {
		fmt.Println(v)
	}
	fmt.Println(v)

}
func FetchValue2(ctx context.Context, k Key) {
	v := ctx.Value(k)
	s, ok := v.(string) // type assertion // making sure data is of the same type that we want to fetch
	if !ok {
		fmt.Println("value not there or of a different type")
		return
	}
	fmt.Println(s)
}
