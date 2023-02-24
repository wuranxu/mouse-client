package model

import (
	"context"
	"errors"
	"log"
)

var (
	QpsNotValid = errors.New("qps not valid")
)

type IRunner interface {
	Run(context.Context)
	Stat(context.Context)
}

// RateLimiterModel limit qps model
type RateLimiterModel struct {
	qps     int
	limiter *RateLimiter
}

func (r *RateLimiterModel) Run(ctx context.Context, worker IRunner) error {
	go worker.Stat(ctx)
	for {
		select {
		case <-ctx.Done():
			// stop task
			return nil
		default:
			if err := r.limiter.Wait(ctx); err != nil {
				select {
				case <-ctx.Done():
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

func NewRateLimiterModel(qps int) (IModel, error) {
	if qps <= 0 {
		return nil, QpsNotValid
	}
	res := &RateLimiterModel{
		qps:     qps,
		limiter: NewRateLimiter(float64(qps), qps),
	}
	return res, nil
}
