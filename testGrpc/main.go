package main

import (
	"context"
	"fmt"
	"io"
	"log"

	rpc_auth "github.com/architagr/golang-microservice-tutorial/rpc/rpc_auth"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		fmt.Println("error in dial", err)
	}
	defer conn.Close()

	loginClient := rpc_auth.NewLoginServiceClient(conn)
	validateClient := rpc_auth.NewValidateTokenServiceClient(conn)

	loginResponse, err := loginClient.LoginSimpleRPC(context.Background(), &rpc_auth.LoginRequest{
		Username:   "steve@yopmail.com",
		Password:   "steve123",
		RememberMe: true,
	})

	if err != nil {
		fmt.Println("Simple RPC login error", err)
	} else {
		fmt.Println("Simple RPC login", loginResponse)
	}

	stream, err := validateClient.Validate(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a claims : %v", err)
			}
			log.Printf("Got claims - %+v -", in)
		}
	}()

	if err := stream.Send(&rpc_auth.ValidateTokenRequest{
		Token: loginResponse.Token,
	}); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	fmt.Println("sending validate request 2nd time")
	if err := stream.Send(&rpc_auth.ValidateTokenRequest{
		Token: loginResponse.Token+"1",
	}); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
	stream.CloseSend()
	<-waitc
}
