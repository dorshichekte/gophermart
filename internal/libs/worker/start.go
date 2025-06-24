package worker

import (
	"context"
	"time"
)

func (w *Worker) Start(ctx context.Context, tasks ...func(ctx context.Context)) {
	timer := time.NewTicker(5 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			for _, task := range tasks {
				task(ctx)
			}
		case <-ctx.Done():
			return
		}
	}
}
