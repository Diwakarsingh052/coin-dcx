package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

//method is the RPC name. req and reply are the corresponding request and response messages.
//cc is the ClientConn on which the RPC was invoked.
//invoker is the handler to complete the RPC,
//and it is the responsibility of the interceptor to call it

func UnaryClientInterceptor(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Preprocessor phase
	log.Println("Method intercept: " + method)
	log.Println("interceptor", req)

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	if err != nil {
		fmt.Println()
		return err
	}
	// Postprocessor phase
	log.Println("unary interceptor ", reply)
	fmt.Println()
	return nil
}
