package config

import (
	"fmt"
	"os"
	"strconv"
)

type ConfigMap struct {
	ApiPort                 int
	GRPCPort                int
	ServiceName             string
	OtelEnabled             bool
	KafkaBrokers            string
	KafkaTopic              string
	KafkaGroupID            string
}

var Config *ConfigMap

func LoadEnvs() *ConfigMap {
	if Config != nil {
		return Config
	}

	Config = &ConfigMap{
		ApiPort:                 GetEnvNumber("API_PORT"),
		GRPCPort:                GetEnvNumber("GRPC_PORT"),
		ServiceName:             GetEnvString("SERVICE_NAME"),
		OtelEnabled:             GetEnvBool("OTEL_ENABLED", false),
		KafkaBrokers:            GetEnvString("KAFKA_BROKERS"),
		KafkaTopic:              GetEnvString("KAFKA_TOPIC"),
		KafkaGroupID:            GetEnvString("KAFKA_GROUP_ID"),
	}

	return Config
}

func GetEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value == "true"
}

func GetEnvString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}
	return value
}

func GetEnvNumber(key string) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s must be a valid integer", key))
	}

	return value
}
