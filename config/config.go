package config

import "os"

type Config struct {
    NodeAPIBaseURL string
}

func Load() *Config {
    return &Config{
        NodeAPIBaseURL: getEnv(
            "NODE_API_BASE_URL",
            "http://localhost:5007/api/v1",
        ),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
