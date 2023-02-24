package model

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	UserNotValid   = errors.New("user num is not valid")
	PeriodNotValid = errors.New("ram range not valid")
)

type IModel interface {
	Run(ctx context.Context, worker IRunner) error
}

type RangeModel struct {
	user   int
	period int
	wg     sync.WaitGroup
}

func (r *RangeModel) Run(ctx context.Context, worker IRunner) error {
	step := r.user / r.period
	if step == 0 {
		step = 1
	}
	var current int
	go worker.Stat(ctx)
	r.wg.Add(r.user)
	for current < r.user {
		current += step
		users := step
		if current >= r.user {
			users = step + r.user - current
		}
		for i := 1; i <= users; i++ {
			go func(c context.Context) {
				defer func() {
					r.wg.Done()
					fmt.Println("wanle?")
				}()
				for {
					select {
					case <-c.Done():
						return
					default:
						worker.Run(c)
					}
				}
			}(ctx)
		}
		time.Sleep(time.Second)
	}
	go r.wg.Wait()
	return nil
}

func NewRangeModel(user, period int) (IModel, error) {
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
	return res, nil
}
