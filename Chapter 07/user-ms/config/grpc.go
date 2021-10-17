package config

import "gomodules/configmodule"

var (
	// GetGrpcConnectionString returns connection string from grpc section in toml file
	GetGrpcConnectionString = getGrpcConnectionString
)

func getGrpcConnectionString() string {
	return configmodule.GetConfigString("grpc.connection_string")
}
