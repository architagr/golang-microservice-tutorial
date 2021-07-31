package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/architagr/golang-microservice-tutorial/authentication/models"
	rpc_auth "github.com/architagr/golang-microservice-tutorial/rpc/rpc_auth"
)

type LoginRpcServer struct {
	rpc_auth.UnimplementedLoginServiceServer
}

func (LoginRpcServer) LoginSimpleRPC(ctx context.Context, in *rpc_auth.LoginRequest) (*rpc_auth.LoginResponse, error) {
	logger := log.New(os.Stdout, "loginRpc", 1)
	flags, err := models.GetFlags()
	if err != nil {
		return nil, err
	}
	service := NewLogin(logger, flags)
	loginModel := models.LoginRequest{
		UserName:   in.Username,
		Password:   in.Password,
		RememberMe: in.RememberMe,
	}
	token, err := service.GetToken(loginModel, "")

	if err != nil {
		return nil, err
	}
	return &rpc_auth.LoginResponse{
		Token: token,
	}, nil
}

func (LoginRpcServer) LoginServerStreamRPC(in *rpc_auth.LoginRequestList, stream rpc_auth.LoginService_LoginServerStreamRPCServer) error {
	for _, loginCred := range in.Data {
		logger := log.New(os.Stdout, "loginRpc", 1)
		flags, err := models.GetFlags()
		if err != nil {
			return err
		}

		service := NewLogin(logger, flags)
		loginModel := models.LoginRequest{
			UserName:   loginCred.Username,
			Password:   loginCred.Password,
			RememberMe: loginCred.RememberMe,
		}
		token, err := service.GetToken(loginModel, "")

		if err != nil {
			return err
		}
		if err := stream.Send(&rpc_auth.LoginResponse{
			LoginDetails: loginCred,
			Token:        token,
		}); err != nil {
			return err
		}

	}
	return nil
}

func (LoginRpcServer) LoginClientStreamRPC(stream rpc_auth.LoginService_LoginClientStreamRPCServer) error {
	result := &rpc_auth.LoginResponseList{
		Data: make([]*rpc_auth.LoginResponse, 0, 10),
	}
	logger := log.New(os.Stdout, "loginRpc", 1)
	flags, err := models.GetFlags()
	if err != nil {
		return err
	}

	service := NewLogin(logger, flags)
	for {
		loginCred, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(result)
		}
		loginModel := models.LoginRequest{
			UserName:   loginCred.Username,
			Password:   loginCred.Password,
			RememberMe: loginCred.RememberMe,
		}

		token, err1 := service.GetToken(loginModel, "")
		if err1 != nil {
			return err1
		}

		result.Data = append(result.Data, &rpc_auth.LoginResponse{
			LoginDetails: loginCred,
			Token:        token,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (LoginRpcServer) LoginBiDirectionalRPC(stream rpc_auth.LoginService_LoginBiDirectionalRPCServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println("error in reading data", err)
			return err
		}
		logger := log.New(os.Stdout, "loginRpc", 1)
		flags, err := models.GetFlags()
		if err != nil {
			return err
		}

		service := NewLogin(logger, flags)
		loginModel := models.LoginRequest{
			UserName:   in.Username,
			Password:   in.Password,
			RememberMe: in.RememberMe,
		}
		token, err := service.GetToken(loginModel, "")

		if err != nil {
			return err
		}
		if err := stream.Send(&rpc_auth.LoginResponse{
			LoginDetails: in,
			Token:        token,
		}); err != nil {
			return err
		}
	}
	return nil
}
