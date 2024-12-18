package config

import (
	"os"
)

// init: auth関連の設定を初期化
func initDBConfig() (*DatabaseConfig, error) {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  "disable",
		Timezone: "Asia/Tokyo",
	}, nil
}
