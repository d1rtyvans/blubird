package config

import (
    "os"
    "fmt"
)

// TODO: Only require this if it's prod environment...
type DbConfig struct {
    Password string
}

type WeatherApiConfig struct {
    ApiKey string
}

type Config struct {
    Db      DbConfig
    DarkSky WeatherApiConfig
}

func New() *Config {
    return &Config{
        // Db: DbConfig{ }, // TODO
        DarkSky: WeatherApiConfig{
            ApiKey: getEnv("DARK_SKY_API_KEY"),
        },
    }
}

func getEnv(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    panic(fmt.Sprintf("'%s' key was not found in .env file.", key))
}
