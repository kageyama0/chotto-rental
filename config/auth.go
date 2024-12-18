package config

import (
	"os"
	"time"
)

// init: auth関連の設定を初期化
func initAuthConfig() (*AuthConfig, error) {
	return &AuthConfig{
		JWTSecret:         []byte(os.Getenv("JWT_SECRET")),
		AccessTokenExpiry: time.Hour,
		SessionExpiry:     30 * 24 * time.Hour,
	}, nil
}
