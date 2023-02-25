package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()                            // empty context where we will put timeouts,values in the future
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // ctx with a timeout , cancel can cancel req anytime
	defer cancel()                                         // Whenever cancel is called it's going to stop the request at the same time
	PrintData(ctx, "hello this is a test func")            //if  Cancel() is called in defer then will clean up the resources

}

func PrintData(ctx context.Context, input string) { // ctx should be the first param in func
	select {
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished
		fmt.Println("timeout")
		fmt.Println(ctx.Err()) // to check err inside the context

	case <-time.After(4 * time.Second): // run the code after specified amt of time
		fmt.Println(input)

	}
}
