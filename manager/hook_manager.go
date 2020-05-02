package manager

import (
	"core_x/contract"
	"core_x/contract/proto"
	"core_x/errors"
	"context"
	"github.com/urfave/cli/v2"
)

type HookManager struct {
	cxt context.Context
	clients map[string]contract.IClientHook
	actionManger contract.IActionManager
}

func NewHookManager() contract.IModule  {
	return &HookManager {
		clients: make(map[string]contract.IClientHook),
	}
}


func (h HookManager) Init(c cli.Context) error {


	return nil
}

func (h *HookManager) Load() contract.AppLoad {
	return func(context *contract.AppContext) {
		h.actionManger = context.ActionManager
		context.HookManager = h
	}
}

func (h HookManager) Flag() []cli.Flag {
	return []cli.Flag{}
}

func (h HookManager) Priority() int {
	return 2
}

func (h *HookManager) Context(ctx context.Context) {
	h.cxt = ctx
}


func (h HookManager) GetClient(id string) (contract.IClientHook, error) {
	client, exist := h.clients[id]

	if !exist {
		return nil, errors.New("Client have exist", errors.EXIST)
	}
	return client, nil
}

func (h HookManager) GetClientByModule(module string) (contract.IClientHook, error) {

	var client contract.IClientHook
	var exist bool = false

	for _, val := range h.clients {
		if val.Module() == module {
			client = val
			exist = true
			break
		}
	}

	if !exist {
		return nil, errors.New("Client have exist", errors.EXIST)
	}
	return client, nil
}

func (h HookManager) SendToModule(ctx context.Context, moduleName string, event *proto.Event) (*proto.Event, error) {
	client, err := h.GetClientByModule(moduleName)
	if err != nil {
		return nil, err
	}
	err = client.Send(event)
	if err != nil {
		return nil, err
	}
	rs, err := client.RecvWithIdEvent(ctx, event.Id)
	if err != nil {
		return  nil, err
	}
	return rs, nil
}

func (h HookManager) AddConnect(client contract.IClientHook) error {
	if h.HaveConnect(client.Id()) {
		return errors.New("Client have exist", errors.EXIST)
	}
	h.clients[client.Id()] = client
	return nil
}

func (h HookManager) RemoveConnect(id string) error {
	if h.HaveConnect(id) {
		return errors.New("Client have exist", errors.EXIST)
	}
	delete(h.clients, id)
	return nil
}

func (h HookManager) Process(event *proto.Event) error {
	if !h.actionManger.HasAction(event.Name) {
		return errors.New("Action have exist", errors.EXIST)
	}
	moduleList, err := h.actionManger.GetHook(event.Name)
	if err != nil {
		return err
	}

	action, err := h.actionManger.Get(event.Name)
	if err != nil {
		return err
	}



	for _, module := range moduleList {
		ctx := context.Background()
		if action.TimeOut != 0 {
			ctx, _ = context.WithTimeout(ctx, action.TimeOut)
		}
		rs, err := h.SendToModule(ctx, module, event)
		if err != nil {

		} else {
			event = rs
		}
	}

	return nil
}

func (h HookManager) HaveConnect(id string) bool {
	_, exist := h.clients[id]
	return exist
}

