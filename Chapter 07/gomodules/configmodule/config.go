package configmodule

import (
	"time"

	"github.com/spf13/viper"
)

var (
	// InitConfig reads configuration from TOML file
	InitConfig = initConfig

	// GetConfigString reads string from coniguration TOML file
	GetConfigString = getConfigString

	// GetConfigDuration reads time.Duration from coniguration TOML file
	GetConfigDuration = getConfigDuration

	// GetConfigInt reads int from coniguration TOML file
	GetConfigInt = getConfigInt
)

func initConfig(fileName string, additionalDirs []string) (err error) {
	viper.SetConfigName(fileName)

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	for _, dir := range additionalDirs {
		viper.AddConfigPath(dir)
	}

	//Read configuration file from disk
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Create configuration
	viper.ConfigFileUsed()
	viper.WatchConfig()

	return
}

func getConfigString(key string) string {
	return viper.GetString(key)
}

func getConfigInt(key string) int {
	return viper.GetInt(key)
}

func getConfigDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
