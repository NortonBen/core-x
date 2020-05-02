package models

import (
	"core_x/errors"
	"time"
)

type TypeAction int32
type StateModule int32

const (
	TypeAction_Hook     TypeAction = 0
	TypeAction_EndPoint TypeAction = 1
	StateModule_RUN		StateModule = 1
	StateModule_STOP	StateModule = 0
)

type RegisterAction struct {
	Type                 TypeAction `protobuf:"varint,1,opt,name=type,proto3,enum=proto.TypeAction" json:"type,omitempty"`
	Action               string     `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Priority             int32      `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Description          string     `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
}


type ModuleData struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Auth                 []string          `protobuf:"bytes,4,rep,name=auth,proto3" json:"auth,omitempty"`
	Url                  string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Actions              []*ActionData     `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions,omitempty"`
	Registers            []*RegisterAction `protobuf:"bytes,7,rep,name=registers,proto3" json:"registers,omitempty"`
	Time *time.Time
	State StateModule
	Section 			string
	Token 				string
}

func (m *ModuleData) Validate() error {
	if len(m.Id) == 0 {
		return errors.New("id_not_validate", errors.VALIDATE)
	}
	if len(m.Token) == 0 {
		return errors.New("token_not_validate", errors.VALIDATE)
	}
	if len(m.Name) == 0 {
		return errors.New("name_not_validate", errors.VALIDATE)
	}
	if len(m.Version) == 0 {
		return errors.New("version_not_validate", errors.VALIDATE)
	}

	return nil
}