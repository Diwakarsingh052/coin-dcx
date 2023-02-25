package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

// In client-streaming RPC, the client sends multiple messages/request to the server
// instead of a single request.
// The server sends back a single response to the client.
// limit of 100 requests per second

type wrappedStream struct {
	//ServerStream defines the server-side behavior of a streaming RPC.
	//it helps in sending and receiving resp at server side
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func ServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("====== [Server Stream Interceptor] ",
		info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return nil
}

// RecvMsg method is defined to implement ServerStream Interface
func (w *wrappedStream) RecvMsg(message interface{}) error {

	log.Printf("====== [Server Stream Interceptor Wrapper] "+
		"Receive a message (Type: %T) at %v",
		message, time.Now().Local())
	//intercepting the request sent to the server by client
	err := w.ServerStream.RecvMsg(message)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(message, "", " ")
	log.Println("****server interceptor message recv*****", string(b))
	fmt.Println()
	return nil

}

func (w *wrappedStream) SendMsg(message interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] "+
		"Send a message (Type: %T) at %v",
		message, time.Now().Local())
	//sending resp back to the client
	log.Println(message)
	return w.ServerStream.SendMsg(message)

}
