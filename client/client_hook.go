package client

import (
	"core_x/contract"
	"core_x/contract/proto"
	"core_x/errors"
	"context"
	"io"
)

type ClientHook struct {
	stream proto.CoreService_HooksServer
	id string
	token string
	module string

	list chan *proto.HookStream
}


func (c ClientHook) Recv(ctx context.Context, callback contract.CallBackFuncEvent) error {
	go func() {
		for  {
			event, err := c.stream.Recv()
			if err == io.EOF {
				callback(event, true)
				return
			}
			callback(event, false)
		}
	}()

	return nil
}

func NewClientHook(id string, token string, module string, stream proto.CoreService_HooksServer) *ClientHook  {
	client := &ClientHook{
		stream: stream,
		id:     id,
		token:  token,
		module: module,
		list: make(chan *proto.HookStream),
	}

	return client
}

func (c ClientHook) Loop() error  {
	for  {
		event, err := c.stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		c.list <- event
	}
}

func (c ClientHook) Stream() proto.CoreService_HooksServer {
	return c.stream
}

func (c ClientHook) Id() string {
	return c.id
}

func (c ClientHook) Token() string {
	return c.token
}

func (c ClientHook) Module() string {
	return c.module
}

func (c ClientHook) Send(event *proto.Event) error {
	return c.stream.Send(event)
}

func (c ClientHook) RecvWithIdEvent(ctx context.Context, id string) (*proto.Event, error) {
	var event proto.Event
	for  {
		select {
			case <- ctx.Done(): {
				return nil, errors.Error{
					Message: "TimeOut",
					Code: errors.TIMEOUT,
				}
			}
			case event = <- c.list: {
				if event.Id == id {
					break;
				}
			}
		}
	}

	return &event, nil
}






