package config

import (
	"log"
	"os"
	"time"
)

func GetJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(secret)
}

func GetRefreshJwtSecret() []byte {
	secret := os.Getenv("REFRESH_JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("REFRESH_JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(secret)
}

func GetJWTExpirationDuration() time.Duration {
	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return time.Hour * 1
	}
	return duration
}

func GetRefreshJWTExpirationDuration() time.Duration {
	duration, err := time.ParseDuration(os.Getenv("REFRESH_JWT_EXPIRES_IN"))
	if err != nil {
		return time.Hour * 24 * 7
	}
	return duration
}
