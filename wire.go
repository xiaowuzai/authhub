//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"git.corp.zgcszkw.com/authhub/backend"
	"git.corp.zgcszkw.com/authhub/config"
	"git.corp.zgcszkw.com/authhub/server"
	"git.corp.zgcszkw.com/authhub/service"
	"github.com/google/wire"
)

func initApp(ctx context.Context, cfg *config.Config) (*server.Server, error) {
	panic(wire.Build(
		service.ProviderSet,
		server.ProviderSet,
		backend.ProviderSet,
	))
}
