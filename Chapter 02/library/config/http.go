package config

import "time"

var (
	// GetHTTPServerAddress returns server address from http section in toml file
	GetHTTPServerAddress = getHTTPServerAddress

	// GetHTTPReadTimeout returns read timeout from http section in toml file
	GetHTTPReadTimeout = getHTTPReadTimeout

	// GetHTTPWriteTimeout returns write timeout from http section in toml file
	GetHTTPWriteTimeout = getHTTPWriteTimeout
)

func getHTTPServerAddress() string {
	return getConfigString("http.server_address")
}

func getHTTPReadTimeout() time.Duration {
	return getConfigDuration("http.read_timeout")
}

func getHTTPWriteTimeout() time.Duration {
	return getConfigDuration("http.write_timeout")
}
