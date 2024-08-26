package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		DB   `yaml:"db"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"4102"      env-required:"true" yaml:"port"`
		Host string `env:"HTTP_HOST" env-default:"localhost" env-required:"true" yaml:"host"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL" env-required:"true" yaml:"log_level"`
	}

	DB struct {
		Driver     string `yaml:"driver"`
		DataSource string `env:"data_source" env-required:"true" yaml:"data_source"`
	}
)

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) ParseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
