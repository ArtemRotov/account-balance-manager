package config

import (
	"fmt"
	"path"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Log    `yaml:"log"`
		PG     `yaml:"postgres"`
		Hasher `yaml:"hasher"`
		JWT    `yaml:"jwt"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"level"`
	}

	PG struct {
		URL         string `yaml:"url"`
		MaxPoolSize int    `yaml:"max_pool_size"`
	}

	Hasher struct {
		Salt string `yaml:"salt"`
	}

	JWT struct {
		SignKey  string        `yaml:"sign_key"`
		TokenTTL time.Duration `yaml:"token_ttl"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// err = cleanenv.UpdateEnv(cfg)
	// if err != nil {
	// 	return nil, fmt.Errorf("error updating env: %w", err)
	// }

	return cfg, nil
}
