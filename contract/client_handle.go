package contract

import "core_x/contract/proto"

type IClientHandle interface {
	Stream() proto.CoreService_EventStreamServer
	Id() string
	Token() string
	Module() string
	Send(event *proto.Event) error
}
