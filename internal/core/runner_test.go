package core

import (
	"context"
	"github.com/wuranxu/mouse-client/internal/core/model"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	data, err := ioutil.ReadFile("D:\\projects\\github\\wuranxu\\mouse-client\\internal\\core\\test_data.yaml")
	if err != nil {
		t.Error("read file failed", err)
		return
	}
	runner, err := NewRunner(1, data)
	if err != nil {
		t.Error("start runner failed", err)
		return
	}

	//limiterModel, err := model.NewRateLimiterModel(20, 300*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	limiterModel, err := model.NewRangeModel(50, 1)
	if err != nil {
		t.Error("create model failed", err)
		return
	}

	limiterModel.Run(ctx, runner)
}
