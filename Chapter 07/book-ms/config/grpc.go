package config

import "gomodules/configmodule"

var (
	// GetUserGRPCConnectionString returns user grpc connection string from grpc section in toml file
	GetUserGRPCConnectionString = getUserGRPCConnectionString
)

func getUserGRPCConnectionString() string {
	return configmodule.GetConfigString("grpc.user_grpc_connection_string")
}
