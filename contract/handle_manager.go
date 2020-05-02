package contract

import (
	"core_x/contract/proto"
)

type IHandleManager interface {
	SendToModule(moduleName string, event *proto.Event) error
	SendToId(id string, event *proto.Event) error
	GetClientByModule(module string) (IClientHandle, error)
	GetClientById(id string) (IClientHandle, error)
	AddConnect(client IClientHandle) error
	RemoveConnect(id string) error
	Process(event *proto.Event) error
	HasConnect(id string) bool
}