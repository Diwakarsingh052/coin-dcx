package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	ctx := context.Background() // empty context where we will put timeouts,values in the future
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // clean resources taken up by context
	go func() {

		time.Sleep(1 * time.Millisecond)
		c1 <- "one" // send

	}()
	PrintData(ctx, c1)
	//db := sql.DB{}
	//http.NewRequestWithContext()
	//db.QueryRowContext()

}

func PrintData(ctx context.Context, c chan string) { // ctx should be the first param in func
	select {
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished // or time is up
		fmt.Println("timeout")
		fmt.Println(ctx.Err())

	case s := <-c:
		fmt.Println(s)

		//case <-time.After(1 * time.Second):
		//	//exec the case when the time specified is over
		//	fmt.Println()
	}
}
