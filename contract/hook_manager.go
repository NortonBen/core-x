package contract

import (
	"core_x/contract/proto"
	"context"
)

type IHookManager interface {
	SendToModule(ctx context.Context, moduleName string, event *proto.Event) (*proto.Event, error)
	AddConnect(client IClientHook) error
	RemoveConnect(id string) error
	Process(event *proto.Event) error
	HaveConnect(id string) bool
	GetClient(id string) (IClientHook, error)
	GetClientByModule(module string) (IClientHook, error)
}

