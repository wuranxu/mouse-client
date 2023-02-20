package api

import (
	"context"
	"github.com/wuranxu/mouse-client/proto"
)

type MouseServiceApi struct {
	proto.UnimplementedMouseServiceServer
	tasks map[int64]proto.Task
}

func (m *MouseServiceApi) Connect(srv proto.MouseService_ConnectServer) error {
	return nil
}

func (m *MouseServiceApi) Disconnect(ctx context.Context, msg *proto.Message) (*proto.MouseResponse, error) {
	return nil, nil
}

func (m *MouseServiceApi) Stat(ctx context.Context, msg *proto.Message) (*proto.MouseResponse, error) {
	return nil, nil
}

func (m *MouseServiceApi) Start(ctx context.Context, task *proto.Task) (*proto.MouseResponse, error) {
	return nil, nil
}

func (m *MouseServiceApi) Stop(ctx context.Context, task *proto.StopTask) (*proto.MouseResponse, error) {
	return nil, nil
}
