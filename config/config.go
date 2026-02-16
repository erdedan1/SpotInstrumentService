package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	GRPCServer            GRPCServerConfig
	GRPCClient            GRPCClientConfig
	SpotInstrumentService SpotInstrumentConfig
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	v := validator.New() //todo мб поменять валидатор
	if err := v.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
