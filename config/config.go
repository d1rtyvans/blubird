package config

import (
    "os"
    "fmt"
)

// TODO: Only require this if it's prod environment...
type DbConfig struct {
    Password string
}

type Config struct {
    Db DbConfig
}

func New() *Config {
    return &Config{
        Db: DbConfig{
        },
    }
}

func getEnv(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    panic(fmt.Sprintf("'%s' key was not found in .env file.", key))
}
