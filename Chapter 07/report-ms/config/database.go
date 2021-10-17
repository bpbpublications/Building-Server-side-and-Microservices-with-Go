package config

import (
	"gomodules/configmodule"
	"time"
)

var (
	// GetDatabaseConnectionString returns connection string from database section in toml file
	GetDatabaseConnectionString = getDatabaseConnectionString

	// GetDatabaseMaxIdleConnections returns max idle connection from database section in toml file
	GetDatabaseMaxIdleConnections = getDatabaseMaxIdleConnections

	// GetDatabaseMaxOpenConnections returns max open connections from database section in toml file
	GetDatabaseMaxOpenConnections = getDatabaseMaxOpenConnections

	// GetDatabaseConnectionMaxLifetime returns connection max lifetime from database section in toml file
	GetDatabaseConnectionMaxLifetime = getDatabaseConnectionMaxLifetime
)

func getDatabaseConnectionString() string {
	return configmodule.GetConfigString("database.connection_string")
}

func getDatabaseMaxIdleConnections() int {
	return configmodule.GetConfigInt("database.max_idle_connections")
}

func getDatabaseMaxOpenConnections() int {
	return configmodule.GetConfigInt("database.max_open_connections")
}

func getDatabaseConnectionMaxLifetime() time.Duration {
	return configmodule.GetConfigDuration("database.connection_max_lifetime")
}
