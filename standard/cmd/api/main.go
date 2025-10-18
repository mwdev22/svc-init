package main

import (
	"log"

	"github.com/mwdev22/core/config"
	"github.com/mwdev22/svc-init/internal/api"
)

// @title           REST Boilerplate API
// @version         1.0
// @description     API documentation for your Go project.
// @termsOfService  http://example.com/terms/

func main() {
	cfg := config.New()

	app := api.New(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("error while running app: %v", err)
	}
}
