// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/transport.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Connection struct {
	LocalAddress         *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        *core.Address `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{0}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Connection) GetRemoteAddress() *core.Address {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

type SocketEvent struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to EventSelector:
	//	*SocketEvent_Read_
	//	*SocketEvent_Write_
	//	*SocketEvent_Closed_
	EventSelector        isSocketEvent_EventSelector `protobuf_oneof:"event_selector"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SocketEvent) Reset()         { *m = SocketEvent{} }
func (m *SocketEvent) String() string { return proto.CompactTextString(m) }
func (*SocketEvent) ProtoMessage()    {}
func (*SocketEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1}
}

func (m *SocketEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent.Unmarshal(m, b)
}
func (m *SocketEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent.Marshal(b, m, deterministic)
}
func (m *SocketEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent.Merge(m, src)
}
func (m *SocketEvent) XXX_Size() int {
	return xxx_messageInfo_SocketEvent.Size(m)
}
func (m *SocketEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent proto.InternalMessageInfo

func (m *SocketEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isSocketEvent_EventSelector interface {
	isSocketEvent_EventSelector()
}

type SocketEvent_Read_ struct {
	Read *SocketEvent_Read `protobuf:"bytes,2,opt,name=read,proto3,oneof"`
}

type SocketEvent_Write_ struct {
	Write *SocketEvent_Write `protobuf:"bytes,3,opt,name=write,proto3,oneof"`
}

type SocketEvent_Closed_ struct {
	Closed *SocketEvent_Closed `protobuf:"bytes,4,opt,name=closed,proto3,oneof"`
}

func (*SocketEvent_Read_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Write_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Closed_) isSocketEvent_EventSelector() {}

func (m *SocketEvent) GetEventSelector() isSocketEvent_EventSelector {
	if m != nil {
		return m.EventSelector
	}
	return nil
}

func (m *SocketEvent) GetRead() *SocketEvent_Read {
	if x, ok := m.GetEventSelector().(*SocketEvent_Read_); ok {
		return x.Read
	}
	return nil
}

func (m *SocketEvent) GetWrite() *SocketEvent_Write {
	if x, ok := m.GetEventSelector().(*SocketEvent_Write_); ok {
		return x.Write
	}
	return nil
}

func (m *SocketEvent) GetClosed() *SocketEvent_Closed {
	if x, ok := m.GetEventSelector().(*SocketEvent_Closed_); ok {
		return x.Closed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketEvent_Read_)(nil),
		(*SocketEvent_Write_)(nil),
		(*SocketEvent_Closed_)(nil),
	}
}

type SocketEvent_Read struct {
	Data                 *Body    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Read) Reset()         { *m = SocketEvent_Read{} }
func (m *SocketEvent_Read) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Read) ProtoMessage()    {}
func (*SocketEvent_Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 0}
}

func (m *SocketEvent_Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Read.Unmarshal(m, b)
}
func (m *SocketEvent_Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Read.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Read.Merge(m, src)
}
func (m *SocketEvent_Read) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Read.Size(m)
}
func (m *SocketEvent_Read) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Read.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Read proto.InternalMessageInfo

func (m *SocketEvent_Read) GetData() *Body {
	if m != nil {
		return m.Data
	}
	return nil
}

type SocketEvent_Write struct {
	Data                 *Body    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	EndStream            bool     `protobuf:"varint,2,opt,name=end_stream,json=endStream,proto3" json:"end_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Write) Reset()         { *m = SocketEvent_Write{} }
func (m *SocketEvent_Write) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Write) ProtoMessage()    {}
func (*SocketEvent_Write) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 1}
}

func (m *SocketEvent_Write) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Write.Unmarshal(m, b)
}
func (m *SocketEvent_Write) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Write.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Write) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Write.Merge(m, src)
}
func (m *SocketEvent_Write) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Write.Size(m)
}
func (m *SocketEvent_Write) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Write.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Write proto.InternalMessageInfo

func (m *SocketEvent_Write) GetData() *Body {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SocketEvent_Write) GetEndStream() bool {
	if m != nil {
		return m.EndStream
	}
	return false
}

type SocketEvent_Closed struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketEvent_Closed) Reset()         { *m = SocketEvent_Closed{} }
func (m *SocketEvent_Closed) String() string { return proto.CompactTextString(m) }
func (*SocketEvent_Closed) ProtoMessage()    {}
func (*SocketEvent_Closed) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{1, 2}
}

func (m *SocketEvent_Closed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketEvent_Closed.Unmarshal(m, b)
}
func (m *SocketEvent_Closed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketEvent_Closed.Marshal(b, m, deterministic)
}
func (m *SocketEvent_Closed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketEvent_Closed.Merge(m, src)
}
func (m *SocketEvent_Closed) XXX_Size() int {
	return xxx_messageInfo_SocketEvent_Closed.Size(m)
}
func (m *SocketEvent_Closed) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketEvent_Closed.DiscardUnknown(m)
}

var xxx_messageInfo_SocketEvent_Closed proto.InternalMessageInfo

type SocketBufferedTrace struct {
	TraceId              uint64         `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Connection           *Connection    `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	Events               []*SocketEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	ReadTruncated        bool           `protobuf:"varint,4,opt,name=read_truncated,json=readTruncated,proto3" json:"read_truncated,omitempty"`
	WriteTruncated       bool           `protobuf:"varint,5,opt,name=write_truncated,json=writeTruncated,proto3" json:"write_truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SocketBufferedTrace) Reset()         { *m = SocketBufferedTrace{} }
func (m *SocketBufferedTrace) String() string { return proto.CompactTextString(m) }
func (*SocketBufferedTrace) ProtoMessage()    {}
func (*SocketBufferedTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{2}
}

func (m *SocketBufferedTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketBufferedTrace.Unmarshal(m, b)
}
func (m *SocketBufferedTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketBufferedTrace.Marshal(b, m, deterministic)
}
func (m *SocketBufferedTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketBufferedTrace.Merge(m, src)
}
func (m *SocketBufferedTrace) XXX_Size() int {
	return xxx_messageInfo_SocketBufferedTrace.Size(m)
}
func (m *SocketBufferedTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketBufferedTrace.DiscardUnknown(m)
}

var xxx_messageInfo_SocketBufferedTrace proto.InternalMessageInfo

func (m *SocketBufferedTrace) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *SocketBufferedTrace) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *SocketBufferedTrace) GetEvents() []*SocketEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *SocketBufferedTrace) GetReadTruncated() bool {
	if m != nil {
		return m.ReadTruncated
	}
	return false
}

func (m *SocketBufferedTrace) GetWriteTruncated() bool {
	if m != nil {
		return m.WriteTruncated
	}
	return false
}

type SocketStreamedTraceSegment struct {
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are valid to be assigned to MessagePiece:
	//	*SocketStreamedTraceSegment_Connection
	//	*SocketStreamedTraceSegment_Event
	MessagePiece         isSocketStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *SocketStreamedTraceSegment) Reset()         { *m = SocketStreamedTraceSegment{} }
func (m *SocketStreamedTraceSegment) String() string { return proto.CompactTextString(m) }
func (*SocketStreamedTraceSegment) ProtoMessage()    {}
func (*SocketStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a9cebdb27ee552, []int{3}
}

func (m *SocketStreamedTraceSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketStreamedTraceSegment.Unmarshal(m, b)
}
func (m *SocketStreamedTraceSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketStreamedTraceSegment.Marshal(b, m, deterministic)
}
func (m *SocketStreamedTraceSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketStreamedTraceSegment.Merge(m, src)
}
func (m *SocketStreamedTraceSegment) XXX_Size() int {
	return xxx_messageInfo_SocketStreamedTraceSegment.Size(m)
}
func (m *SocketStreamedTraceSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketStreamedTraceSegment.DiscardUnknown(m)
}

var xxx_messageInfo_SocketStreamedTraceSegment proto.InternalMessageInfo

func (m *SocketStreamedTraceSegment) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

type isSocketStreamedTraceSegment_MessagePiece interface {
	isSocketStreamedTraceSegment_MessagePiece()
}

type SocketStreamedTraceSegment_Connection struct {
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3,oneof"`
}

type SocketStreamedTraceSegment_Event struct {
	Event *SocketEvent `protobuf:"bytes,3,opt,name=event,proto3,oneof"`
}

func (*SocketStreamedTraceSegment_Connection) isSocketStreamedTraceSegment_MessagePiece() {}

func (*SocketStreamedTraceSegment_Event) isSocketStreamedTraceSegment_MessagePiece() {}

func (m *SocketStreamedTraceSegment) GetMessagePiece() isSocketStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (m *SocketStreamedTraceSegment) GetConnection() *Connection {
	if x, ok := m.GetMessagePiece().(*SocketStreamedTraceSegment_Connection); ok {
		return x.Connection
	}
	return nil
}

func (m *SocketStreamedTraceSegment) GetEvent() *SocketEvent {
	if x, ok := m.GetMessagePiece().(*SocketStreamedTraceSegment_Event); ok {
		return x.Event
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketStreamedTraceSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketStreamedTraceSegment_Connection)(nil),
		(*SocketStreamedTraceSegment_Event)(nil),
	}
}

func init() {
	proto.RegisterType((*Connection)(nil), "envoy.data.tap.v2alpha.Connection")
	proto.RegisterType((*SocketEvent)(nil), "envoy.data.tap.v2alpha.SocketEvent")
	proto.RegisterType((*SocketEvent_Read)(nil), "envoy.data.tap.v2alpha.SocketEvent.Read")
	proto.RegisterType((*SocketEvent_Write)(nil), "envoy.data.tap.v2alpha.SocketEvent.Write")
	proto.RegisterType((*SocketEvent_Closed)(nil), "envoy.data.tap.v2alpha.SocketEvent.Closed")
	proto.RegisterType((*SocketBufferedTrace)(nil), "envoy.data.tap.v2alpha.SocketBufferedTrace")
	proto.RegisterType((*SocketStreamedTraceSegment)(nil), "envoy.data.tap.v2alpha.SocketStreamedTraceSegment")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/transport.proto", fileDescriptor_03a9cebdb27ee552)
}

var fileDescriptor_03a9cebdb27ee552 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0x85, 0xe5, 0x43, 0xf6, 0x22, 0xac, 0x19, 0x13, 0x83, 0x8d, 0x66, 0x0d, 0xeb, 0x07, 0xfa,
	0x30, 0x35, 0xf8, 0x42, 0x62, 0xa2, 0xa1, 0xbb, 0x26, 0xf8, 0xb6, 0x29, 0x24, 0xfa, 0xd6, 0xcc,
	0x76, 0x2e, 0xd8, 0xd8, 0x76, 0x9a, 0xe9, 0x80, 0xf2, 0xe8, 0x3f, 0xf0, 0x27, 0xf9, 0xee, 0x9f,
	0x32, 0xf3, 0x01, 0x4b, 0x8c, 0x28, 0xee, 0x1b, 0x33, 0x39, 0xe7, 0xdc, 0x73, 0xcf, 0x9c, 0x02,
	0x4f, 0x31, 0x5f, 0x89, 0xb5, 0xcf, 0x99, 0x62, 0xbe, 0x62, 0x85, 0xbf, 0x1a, 0xb2, 0xb4, 0xf8,
	0xc4, 0x7c, 0x25, 0x59, 0x5e, 0x16, 0x42, 0x2a, 0x5a, 0x48, 0xa1, 0x04, 0xb9, 0x67, 0x70, 0x54,
	0xe3, 0xa8, 0x62, 0x05, 0x75, 0x38, 0xef, 0xd4, 0xf2, 0x59, 0x91, 0xf8, 0xab, 0xa1, 0x1f, 0x0b,
	0x89, 0x3e, 0xe3, 0x5c, 0x62, 0x59, 0x5a, 0xa2, 0x77, 0xb6, 0x67, 0x40, 0x2c, 0xb2, 0x4c, 0xe4,
	0x0e, 0x74, 0xba, 0x10, 0x62, 0x91, 0xa2, 0x6f, 0x4e, 0x57, 0xcb, 0xb9, 0xaf, 0x92, 0x0c, 0x4b,
	0xc5, 0xb2, 0xc2, 0x02, 0xfa, 0xdf, 0xab, 0x00, 0xe7, 0x22, 0xcf, 0x31, 0x56, 0x89, 0xc8, 0xc9,
	0x5b, 0xe8, 0xa4, 0x22, 0x66, 0x69, 0xe4, 0x66, 0xf5, 0x8e, 0x1e, 0x55, 0x07, 0xed, 0xa1, 0x47,
	0xad, 0x4b, 0x56, 0x24, 0x74, 0x35, 0xa4, 0xda, 0x0d, 0x1d, 0x5b, 0x44, 0x78, 0xdb, 0x10, 0xdc,
	0x89, 0x8c, 0xa1, 0x2b, 0x31, 0x13, 0x0a, 0xb7, 0x0a, 0xb5, 0x7f, 0x2a, 0x74, 0x2c, 0xc3, 0x1d,
	0xfb, 0x3f, 0x6a, 0xd0, 0x9e, 0x8a, 0xf8, 0x33, 0xaa, 0x77, 0x2b, 0xcc, 0x15, 0x19, 0xc1, 0xf1,
	0xd6, 0x75, 0xaf, 0xea, 0xd4, 0xec, 0x5e, 0x74, 0xb3, 0x17, 0x9d, 0x6d, 0x10, 0xe1, 0x35, 0x98,
	0xbc, 0x81, 0xba, 0x44, 0xc6, 0xdd, 0x12, 0x03, 0xfa, 0xe7, 0xa8, 0xe9, 0xce, 0x30, 0x1a, 0x22,
	0xe3, 0x93, 0x4a, 0x68, 0x78, 0x64, 0x0c, 0x8d, 0x2f, 0x32, 0x51, 0xe8, 0x76, 0x78, 0x7e, 0x88,
	0xc0, 0x07, 0x4d, 0x98, 0x54, 0x42, 0xcb, 0x24, 0x17, 0xd0, 0x8c, 0x53, 0x51, 0x22, 0xef, 0xd5,
	0x8d, 0xc6, 0x8b, 0x43, 0x34, 0xce, 0x0d, 0x63, 0x52, 0x09, 0x1d, 0xd7, 0x1b, 0x41, 0x5d, 0x1b,
	0x23, 0x2f, 0xa1, 0xae, 0x89, 0x2e, 0x85, 0x07, 0xfb, 0xb4, 0x02, 0xc1, 0xd7, 0xa1, 0x41, 0x7a,
	0x1f, 0xa1, 0x61, 0x1c, 0xfd, 0x3f, 0x95, 0x3c, 0x04, 0xc0, 0x9c, 0x47, 0xa5, 0x92, 0xc8, 0x32,
	0x93, 0x61, 0x2b, 0x3c, 0xc6, 0x9c, 0x4f, 0xcd, 0x85, 0xd7, 0x82, 0xa6, 0xf5, 0x19, 0xdc, 0x81,
	0x2e, 0x6a, 0xdf, 0x51, 0x89, 0x29, 0xc6, 0x4a, 0xc8, 0xfe, 0xb7, 0x23, 0xb8, 0x6b, 0x17, 0x0a,
	0x96, 0xf3, 0x39, 0x4a, 0xe4, 0x33, 0xc9, 0x62, 0x24, 0xf7, 0xa1, 0xa5, 0xf4, 0x8f, 0x28, 0xe1,
	0xc6, 0x48, 0x3d, 0xbc, 0x65, 0xce, 0xef, 0x39, 0x09, 0x00, 0xe2, 0x6d, 0x0f, 0xdd, 0x8b, 0xf5,
	0xf7, 0xb9, 0xbc, 0x6e, 0x6c, 0xb8, 0xc3, 0x22, 0xaf, 0xa1, 0x69, 0x8c, 0xe8, 0xd2, 0xd5, 0x06,
	0xed, 0xe1, 0xd9, 0x01, 0x61, 0x87, 0x8e, 0x42, 0x9e, 0xe8, 0xe6, 0x32, 0x1e, 0x29, 0xb9, 0xcc,
	0x63, 0xa6, 0xdc, 0x8b, 0xb5, 0x74, 0x3b, 0x19, 0x9f, 0x6d, 0x2e, 0xc9, 0x33, 0x38, 0x31, 0x2f,
	0xbb, 0x83, 0x6b, 0x18, 0x5c, 0xd7, 0x5c, 0x6f, 0x81, 0xfd, 0x9f, 0x55, 0xf0, 0xec, 0x1c, 0x1b,
	0x98, 0xcb, 0x60, 0x8a, 0x8b, 0x4c, 0xb7, 0xfa, 0x2f, 0x51, 0x5c, 0xdc, 0x2c, 0x8a, 0x49, 0xe5,
	0xb7, 0x30, 0x1a, 0x66, 0x33, 0x57, 0xde, 0x43, 0xb2, 0xd0, 0xb5, 0x35, 0x9c, 0xe0, 0x04, 0x3a,
	0x19, 0x96, 0x25, 0x5b, 0x60, 0x54, 0x24, 0x18, 0x63, 0x30, 0x82, 0xc7, 0x89, 0xb0, 0x12, 0x85,
	0x14, 0x5f, 0xd7, 0x7b, 0xd4, 0x82, 0xee, 0x6c, 0xf3, 0xff, 0x76, 0xa9, 0x3f, 0xcd, 0xcb, 0xea,
	0x55, 0xd3, 0x7c, 0xa3, 0xaf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x0d, 0xac, 0x8e, 0x11,
	0x05, 0x00, 0x00,
}
