package api

import (
	"context"
	"log"

	"github.com/mwdev22/database"
	config "github.com/mwdev22/gocfg"
	"github.com/mwdev22/grpclib/grpcserver"
)

type Api struct {
	cfg    *config.Config
	cache  database.Cache
	server *grpcserver.Server
}

func New(cfg *config.Config, server *grpcserver.Server, opts ...func(*Api)) *Api {
	api := &Api{
		cfg:    cfg,
		server: server,
	}
	for _, opt := range opts {
		opt(api)
	}
	return api
}

func WithCache(cache database.Cache) func(*Api) {
	return func(api *Api) {
		api.cache = cache
	}
}

func (a *Api) Start(ctx context.Context) error {

	// a.server.RegisterService()

	addr, err := a.server.Start(ctx)
	if err != nil {
		log.Fatalf("couldn't run grpc server: %s", err)
	}

	log.Printf("listening on %s", addr)

	return nil
}
