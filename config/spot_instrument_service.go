package config

type SpotInstrumentConfig struct {
	Log_LVL string `env:"LOG_LVL" validate:"required"`
}
