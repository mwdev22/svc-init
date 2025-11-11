package main

import (
	"context"
	"log"

	config "github.com/mwdev22/gocfg"
	"github.com/mwdev22/grpclib/grpcserver"
	"github.com/mwdev22/svc-init/api"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.New(
		config.WithDatabaseConfig(
			&config.DatabaseConfig{},
		),
	)

	server := grpcserver.New(
		cfg.Addr,
		grpcserver.WithCreds(insecure.NewCredentials()),
	)

	app := api.New(cfg, server)

	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatalf("error while running app: %v", err)
	}
}
