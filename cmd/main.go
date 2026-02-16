package main

import (
	"SpotInstrumentService/config"
	"SpotInstrumentService/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	spotInstrumentSerivce := app.New(cfg)
	defer spotInstrumentSerivce.L.Sync()

	if err := spotInstrumentSerivce.Start(); err != nil {
		panic(err)
	}
}
