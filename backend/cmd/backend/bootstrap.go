package main

import (
	"context"

	"s0nh-server/internal/infrastructure"
	"s0nh-server/pkg/server"

	"go.uber.org/fx"
)

var Module = fx.Options(
	server.Module,
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	srv server.Server,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func(srv server.Server) {
				if err := srv.ListenAndServe(); err != nil {
					panic(err)
				}
			}(srv)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
