package config

import "gomodules/configmodule"

var (
	// GetRabbitMQConnectionString returns connection string from rabbitMQ section in toml file
	GetRabbitMQConnectionString = getRabbitMQConnectionString
)

func getRabbitMQConnectionString() string {
	return configmodule.GetConfigString("rabbitmq.connection_string")
}
