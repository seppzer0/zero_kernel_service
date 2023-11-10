package infrastructure

import (
	"zero-kernel-service/internal/infrastructure/config"

	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
)
