package config

import (
	"gomodules/configmodule"
	"time"
)

var (
	// GetHTTPServerAddress returns server address from http section in toml file
	GetHTTPServerAddress = getHTTPServerAddress

	// GetHTTPReadTimeout returns read timeout from http section in toml file
	GetHTTPReadTimeout = getHTTPReadTimeout

	// GetHTTPWriteTimeout returns write timeout from http section in toml file
	GetHTTPWriteTimeout = getHTTPWriteTimeout
)

func getHTTPServerAddress() string {
	return configmodule.GetConfigString("http.server_address")
}

func getHTTPReadTimeout() time.Duration {
	return configmodule.GetConfigDuration("http.read_timeout")
}

func getHTTPWriteTimeout() time.Duration {
	return configmodule.GetConfigDuration("http.write_timeout")
}
