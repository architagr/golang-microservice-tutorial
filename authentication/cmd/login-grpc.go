package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/architagr/golang-microservice-tutorial/authentication/models"
	"github.com/architagr/golang-microservice-tutorial/authentication/services"
	"github.com/architagr/golang-microservice-tutorial/authentication/token"
	rpc_auth "github.com/architagr/golang-microservice-tutorial/rpc/rpc_auth"

	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)
var flags *models.Flags

func main() {
	flag.Parse()
	flags = models.NewFlags(*ip, *port)
	token.Init()
	url, _ := flags.GetApplicationUrl()

	lis, err := net.Listen("tcp", *url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	rpc_auth.RegisterLoginServiceServer(grpcServer, &services.LoginRpcServer{})
	rpc_auth.RegisterValidateTokenServiceServer(grpcServer,&services.ValidateRpcServer{})
	fmt.Println("starting grpc server on", *url)
	err1 := grpcServer.Serve(lis)
	if err1 == nil {
		fmt.Println("grpc server running on", *url)
	}else{
		fmt.Println("grpc server running error on", err1)
	}
}

