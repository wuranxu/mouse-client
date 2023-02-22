package api

import (
	"context"
	"github.com/wuranxu/mouse-client/proto"
	tool "github.com/wuranxu/mouse-tool"
)

type MouseServiceApi struct {
	proto.UnimplementedMouseServiceServer
	sceneId int64
	server  proto.MouseService_ConnectServer
	quit    chan struct{}
	Client  *tool.EtcdClient
}

func (m *MouseServiceApi) work() {
	// TODO  do something from console
}

func (m *MouseServiceApi) Connect(srv proto.MouseService_ConnectServer) error {
	m.server = srv
	go m.work()
	return nil
}

func (m *MouseServiceApi) Disconnect(ctx context.Context, msg *proto.Message) (*proto.MouseResponse, error) {
	m.quit <- struct{}{}
	m.server = nil
	return &proto.MouseResponse{Code: 0, Msg: "success"}, nil
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
