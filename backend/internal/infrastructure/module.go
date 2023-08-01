package infrastructure

import (
	"s0nh-server/internal/infrastructure/config"

	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
)
