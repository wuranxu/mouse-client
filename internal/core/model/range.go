package model

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	UserNotValid   = errors.New("user num is not valid")
	PeriodNotValid = errors.New("ram range not valid")
)

type RangeModel struct {
	user   int
	period int
	// end time
	deadline <-chan time.Time
	wg       sync.WaitGroup
}

func (r *RangeModel) Run(ctx context.Context, worker IRunner) error {
	step := r.user / r.period
	var current int
	r.wg.Add(r.user)
	for current < r.user {
		current += step
		users := step
		if current >= r.user {
			users = step + r.user - current
		}
		for i := 1; i <= users; i++ {
			go func(c context.Context) {
				defer r.wg.Done()
				for {
					select {
					case <-c.Done():
						return
					case <-r.deadline:
						return
					default:
						worker.Run(c)
					}
				}
			}(ctx)
		}
		time.Sleep(time.Second)
	}
	r.wg.Wait()
	return nil
}

func NewRangeModel(user, period int, last time.Duration) (*RangeModel, error) {
	if user <= 0 {
		return nil, UserNotValid
	}
	if period <= 0 {
		return nil, PeriodNotValid
	}
	res := &RangeModel{
		user:   user,
		period: period,
	}
	if last > 0 {
		res.deadline = time.After(last)
	}
	return res, nil
}
