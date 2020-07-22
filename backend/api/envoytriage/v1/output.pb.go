// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoytriage/v1/output.proto

package envoytriagev1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type HostStatus struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Healthy              bool     `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{0}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *HostStatus) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

type ClusterStatus struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HostStatuses         []*HostStatus `protobuf:"bytes,2,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

type Clusters struct {
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{2}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

type ConfigDump struct {
	Value                *_struct.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigDump) Reset()         { *m = ConfigDump{} }
func (m *ConfigDump) String() string { return proto.CompactTextString(m) }
func (*ConfigDump) ProtoMessage()    {}
func (*ConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{3}
}

func (m *ConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigDump.Unmarshal(m, b)
}
func (m *ConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigDump.Marshal(b, m, deterministic)
}
func (m *ConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigDump.Merge(m, src)
}
func (m *ConfigDump) XXX_Size() int {
	return xxx_messageInfo_ConfigDump.Size(m)
}
func (m *ConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigDump proto.InternalMessageInfo

func (m *ConfigDump) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type ListenerStatus struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LocalAddress         string   `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{4}
}

func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerStatus.Unmarshal(m, b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return xxx_messageInfo_ListenerStatus.Size(m)
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

type Listeners struct {
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{5}
}

func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listeners.Unmarshal(m, b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return xxx_messageInfo_Listeners.Size(m)
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

type Runtime struct {
	Entries              []*Runtime_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{6}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetEntries() []*Runtime_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Runtime_Entry struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Runtime_Entry_Value
	Type                 isRuntime_Entry_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Runtime_Entry) Reset()         { *m = Runtime_Entry{} }
func (m *Runtime_Entry) String() string { return proto.CompactTextString(m) }
func (*Runtime_Entry) ProtoMessage()    {}
func (*Runtime_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{6, 0}
}

func (m *Runtime_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime_Entry.Unmarshal(m, b)
}
func (m *Runtime_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime_Entry.Marshal(b, m, deterministic)
}
func (m *Runtime_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime_Entry.Merge(m, src)
}
func (m *Runtime_Entry) XXX_Size() int {
	return xxx_messageInfo_Runtime_Entry.Size(m)
}
func (m *Runtime_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime_Entry proto.InternalMessageInfo

func (m *Runtime_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isRuntime_Entry_Type interface {
	isRuntime_Entry_Type()
}

type Runtime_Entry_Value struct {
	Value string `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}

func (*Runtime_Entry_Value) isRuntime_Entry_Type() {}

func (m *Runtime_Entry) GetType() isRuntime_Entry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Runtime_Entry) GetValue() string {
	if x, ok := m.GetType().(*Runtime_Entry_Value); ok {
		return x.Value
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Runtime_Entry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Runtime_Entry_Value)(nil),
	}
}

type ServerInfo struct {
	Value                *_struct.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{7}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Stats struct {
	// Counters and gauges are returned here.
	Stats                []*Stats_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{8}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetStats() []*Stats_Stat {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Stats_Stat struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                uint64   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats_Stat) Reset()         { *m = Stats_Stat{} }
func (m *Stats_Stat) String() string { return proto.CompactTextString(m) }
func (*Stats_Stat) ProtoMessage()    {}
func (*Stats_Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1adb45950dad7055, []int{8, 0}
}

func (m *Stats_Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats_Stat.Unmarshal(m, b)
}
func (m *Stats_Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats_Stat.Marshal(b, m, deterministic)
}
func (m *Stats_Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats_Stat.Merge(m, src)
}
func (m *Stats_Stat) XXX_Size() int {
	return xxx_messageInfo_Stats_Stat.Size(m)
}
func (m *Stats_Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stats_Stat proto.InternalMessageInfo

func (m *Stats_Stat) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Stats_Stat) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*HostStatus)(nil), "clutch.envoytriage.v1.HostStatus")
	proto.RegisterType((*ClusterStatus)(nil), "clutch.envoytriage.v1.ClusterStatus")
	proto.RegisterType((*Clusters)(nil), "clutch.envoytriage.v1.Clusters")
	proto.RegisterType((*ConfigDump)(nil), "clutch.envoytriage.v1.ConfigDump")
	proto.RegisterType((*ListenerStatus)(nil), "clutch.envoytriage.v1.ListenerStatus")
	proto.RegisterType((*Listeners)(nil), "clutch.envoytriage.v1.Listeners")
	proto.RegisterType((*Runtime)(nil), "clutch.envoytriage.v1.Runtime")
	proto.RegisterType((*Runtime_Entry)(nil), "clutch.envoytriage.v1.Runtime.Entry")
	proto.RegisterType((*ServerInfo)(nil), "clutch.envoytriage.v1.ServerInfo")
	proto.RegisterType((*Stats)(nil), "clutch.envoytriage.v1.Stats")
	proto.RegisterType((*Stats_Stat)(nil), "clutch.envoytriage.v1.Stats.Stat")
}

func init() {
	proto.RegisterFile("envoytriage/v1/output.proto", fileDescriptor_1adb45950dad7055)
}

var fileDescriptor_1adb45950dad7055 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x76, 0xd3, 0xa4, 0x69, 0x5f, 0x1b, 0x1b, 0x07, 0x2d, 0x21, 0x7a, 0x88, 0xab, 0x42, 0x0e,
	0x32, 0x21, 0xf5, 0x20, 0x7a, 0x10, 0xb5, 0x2a, 0x2d, 0x08, 0xc2, 0x14, 0x3c, 0xe8, 0x21, 0x6c,
	0xb7, 0x2f, 0xd9, 0xa5, 0x93, 0x9d, 0x65, 0xe6, 0xcd, 0xc2, 0xde, 0xfd, 0xe1, 0xb2, 0xb3, 0x3b,
	0x76, 0x23, 0x8d, 0xd2, 0x4b, 0x78, 0xdf, 0xbc, 0xef, 0x7d, 0xf9, 0xbe, 0x8f, 0x85, 0xc7, 0x98,
	0x15, 0xaa, 0x24, 0x9d, 0x46, 0x2b, 0x9c, 0x15, 0xf3, 0x99, 0xb2, 0x94, 0x5b, 0xe2, 0xb9, 0x56,
	0xa4, 0xd8, 0xa3, 0x58, 0x5a, 0x8a, 0x13, 0xde, 0xe2, 0xf0, 0x62, 0x3e, 0x7e, 0xb2, 0x52, 0x6a,
	0x25, 0x71, 0xe6, 0x48, 0x97, 0x76, 0x39, 0x33, 0xa4, 0x6d, 0xdc, 0x1c, 0x85, 0xef, 0x01, 0xce,
	0x94, 0xa1, 0x0b, 0x8a, 0xc8, 0x1a, 0x36, 0x82, 0x7e, 0x74, 0x75, 0xa5, 0xd1, 0x98, 0x51, 0x30,
	0x09, 0xa6, 0xfb, 0xc2, 0xc3, 0x6a, 0x93, 0x60, 0x24, 0x29, 0x29, 0x47, 0x9d, 0x49, 0x30, 0xdd,
	0x13, 0x1e, 0x86, 0xd7, 0x30, 0x38, 0x95, 0xd6, 0x10, 0xea, 0x46, 0x84, 0x41, 0x37, 0x8b, 0xd6,
	0xd8, 0x28, 0xb8, 0x99, 0x7d, 0x81, 0x41, 0xa2, 0x0c, 0x2d, 0x8c, 0xa3, 0xa0, 0x19, 0x75, 0x26,
	0x3b, 0xd3, 0x83, 0x93, 0xa7, 0xfc, 0x56, 0xcf, 0xfc, 0xc6, 0x92, 0x38, 0x4c, 0xfe, 0xcc, 0x68,
	0xc2, 0x9f, 0xb0, 0xd7, 0xfc, 0x99, 0x61, 0xdf, 0x60, 0x18, 0xd7, 0xf3, 0x8d, 0x6c, 0xe0, 0x64,
	0x9f, 0x6f, 0x91, 0xdd, 0xf0, 0x29, 0x8e, 0xe2, 0x36, 0x44, 0x13, 0xbe, 0x05, 0x38, 0x55, 0xd9,
	0x32, 0x5d, 0x7d, 0xb2, 0xeb, 0x9c, 0xbd, 0x84, 0x5e, 0x11, 0x49, 0x5b, 0xe7, 0x38, 0x38, 0x39,
	0xe6, 0x75, 0x8f, 0xdc, 0xf7, 0xc8, 0xbf, 0x57, 0x5b, 0x51, 0x93, 0xc2, 0x73, 0xb8, 0xff, 0x35,
	0x35, 0x84, 0xd9, 0x3f, 0x6b, 0x78, 0x06, 0x03, 0xa9, 0xe2, 0x48, 0x2e, 0x7c, 0xcb, 0x1d, 0xb7,
	0x3c, 0x74, 0x8f, 0x1f, 0xea, 0xb7, 0x70, 0x01, 0xfb, 0x5e, 0xca, 0x30, 0x01, 0x0f, 0x64, 0x03,
	0xfe, 0x4e, 0xf9, 0x62, 0x4b, 0xca, 0x4d, 0x1f, 0x62, 0x28, 0x37, 0x30, 0x9a, 0xf0, 0x57, 0x00,
	0x7d, 0x61, 0x33, 0x4a, 0xd7, 0xc8, 0xde, 0x41, 0x1f, 0x33, 0xd2, 0xe9, 0x7f, 0xbb, 0x6b, 0x0e,
	0xf8, 0xe7, 0x8c, 0x74, 0x29, 0xfc, 0xd1, 0xf8, 0x0d, 0xf4, 0xdc, 0x0b, 0x1b, 0xc2, 0xce, 0x35,
	0x96, 0x4d, 0xda, 0x6a, 0x64, 0xc7, 0xbe, 0x40, 0x17, 0xf2, 0xec, 0x5e, 0x53, 0xd5, 0xc7, 0x5d,
	0xe8, 0x52, 0x99, 0x63, 0x55, 0xf7, 0x05, 0xea, 0x02, 0xf5, 0x79, 0xb6, 0x54, 0x77, 0xac, 0x3b,
	0x87, 0x5e, 0x15, 0xc7, 0xb0, 0xd7, 0xd0, 0xab, 0x6a, 0xf1, 0xee, 0xb7, 0x7d, 0x50, 0x8e, 0xec,
	0x7e, 0x45, 0xcd, 0x1f, 0x73, 0xe8, 0x56, 0xf0, 0x16, 0xdf, 0x0f, 0xdb, 0xbe, 0xbb, 0xde, 0xf5,
	0xd1, 0x8f, 0x41, 0x4b, 0xb3, 0x98, 0x5f, 0xee, 0x3a, 0x67, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0x1e, 0x3e, 0x30, 0x94, 0x03, 0x00, 0x00,
}