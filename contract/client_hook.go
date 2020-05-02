package contract

import (
	"core_x/contract/proto"
	"context"
)

type CallBackFuncEvent func(stream *proto.HookStream, end bool)

type IClientHook interface {
	Stream() proto.CoreService_HooksServer
	Id() string
	Token() string
	Module() string
	Send(event *proto.Event) error
	RecvWithIdEvent(ctx context.Context, id string) (*proto.Event, error)
	Recv(ctx context.Context,callback CallBackFuncEvent)  error
}