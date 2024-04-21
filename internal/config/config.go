package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ConnectionString string `yaml:"conn_str"`
	ServerAddress    string `yaml:"address"`
	JwtSecret        string `yaml:"jwt_secret"`
	DriverName       string `yaml:"driver_name"`
	MigrationsDir    string `yaml:"migrations_dir"`
}

func LoadConfig(configPath string) (Config, error) {
	config := Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
