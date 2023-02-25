package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

//desc contains a description of the stream,
//desc - streaming RPC service's method specification.
//cc is the ClientConn on which the RPC was invoked.
//streamer is the handler to create a ClientStream
//and it is the responsibility of the interceptor to call it.
//ClientStream defines the client-side behavior of a streaming RPC.

type wrappedStream struct {
	//ClientStream defines the client-side behavior of a streaming RPC.
	//it helps in sending and receiving resp at client side
	grpc.ClientStream
}

// creating the stream interceptor
func clientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("======= [Client Interceptor] ", method)
	log.Println("[Client Interceptor]", desc)
	//s represents the gRPC client stream returned by the streamer function.
	//streamer is the handler to create a ClientStream
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return newWrappedStream(s), nil
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] "+
		"Receive a message (Type: %T) at %v",
		m, time.Now().Local())

	//receiving message from the server
	err := w.ClientStream.RecvMsg(m)
	if err != nil {
		return err
	}

	log.Println("client interceptor recv", m)
	fmt.Println()
	return nil
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] "+
		"Send a message (Type: %T) at %v",
		m, time.Now().Local())
	b, _ := json.MarshalIndent(m, "", " ")
	log.Println("client interceptor send", string(b))
	fmt.Println()
	//sending msg to the server
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{ClientStream: s}
}
