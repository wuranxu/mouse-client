package model

import (
	"context"
	"errors"
	"github.com/wuranxu/mouse-client/internal/core/scene"
	"log"
	"time"
)

var (
	QpsNotValid = errors.New("qps not valid")
)

type StressModel interface {
	Run(context.Context, *scene.Scene) error
}

type IRunner interface {
	Run(context.Context)
}

// RateLimiterModel limit qps model
type RateLimiterModel struct {
	qps      int
	deadline <-chan time.Time
	limiter  *RateLimiter
}

func (r *RateLimiterModel) Run(ctx context.Context, worker IRunner) error {
	for {
		select {
		case <-ctx.Done():
			// stop task
			return nil
		case <-r.deadline:
			return nil
		default:
			if err := r.limiter.Wait(ctx); err != nil {
				select {
				case <-ctx.Done():
					return nil
				case <-r.deadline:
					return nil
				default:
					log.Println("rate limiter pool get task failed, ", err)
					continue
				}
			}
			go worker.Run(ctx)
		}
	}
}

func NewRateLimiterModel(qps int, last time.Duration) (*RateLimiterModel, error) {
	if qps <= 0 {
		return nil, QpsNotValid
	}
	res := &RateLimiterModel{
		qps:     qps,
		limiter: NewRateLimiter(float64(qps), qps),
	}
	if last > 0 {
		res.deadline = time.After(last)
	}
	return res, nil
}
