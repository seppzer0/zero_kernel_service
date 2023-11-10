package worker

import (
	"zero-kernel-service/internal/infrastructure/config"

	"github.com/hibiken/asynq"
)

type Scheduler struct {
	scheduler *asynq.Scheduler
}

func NewScheduler(cfg config.Config) *Scheduler {
	return &Scheduler{
		scheduler: &asynq.Scheduler{},
	}
}
