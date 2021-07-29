package services

import (
	"io"
	"log"
	"os"

	"github.com/architagr/golang-microservice-tutorial/authentication/services"
	rpc_auth "github.com/architagr/golang-microservice-tutorial/rpc/rpc_auth"
)


type ValidateRpcServer struct {
	rpc_auth.UnimplementedValidateTokenServiceServer
}


func (ValidateRpcServer) Validate(stream rpc_auth.ValidateTokenService_ValidateServer) error{
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger := log.New(os.Stdout, "validateRpc", 1)

		service := services.NewLogin(logger, flags)
		
		valid, claims := service.VerifyToken(in, "")

		if !valid {
			if err := stream.Send(&rpc_auth.ValidateTokenResponse{
				IsValid: valid,
				ComapnyId: "",
				Username:"",
				Roles:nil,
			}); err != nil {
				return err
			}
		} else{
		if err := stream.Send(&rpc_auth.ValidateTokenResponse{
			IsValid: valid,
			ComapnyId: claims.ComapnyId,
			Username:claims.Username,
			Roles:claims.Roles,
		}); err != nil {
			return err
		}
	}
	}
	return nil
}