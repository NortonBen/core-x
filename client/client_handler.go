package client

import (
	"core_x/contract"
	"core_x/contract/proto"
)

type ClientHandler struct {
	stream proto.CoreService_EventStreamServer
	id string
	token string
	module string
}

func NewClientHandler(id string, token string, module string, stream proto.CoreService_EventStreamServer) contract.IClientHandle  {
	client := &ClientHandler{
		stream: stream,
		id:     id,
		token:  token,
		module: module,
	}

	return client
}


func (c ClientHandler) Stream() proto.CoreService_EventStreamServer {
	return c.stream
}

func (c ClientHandler) Id() string {
	return c.id
}

func (c ClientHandler) Token() string {
	return c.token
}

func (c ClientHandler) Module() string {
	return c.module
}

func (c ClientHandler) Send(event *proto.Event) error {
	return c.stream.Send(event)
}



