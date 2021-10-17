package grpcserver

import (
	"context"
	"report-ms/config"
	"report-ms/server/grpcserver/usergrpc"

	"google.golang.org/grpc"
)

var (
	// StartUserGrpcClient starts gRPC client for user grpc
	StartUserGrpcClient = startUserGrpcClient
)

func startUserGrpcClient() (err error) {
	address := config.GetUserGRPCConnectionString()

	opts := grpc.WithInsecure()
	clientConn, err := grpc.Dial(address, opts)
	if err != nil {
		return
	}

	userGrpcClient = usergrpc.NewUserGrpcClient(clientConn)

	return
}

var userGrpcClient usergrpc.UserGrpcClient

func AuthorizeUser(ctx context.Context, authorizationToken string) (userRole int64, err error) {
	request := &usergrpc.AuthorizeUserRequest{AuthorizationToken: authorizationToken}

	response, err := userGrpcClient.AuthorizeUser(ctx, request)
	if err != nil {
		return
	}

	userRole = response.UserRole

	return
}
