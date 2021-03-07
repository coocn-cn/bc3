package service

import (
	"context"
	"time"

	"github.com/coocn-cn/bc3/pkg/log"
	"github.com/miniflux/v2/storage"
	"github.com/miniflux/v2/worker"
)

func feedScheduler(ctx context.Context, store *storage.Storage, pool *worker.Pool, frequency, batchSize int) {
	log := log.G(ctx).WithContext(ctx)

	for range time.Tick(time.Duration(frequency) * time.Minute) {
		jobs, err := store.NewBatch(batchSize)
		if err != nil {
			log.WithError(err).Error("Create batch failed")
		} else {
			log.WithField("count", len(jobs)).Error("Pushing jobs")
			pool.Push(jobs)
		}
	}
}
