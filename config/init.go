package config

func InitConfig() (*Config, error) {
	authConfig, err := initAuthConfig()
	if err != nil {
		return nil, err
	}

	databaseConfig, err := initDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Auth:     *authConfig,
		Database: *databaseConfig,
	}, nil
}
