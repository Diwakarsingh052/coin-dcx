package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// UnaryServerInterceptor Server - Unary Interceptor
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Preprocessing logic
	// Gets info about the current RPC call by examining the args passed in

	log.Println(" [Unary Server Interceptor] ", info.FullMethod)
	log.Println("req", req)
	fmt.Println()

	// Invoking the handler to complete the normal execution of a unary RPC.
	response, err := handler(ctx, req)

	fmt.Println()
	// Post-processing logic
	if err != nil {
		log.Println("Error", err)
		return nil, err
	}
	log.Printf(" Response from the unary rpc server : %s", response)

	fmt.Println()
	return response, nil
}
