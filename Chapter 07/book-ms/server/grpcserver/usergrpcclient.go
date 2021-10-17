package grpcserver

import (
	"book-ms/config"
	"book-ms/server/grpcserver/usergrpc"
	"context"

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

func GetUserID(ctx context.Context, authorizationToken string) (userID string, err error) {
	request := &usergrpc.GetUserIDRequest{AuthorizationToken: authorizationToken}

	response, err := userGrpcClient.GetUserID(ctx, request)
	if err != nil {
		return
	}

	userID = response.UserID

	return
}

func GetUserFullName(ctx context.Context, userID string) (fullName string, err error) {
	request := &usergrpc.GetUserFullNameRequest{UserID: userID}

	response, err := userGrpcClient.GetUserFullName(ctx, request)
	if err != nil {
		return
	}

	fullName = response.FullName

	return
}
