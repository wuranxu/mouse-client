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

	limiterModel, err := model.NewRateLimiterModel(1, 30*time.Second)
	if err != nil {
		t.Error("create model failed", err)
		return
	}

	limiterModel.Run(context.TODO(), runner)
}
