package main

import (
	"flat/api"
	"log"

	"github.com/mwdev22/core/config"
)

// @title           REST Boilerplate API
// @version         1.0
// @description     API documentation

func main() {
	cfg := config.New()

	app := api.New(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("error while running app: %v", err)
	}
}
