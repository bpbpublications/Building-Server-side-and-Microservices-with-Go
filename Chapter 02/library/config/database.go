package config

import "time"

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
	return getConfigString("database.connection_string")
}

func getDatabaseMaxIdleConnections() int {
	return getConfigInt("database.max_idle_connections")
}

func getDatabaseMaxOpenConnections() int {
	return getConfigInt("database.max_open_connections")
}

func getDatabaseConnectionMaxLifetime() time.Duration {
	return getConfigDuration("database.connection_max_lifetime")
}
