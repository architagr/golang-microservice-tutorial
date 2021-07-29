package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	rpc_auth "./../../rpc/rpc_auth"
	"github.com/architagr/golang-microservice-tutorial/authentication/models"
	"github.com/architagr/golang-microservice-tutorial/authentication/services"

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

	url, _ := flags.GetApplicationUrl()

	lis, err := net.Listen("tcp", *url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	rpc_auth.RegisterLoginServiceServer(grpcServer, &services.LoginRpcServer{})
	err1 := grpcServer.Serve(lis)
	if err1 == nil {
		fmt.Println("grpc server running on", *url)
	}
}

