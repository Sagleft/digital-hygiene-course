package main

import (
	"fmt"
	"log"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/tgfun"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("app started")
	swissknife.RunInBackground()
}

func run() error {
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	_, err = tgfun.NewFunnel(
		getData(cfg),
		getScript(),
	)
	if err != nil {
		return fmt.Errorf("create funnel: %w", err)
	}
	return nil
}
