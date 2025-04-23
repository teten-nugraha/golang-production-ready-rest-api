package config

import "os"

type Config struct {
	AppPort         string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	RedisAddr       string
	JwtSecret       string
	RateLimit       string
	CORSAllowOrigin string
	DBSSLMode       string // Add SSL mode for PostgreSQL
}

func NewConfig() *Config {
	return &Config{
		AppPort:         os.Getenv("APP_PORT"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		RedisAddr:       os.Getenv("REDIS_ADDR"),
		JwtSecret:       os.Getenv("JWT_SECRET"),
		RateLimit:       os.Getenv("RATE_LIMIT"),
		CORSAllowOrigin: os.Getenv("CORS_ALLOW_ORIGINS"),
		DBSSLMode:       os.Getenv("DB_SSL_MODE"),
	}
}
