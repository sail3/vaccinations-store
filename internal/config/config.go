package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var serviceVersion = "local"

const (
	port           = "HTTP_PORT"
	debug          = "DEBUG"
	dbHost         = "DB_HOST"
	dbPort         = "DB_PORT"
	dbUser         = "DB_USER"
	dbPassword     = "DB_PASS"
	dbDatabase     = "DB_DATABASE"
	dbTimeout      = "DB_TIMEOUT"
	signatureToken = "SIGNATURE_TOKEN"
)

type Config struct {
	Port           string
	Debug          bool
	DbURL          string
	TimeOut        int
	SignatureToken string
}

func New() Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.Get(dbUser),
		viper.Get(dbPassword),
		viper.Get(dbHost),
		viper.Get(dbPort),
		viper.Get(dbDatabase))
	return Config{
		Debug:          GetEnvBool(debug, false),
		Port:           GetEnvString(port, "8080"),
		DbURL:          databaseURL,
		TimeOut:        GetEnvFloat(dbTimeout, 15),
		SignatureToken: GetEnvString(signatureToken, ""),
	}
}

// GetEnvString returns the value as a string of the environment
// variable that matches the key, if not found it returns the default value.
func GetEnvString(key, defaultValue string) string {
	if val := viper.GetString(key); val != "" {
		return val
	}

	return defaultValue
}

// GetEnvBool returns the value as boolean of the environment
// variable that matches the key, if not found it returns the default value.
func GetEnvBool(key string, defaultValue bool) bool {
	if val := viper.GetString(key); val != "" {
		bVal, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		return bVal
	}

	return defaultValue
}

func GetEnvFloat(key string, def int) int {
	val, err := strconv.Atoi(GetEnvString(key, fmt.Sprintf("%d", def)))
	if err != nil {
		return def
	}
	return val
}
