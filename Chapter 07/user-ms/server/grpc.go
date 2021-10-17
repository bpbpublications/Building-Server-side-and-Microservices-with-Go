package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-ms/config"
	"user-ms/core"
	"user-ms/server/usergrpc"

	"google.golang.org/grpc"
)

var (
	// StartGrpcServer starts listening to gRPC requests
	StartGrpcServer = startGrpcServer
)

func startGrpcServer() (err error) {
	address := config.GetGrpcConnectionString()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return
	}

	server := grpc.NewServer()
	usergrpc.RegisterUserGrpcServer(server, &userGrpcServer{})

	go func() {
		// Listen to operating system's interrupt signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt

		// Gracefully shut down the server when it happens
		server.GracefulStop()
	}()

	server.Serve(lis)

	return
}

type userGrpcServer struct{}

func (s *userGrpcServer) AuthorizeUser(ctx context.Context, request *usergrpc.AuthorizeUserRequest) (response *usergrpc.AuthorizeUserResponse, err error) {
	userRole, err := core.AuthorizeUser(ctx, request.AuthorizationToken)
	if err != nil {
		return
	}

	response = &usergrpc.AuthorizeUserResponse{UserRole: userRole}

	return
}

func (s *userGrpcServer) GetUserID(ctx context.Context, request *usergrpc.GetUserIDRequest) (response *usergrpc.GetUserIDResponse, err error) {
	userID, err := core.GetUserID(ctx, request.AuthorizationToken)
	if err != nil {
		return
	}

	response = &usergrpc.GetUserIDResponse{UserID: userID}

	return
}

func (s *userGrpcServer) GetUserFullName(ctx context.Context, request *usergrpc.GetUserFullNameRequest) (response *usergrpc.GetUserFullNameResponse, err error) {
	fullName, err := core.GetUserFullName(ctx, request.UserID)
	if err != nil {
		return
	}

	response = &usergrpc.GetUserFullNameResponse{FullName: fullName}

	return
}
