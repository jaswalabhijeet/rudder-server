package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Initialize initializes the config
func Initialize() {
	configPath := GetEnv("CONFIG_PATH", "./config.toml")

	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// GetInt is wrapper for viper's GetInt
func GetInt(key string, defaultValue int) int {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetInt(key)
}

// GetFloat64 is wrapper for viper's GetFloat64
func GetFloat64(key string, defaultValue float64) float64 {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetFloat64(key)
}

// GetString is wrapper for viper's GetString
func GetString(key string, defaultValue string) string {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetString(key)
}

// GetDuration is wrapper for viper's GetDuration
func GetDuration(key string, defaultValue time.Duration) time.Duration {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetDuration(key)
}

// GetEnv returns the environment value stored in key variable
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}