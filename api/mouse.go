package api

import (
	"context"
	"errors"
	"github.com/wuranxu/mouse-client/internal/core"
	"github.com/wuranxu/mouse-client/internal/core/model"
	"github.com/wuranxu/mouse-client/proto"
	tool "github.com/wuranxu/mouse-tool"
)

const (
	StartJobErrorCode = iota + 10001
	ExistsJobErrorCode
	NotExistsJobErrorCode
)

var (
	ErrorStartJob     = errors.New("failed to start job")
	ErrorJobExists    = errors.New("job already exists")
	ErrorJobNotExists = errors.New("job not exists")
)

const Success = "operate success"

func WrapMsg(err error, msg string) string {
	if err == nil {
		return msg
	}
	return err.Error() + ": " + msg
}

type MouseServiceApi struct {
	proto.UnimplementedMouseServiceServer
	sceneId int64
	Client  *tool.EtcdClient
	cancel  context.CancelFunc
	runner  *core.Runner
}

//func (m *MouseServiceApi) Stat(ctx context.Context, _ *proto.Empty) (*proto.MouseResponse, error) {
//	return nil, nil
//}

func (m *MouseServiceApi) reset() {
	m.cancel = nil
	m.sceneId = 0
	m.runner = nil
}

func (m *MouseServiceApi) Start(ctx context.Context, task *proto.Task) (*proto.MouseResponse, error) {
	if m.sceneId != 0 {
		return &proto.MouseResponse{Code: ExistsJobErrorCode, Msg: ErrorJobExists.Error()}, nil
	}

	resp := &proto.MouseResponse{Code: 0, Msg: Success}
	runner, err := core.NewRunner(task.SceneId, task.Data)
	if err != nil {
		m.reset()
		resp.Code = StartJobErrorCode
		resp.Msg = WrapMsg(ErrorStartJob, err.Error())
		return resp, nil
	}
	var (
		md model.IModel
	)
	if *task.MaxQps != 0 {
		// qps mode
		md, err = model.NewRateLimiterModel(int(*task.MaxQps))
	} else {
		md, err = model.NewRangeModel(int(*task.Threads), int(*task.Interval))
	}
	if err != nil {
		m.reset()
		resp.Code = StartJobErrorCode
		resp.Msg = WrapMsg(ErrorStartJob, err.Error())
		return resp, nil
	}

	ct, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	if err = md.Run(ct, runner); err != nil {
		m.reset()
		resp.Code = StartJobErrorCode
		resp.Msg = WrapMsg(ErrorStartJob, err.Error())
		return resp, nil
	}
	return resp, nil
}

func (m *MouseServiceApi) Stop(ctx context.Context, task *proto.StopTask) (*proto.MouseResponse, error) {
	resp := &proto.MouseResponse{Code: 0, Msg: Success}
	if m.sceneId != task.SceneId {
		return &proto.MouseResponse{Code: NotExistsJobErrorCode, Msg: ErrorJobNotExists.Error()}, nil
	}
	// stop job
	if m.cancel == nil {
		return &proto.MouseResponse{Code: NotExistsJobErrorCode, Msg: ErrorJobNotExists.Error()}, nil
	}
	m.cancel()
	// reset job
	m.reset()
	return resp, nil
}
