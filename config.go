package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}

func loadConfig() (config, error) {
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		return config{}, fmt.Errorf("get env config: %w", err)
	}
	return cfg, nil
}
