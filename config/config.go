package config

import "os"

type PostgresConfig struct {
	PostgresUsername string
	PostgresPassword string
	PostgresDatabase string
	PostgresHostname string
	PostgresPort     string
	PostgresSslmode  string
}
type Config struct {
	PostgresConfig PostgresConfig
}

func NewConfig() *Config {
	return &Config{
		PostgresConfig: PostgresConfig{
			PostgresUsername: getEnv("POSTGRES_USER", "postgres"),
			PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
			PostgresDatabase: getEnv("POSTGRES_DATABASE", "teacher"),
			PostgresHostname: getEnv("POSTGRES_HOSTNAME", "localhost"),
			PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
			PostgresSslmode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
