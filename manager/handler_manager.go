package manager

import (
	"context"
	"core_x/contract"
	"core_x/contract/proto"
	"core_x/errors"
	"github.com/urfave/cli/v2"
)

type HandlerManager struct {
	cxt context.Context
	actionManager contract.IActionManager
	list map[string]contract.IClientHandle
}

func NewHandlerManager() contract.IModule  {
	return &HandlerManager {
		list: make(map[string]contract.IClientHandle),
	}
}


func (h HandlerManager) Init(c cli.Context) error {


	return nil
}

func (h *HandlerManager) Load() contract.AppLoad {
	return func(context *contract.AppContext) {
		h.actionManager = context.ActionManager
		context.HandlerManager = h
	}
}

func (h HandlerManager) Flag() []cli.Flag {
	return []cli.Flag{}
}

func (h HandlerManager) Priority() int {
	return 2
}

func (h *HandlerManager) Context(ctx context.Context) {
	h.cxt = ctx
}

func (h HandlerManager) GetClientById(id string) (contract.IClientHandle, error) {
	module, exist := h.list[id]
	if !exist {
		return nil, errors.New("Module Exist", errors.EXIST)
	}

	return module, nil
}

func (h HandlerManager) GetClientByModule(moduleName string) (contract.IClientHandle, error) {

	var module contract.IClientHandle
	var exist = false
	for _, item := range h.list {
		if item.Id() == moduleName {
			module = item
		}
	}
	if !exist {
		return nil, errors.New("Module Exist", errors.EXIST)
	}
	return module, nil
}

func (h HandlerManager) SendToModule(moduleName string, event *proto.Event) error {
	module, err := h.GetClientByModule(moduleName)
	if err != nil {
		return err
	}
	return module.Send(event)
}

func (h HandlerManager) SendToId(id string, event *proto.Event) error {
	module, err := h.GetClientById(id)
	if err != nil {
		return err
	}
	return module.Send(event)
}

func (h HandlerManager) AddConnect(client contract.IClientHandle) error {
	if h.HasConnect(client.Id()) {
		return errors.New("Module Exist", errors.EXIST)
	}
	h.list[client.Id()] = client
	return nil
}

func (h HandlerManager) RemoveConnect(id string) error {
	if h.HasConnect(id) {
		return errors.New("Module Not Found", errors.NOT_FOUND)
	}
	return nil
}

func (h HandlerManager) Process(event *proto.Event) error {
	if !h.actionManager.HasAction(event.Name) {
		return errors.New("Not Found", errors.EXIST)
	}
	handlers, err := h.actionManager.GetHandler(event.Name)
	if err != nil {
		return err
	}

	for _, item := range handlers {
		h.SendToModule(item, event)
	}

	return nil
}

func (h HandlerManager) HasConnect(id string) bool {
	_, exist := h.list[id]
	return exist
}

