// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TypeAction int32

const (
	TypeAction_Hook     TypeAction = 0
	TypeAction_EndPoint TypeAction = 1
)

var TypeAction_name = map[int32]string{
	0: "Hook",
	1: "EndPoint",
}

var TypeAction_value = map[string]int32{
	"Hook":     0,
	"EndPoint": 1,
}

func (x TypeAction) String() string {
	return proto.EnumName(TypeAction_name, int32(x))
}

func (TypeAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

type ActionRequest struct {
	IdModule             string            `protobuf:"bytes,1,opt,name=id_module,json=idModule,proto3" json:"id_module,omitempty"`
	IdSection            string            `protobuf:"bytes,2,opt,name=id_section,json=idSection,proto3" json:"id_section,omitempty"`
	MetaData             map[string]string `protobuf:"bytes,3,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 *Event            `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetIdModule() string {
	if m != nil {
		return m.IdModule
	}
	return ""
}

func (m *ActionRequest) GetIdSection() string {
	if m != nil {
		return m.IdSection
	}
	return ""
}

func (m *ActionRequest) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *ActionRequest) GetData() *Event {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ActionData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeOut              uint64   `protobuf:"varint,2,opt,name=time_out,json=timeOut,proto3" json:"time_out,omitempty"`
	EndPoints            []string `protobuf:"bytes,3,rep,name=end_points,json=endPoints,proto3" json:"end_points,omitempty"`
	EnableHook           bool     `protobuf:"varint,4,opt,name=enable_hook,json=enableHook,proto3" json:"enable_hook,omitempty"`
	Durable              bool     `protobuf:"varint,5,opt,name=durable,proto3" json:"durable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionData) Reset()         { *m = ActionData{} }
func (m *ActionData) String() string { return proto.CompactTextString(m) }
func (*ActionData) ProtoMessage()    {}
func (*ActionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ActionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionData.Unmarshal(m, b)
}
func (m *ActionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionData.Marshal(b, m, deterministic)
}
func (m *ActionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionData.Merge(m, src)
}
func (m *ActionData) XXX_Size() int {
	return xxx_messageInfo_ActionData.Size(m)
}
func (m *ActionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionData.DiscardUnknown(m)
}

var xxx_messageInfo_ActionData proto.InternalMessageInfo

func (m *ActionData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionData) GetTimeOut() uint64 {
	if m != nil {
		return m.TimeOut
	}
	return 0
}

func (m *ActionData) GetEndPoints() []string {
	if m != nil {
		return m.EndPoints
	}
	return nil
}

func (m *ActionData) GetEnableHook() bool {
	if m != nil {
		return m.EnableHook
	}
	return false
}

func (m *ActionData) GetDurable() bool {
	if m != nil {
		return m.Durable
	}
	return false
}

type RegisterAction struct {
	Type                 TypeAction `protobuf:"varint,1,opt,name=type,proto3,enum=proto.TypeAction" json:"type,omitempty"`
	Action               string     `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Priority             int32      `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Description          string     `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RegisterAction) Reset()         { *m = RegisterAction{} }
func (m *RegisterAction) String() string { return proto.CompactTextString(m) }
func (*RegisterAction) ProtoMessage()    {}
func (*RegisterAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RegisterAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAction.Unmarshal(m, b)
}
func (m *RegisterAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAction.Marshal(b, m, deterministic)
}
func (m *RegisterAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAction.Merge(m, src)
}
func (m *RegisterAction) XXX_Size() int {
	return xxx_messageInfo_RegisterAction.Size(m)
}
func (m *RegisterAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAction.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAction proto.InternalMessageInfo

func (m *RegisterAction) GetType() TypeAction {
	if m != nil {
		return m.Type
	}
	return TypeAction_Hook
}

func (m *RegisterAction) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RegisterAction) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *RegisterAction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RegisterRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Auth                 []string          `protobuf:"bytes,4,rep,name=auth,proto3" json:"auth,omitempty"`
	Url                  string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Token                string            `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Actions              []*ActionData     `protobuf:"bytes,7,rep,name=actions,proto3" json:"actions,omitempty"`
	Registers            []*RegisterAction `protobuf:"bytes,8,rep,name=registers,proto3" json:"registers,omitempty"`
	MetaData             map[string]string `protobuf:"bytes,9,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegisterRequest) GetAuth() []string {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *RegisterRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RegisterRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterRequest) GetActions() []*ActionData {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *RegisterRequest) GetRegisters() []*RegisterAction {
	if m != nil {
		return m.Registers
	}
	return nil
}

func (m *RegisterRequest) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type RegisterResponse struct {
	IdSection            string   `protobuf:"bytes,1,opt,name=id_section,json=idSection,proto3" json:"id_section,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetIdSection() string {
	if m != nil {
		return m.IdSection
	}
	return ""
}

type PingRequest struct {
	IdModule             string   `protobuf:"bytes,1,opt,name=id_module,json=idModule,proto3" json:"id_module,omitempty"`
	IdSection            string   `protobuf:"bytes,2,opt,name=id_section,json=idSection,proto3" json:"id_section,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetIdModule() string {
	if m != nil {
		return m.IdModule
	}
	return ""
}

func (m *PingRequest) GetIdSection() string {
	if m != nil {
		return m.IdSection
	}
	return ""
}

func (m *PingRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type HookStream struct {
	IdModule             string            `protobuf:"bytes,1,opt,name=id_module,json=idModule,proto3" json:"id_module,omitempty"`
	IdSection            string            `protobuf:"bytes,2,opt,name=id_section,json=idSection,proto3" json:"id_section,omitempty"`
	MetaData             map[string]string `protobuf:"bytes,3,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 *Event            `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HookStream) Reset()         { *m = HookStream{} }
func (m *HookStream) String() string { return proto.CompactTextString(m) }
func (*HookStream) ProtoMessage()    {}
func (*HookStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *HookStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookStream.Unmarshal(m, b)
}
func (m *HookStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookStream.Marshal(b, m, deterministic)
}
func (m *HookStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookStream.Merge(m, src)
}
func (m *HookStream) XXX_Size() int {
	return xxx_messageInfo_HookStream.Size(m)
}
func (m *HookStream) XXX_DiscardUnknown() {
	xxx_messageInfo_HookStream.DiscardUnknown(m)
}

var xxx_messageInfo_HookStream proto.InternalMessageInfo

func (m *HookStream) GetIdModule() string {
	if m != nil {
		return m.IdModule
	}
	return ""
}

func (m *HookStream) GetIdSection() string {
	if m != nil {
		return m.IdSection
	}
	return ""
}

func (m *HookStream) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *HookStream) GetData() *Event {
	if m != nil {
		return m.Data
	}
	return nil
}

type Event struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32             `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Time                 uint64            `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Header               map[string]string `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 []byte            `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Event) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Event) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EventStreamRequest struct {
	IdModule             string            `protobuf:"bytes,1,opt,name=id_module,json=idModule,proto3" json:"id_module,omitempty"`
	IdSection            string            `protobuf:"bytes,2,opt,name=id_section,json=idSection,proto3" json:"id_section,omitempty"`
	MetaData             map[string]string `protobuf:"bytes,3,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EventStreamRequest) Reset()         { *m = EventStreamRequest{} }
func (m *EventStreamRequest) String() string { return proto.CompactTextString(m) }
func (*EventStreamRequest) ProtoMessage()    {}
func (*EventStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *EventStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStreamRequest.Unmarshal(m, b)
}
func (m *EventStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStreamRequest.Marshal(b, m, deterministic)
}
func (m *EventStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStreamRequest.Merge(m, src)
}
func (m *EventStreamRequest) XXX_Size() int {
	return xxx_messageInfo_EventStreamRequest.Size(m)
}
func (m *EventStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventStreamRequest proto.InternalMessageInfo

func (m *EventStreamRequest) GetIdModule() string {
	if m != nil {
		return m.IdModule
	}
	return ""
}

func (m *EventStreamRequest) GetIdSection() string {
	if m != nil {
		return m.IdSection
	}
	return ""
}

func (m *EventStreamRequest) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.TypeAction", TypeAction_name, TypeAction_value)
	proto.RegisterType((*ActionRequest)(nil), "proto.ActionRequest")
	proto.RegisterMapType((map[string]string)(nil), "proto.ActionRequest.MetaDataEntry")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*ActionData)(nil), "proto.ActionData")
	proto.RegisterType((*RegisterAction)(nil), "proto.RegisterAction")
	proto.RegisterType((*RegisterRequest)(nil), "proto.RegisterRequest")
	proto.RegisterMapType((map[string]string)(nil), "proto.RegisterRequest.MetaDataEntry")
	proto.RegisterType((*RegisterResponse)(nil), "proto.RegisterResponse")
	proto.RegisterType((*PingRequest)(nil), "proto.PingRequest")
	proto.RegisterType((*HookStream)(nil), "proto.HookStream")
	proto.RegisterMapType((map[string]string)(nil), "proto.HookStream.MetaDataEntry")
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterMapType((map[string]string)(nil), "proto.Event.HeaderEntry")
	proto.RegisterType((*EventStreamRequest)(nil), "proto.EventStreamRequest")
	proto.RegisterMapType((map[string]string)(nil), "proto.EventStreamRequest.MetaDataEntry")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x0e, 0x6d, 0xd9, 0x96, 0x46, 0xf9, 0x71, 0x06, 0xd9, 0xac, 0xe2, 0xc5, 0x22, 0x86, 0x90,
	0xc5, 0x1a, 0xbb, 0x80, 0x9b, 0x38, 0x97, 0x36, 0x6d, 0x51, 0x04, 0x4d, 0x80, 0x5c, 0x82, 0x06,
	0x4a, 0xef, 0x86, 0x62, 0x11, 0x09, 0x61, 0x5b, 0x72, 0x25, 0xca, 0x80, 0x1f, 0xa2, 0xc7, 0xde,
	0xfa, 0x56, 0xed, 0x2b, 0xf4, 0x54, 0xf4, 0x1d, 0x0a, 0x0e, 0x69, 0x5b, 0x8e, 0x53, 0xb4, 0x45,
	0x52, 0xf4, 0x24, 0xce, 0x70, 0x86, 0xf3, 0xcd, 0xc7, 0x6f, 0x28, 0x58, 0xcb, 0x78, 0x3a, 0x16,
	0x3d, 0xde, 0x1e, 0xa5, 0x89, 0x4c, 0xb0, 0x42, 0x1f, 0xff, 0x33, 0x83, 0xb5, 0xe3, 0x9e, 0x14,
	0x49, 0x1c, 0xf0, 0x37, 0x39, 0xcf, 0x24, 0xfe, 0x05, 0x8e, 0x88, 0xba, 0xc3, 0x24, 0xca, 0x07,
	0xdc, 0x63, 0x4d, 0xd6, 0x72, 0x02, 0x5b, 0x44, 0xe7, 0x64, 0xe3, 0xdf, 0x00, 0x22, 0xea, 0x66,
	0x9c, 0x32, 0xbc, 0x12, 0xed, 0x3a, 0x22, 0xba, 0xd4, 0x0e, 0x7c, 0x01, 0xce, 0x90, 0xcb, 0xb0,
	0x1b, 0x85, 0x32, 0xf4, 0xca, 0xcd, 0x72, 0xcb, 0xed, 0xf8, 0xba, 0x5e, 0x7b, 0xa1, 0x48, 0xfb,
	0x9c, 0xcb, 0xf0, 0x24, 0x94, 0xe1, 0x69, 0x2c, 0xd3, 0x49, 0x60, 0x0f, 0x8d, 0x89, 0x4d, 0xb0,
	0x28, 0xd7, 0x6a, 0xb2, 0x96, 0xdb, 0x59, 0x35, 0xb9, 0xa7, 0x63, 0x1e, 0xcb, 0x80, 0x76, 0x1a,
	0x4f, 0x61, 0x6d, 0x21, 0x19, 0xeb, 0x50, 0xee, 0xf3, 0x89, 0x41, 0xaa, 0x96, 0xb8, 0x05, 0x95,
	0x71, 0x38, 0xc8, 0xb9, 0xc1, 0xa7, 0x8d, 0xa3, 0xd2, 0x63, 0xe6, 0xef, 0x81, 0x1d, 0xf0, 0x6c,
	0x94, 0xc4, 0x19, 0x47, 0x0f, 0x6a, 0x59, 0xde, 0xeb, 0xf1, 0x2c, 0xa3, 0x5c, 0x3b, 0x98, 0x9a,
	0xfe, 0x3b, 0x06, 0xa0, 0xe1, 0x12, 0x26, 0x04, 0x2b, 0x0e, 0x87, 0x53, 0x2e, 0x68, 0x8d, 0x3b,
	0x60, 0x4b, 0x31, 0xe4, 0xdd, 0x24, 0x97, 0x54, 0xc5, 0x0a, 0x6a, 0xca, 0x7e, 0x95, 0x4b, 0x45,
	0x11, 0x8f, 0xa3, 0xee, 0x28, 0x11, 0xb1, 0xcc, 0x88, 0x04, 0x27, 0x70, 0x78, 0x1c, 0x5d, 0x90,
	0x03, 0x77, 0xc1, 0xe5, 0x71, 0x78, 0x35, 0xe0, 0xdd, 0x9b, 0x24, 0xe9, 0x53, 0xa3, 0x76, 0x00,
	0xda, 0x75, 0x96, 0x24, 0x7d, 0x85, 0x2b, 0xca, 0x53, 0x65, 0x7a, 0x15, 0x8d, 0xcb, 0x98, 0xfe,
	0x5b, 0x06, 0xeb, 0x01, 0xbf, 0x16, 0x99, 0xe4, 0xa9, 0xc6, 0x87, 0xff, 0x80, 0x25, 0x27, 0x23,
	0x8d, 0x6d, 0xbd, 0xb3, 0x69, 0xf8, 0x7a, 0x3d, 0x19, 0x71, 0xc3, 0x37, 0x6d, 0xe3, 0x36, 0x54,
	0xc3, 0xe2, 0x95, 0x19, 0x0b, 0x1b, 0x60, 0x8f, 0x52, 0x91, 0xa4, 0x42, 0x4e, 0xbc, 0x72, 0x93,
	0xb5, 0x2a, 0xc1, 0xcc, 0xc6, 0x26, 0xb8, 0x27, 0x3c, 0xeb, 0xa5, 0x62, 0x44, 0x89, 0x16, 0x25,
	0x16, 0x5d, 0xfe, 0x97, 0x12, 0x6c, 0x4c, 0xf1, 0x4c, 0xd5, 0xb3, 0x0e, 0x25, 0x11, 0x19, 0xaa,
	0x4a, 0x22, 0x9a, 0x91, 0x57, 0x2a, 0x90, 0xe7, 0x41, 0x6d, 0xcc, 0xd3, 0x4c, 0x9d, 0x5a, 0x26,
	0xf7, 0xd4, 0x54, 0xd1, 0x61, 0x2e, 0x6f, 0x3c, 0x8b, 0x58, 0xa3, 0xb5, 0xba, 0xdf, 0x3c, 0x1d,
	0x10, 0x17, 0x4e, 0xa0, 0x96, 0xea, 0x7e, 0x65, 0xd2, 0xe7, 0xb1, 0x57, 0xd5, 0xf7, 0x4b, 0x06,
	0xfe, 0x0f, 0x35, 0xdd, 0x55, 0xe6, 0xd5, 0x48, 0x79, 0x9b, 0x0b, 0xca, 0x53, 0x57, 0x19, 0x4c,
	0x23, 0xf0, 0x10, 0x9c, 0xd4, 0x20, 0xcf, 0x3c, 0x9b, 0xc2, 0xff, 0x30, 0xe1, 0x8b, 0x0c, 0x07,
	0xf3, 0x38, 0x3c, 0x2e, 0xaa, 0xdb, 0xa1, 0xa4, 0xbd, 0x5b, 0x49, 0xdf, 0xd1, 0xf7, 0xfd, 0xd4,
	0x7b, 0x00, 0xf5, 0x79, 0x1d, 0xa3, 0xe2, 0xc5, 0x81, 0x64, 0xb7, 0x06, 0xd2, 0xef, 0x82, 0x7b,
	0x21, 0xe2, 0xeb, 0x87, 0x98, 0xed, 0x19, 0xeb, 0xe5, 0x02, 0xeb, 0xfe, 0x27, 0x06, 0xa0, 0x64,
	0x7b, 0x29, 0x53, 0x1e, 0x0e, 0xef, 0x55, 0xe0, 0xd9, 0xf2, 0xe3, 0xb1, 0x6b, 0xe8, 0x9d, 0x57,
	0xf8, 0x5d, 0x2f, 0xc7, 0x47, 0x06, 0x15, 0x3a, 0xec, 0x87, 0x14, 0x8e, 0x66, 0x2c, 0xf5, 0x4c,
	0xe9, 0x19, 0x54, 0x3e, 0x31, 0xe4, 0x04, 0xd0, 0x0a, 0x68, 0x8d, 0xfb, 0x50, 0xbd, 0xe1, 0x61,
	0xc4, 0x53, 0xaf, 0x42, 0xfd, 0x7a, 0x45, 0xd8, 0xed, 0x33, 0xda, 0xd2, 0x8d, 0x9a, 0x38, 0x75,
	0x0a, 0xb5, 0xa9, 0xa4, 0xbf, 0x6a, 0x1a, 0x7b, 0x02, 0x6e, 0x21, 0xf4, 0xa7, 0xda, 0xfa, 0xc0,
	0x00, 0xa9, 0x98, 0x66, 0xf7, 0x21, 0x74, 0x72, 0xb2, 0x7c, 0x8d, 0xff, 0x16, 0xdb, 0x5a, 0xa8,
	0xf4, 0x4b, 0x06, 0xe5, 0xbf, 0x3d, 0x80, 0xf9, 0x13, 0x88, 0x36, 0x58, 0x4a, 0x3f, 0xf5, 0x15,
	0x5c, 0x05, 0xfb, 0xd4, 0x3c, 0xc4, 0x75, 0xd6, 0x79, 0x5f, 0x02, 0xf7, 0x65, 0x92, 0xf2, 0x4b,
	0xfd, 0x5f, 0xc4, 0x23, 0x70, 0x0b, 0x00, 0x71, 0xe7, 0x9b, 0xa0, 0x1b, 0x0b, 0xea, 0xf2, 0x57,
	0xf6, 0x19, 0xb6, 0xa1, 0xa2, 0x6a, 0x64, 0xb8, 0xb9, 0xa4, 0xd8, 0xdb, 0xd1, 0x2d, 0xb6, 0xcf,
	0xf0, 0x11, 0x58, 0x6a, 0x2e, 0x11, 0xcd, 0x5e, 0x61, 0x48, 0x1b, 0x1b, 0xb3, 0x37, 0x45, 0xcf,
	0xb8, 0x4a, 0xc1, 0xe7, 0xea, 0xcf, 0xa5, 0x67, 0x1f, 0xb7, 0xef, 0x7e, 0x74, 0x1a, 0x7f, 0x2e,
	0xf9, 0xa7, 0x07, 0xe0, 0x01, 0x54, 0x0d, 0x1b, 0x5b, 0x77, 0xfd, 0x8f, 0xef, 0xa8, 0x79, 0x55,
	0x25, 0xcf, 0xe1, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xc4, 0x73, 0x95, 0x38, 0x08, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	EventStream(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (CoreService_EventStreamClient, error)
	Hooks(ctx context.Context, opts ...grpc.CallOption) (CoreService_HooksClient, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (CoreService_PingClient, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Response, error)
}

type coreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreServiceClient(cc grpc.ClientConnInterface) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) EventStream(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (CoreService_EventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CoreService_serviceDesc.Streams[0], "/proto.CoreService/EventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServiceEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoreService_EventStreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type coreServiceEventStreamClient struct {
	grpc.ClientStream
}

func (x *coreServiceEventStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) Hooks(ctx context.Context, opts ...grpc.CallOption) (CoreService_HooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CoreService_serviceDesc.Streams[1], "/proto.CoreService/Hooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServiceHooksClient{stream}
	return x, nil
}

type CoreService_HooksClient interface {
	Send(*HookStream) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type coreServiceHooksClient struct {
	grpc.ClientStream
}

func (x *coreServiceHooksClient) Send(m *HookStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreServiceHooksClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) Ping(ctx context.Context, opts ...grpc.CallOption) (CoreService_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CoreService_serviceDesc.Streams[2], "/proto.CoreService/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServicePingClient{stream}
	return x, nil
}

type CoreService_PingClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type coreServicePingClient struct {
	grpc.ClientStream
}

func (x *coreServicePingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreServicePingClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.CoreService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CoreService/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	EventStream(*EventStreamRequest, CoreService_EventStreamServer) error
	Hooks(CoreService_HooksServer) error
	Ping(CoreService_PingServer) error
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Action(context.Context, *ActionRequest) (*Response, error)
}

// UnimplementedCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (*UnimplementedCoreServiceServer) EventStream(req *EventStreamRequest, srv CoreService_EventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventStream not implemented")
}
func (*UnimplementedCoreServiceServer) Hooks(srv CoreService_HooksServer) error {
	return status.Errorf(codes.Unimplemented, "method Hooks not implemented")
}
func (*UnimplementedCoreServiceServer) Ping(srv CoreService_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedCoreServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedCoreServiceServer) Action(ctx context.Context, req *ActionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_EventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServiceServer).EventStream(m, &coreServiceEventStreamServer{stream})
}

type CoreService_EventStreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type coreServiceEventStreamServer struct {
	grpc.ServerStream
}

func (x *coreServiceEventStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _CoreService_Hooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServiceServer).Hooks(&coreServiceHooksServer{stream})
}

type CoreService_HooksServer interface {
	Send(*Event) error
	Recv() (*HookStream, error)
	grpc.ServerStream
}

type coreServiceHooksServer struct {
	grpc.ServerStream
}

func (x *coreServiceHooksServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreServiceHooksServer) Recv() (*HookStream, error) {
	m := new(HookStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreService_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServiceServer).Ping(&coreServicePingServer{stream})
}

type CoreService_PingServer interface {
	SendAndClose(*Response) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type coreServicePingServer struct {
	grpc.ServerStream
}

func (x *coreServicePingServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreServicePingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CoreService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CoreService/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CoreService_Register_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _CoreService_Action_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventStream",
			Handler:       _CoreService_EventStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Hooks",
			Handler:       _CoreService_Hooks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _CoreService_Ping_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
