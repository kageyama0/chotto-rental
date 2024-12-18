package config

import "time"

type AuthConfig struct {
	JWTSecret         []byte
	AccessTokenExpiry time.Duration
	SessionExpiry     time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	Timezone string
}

// config/config.go
type Config struct {
	Auth     AuthConfig
	Database DatabaseConfig
	// Server  ServerConfig
	// ... 他の設定カテゴリ
}
