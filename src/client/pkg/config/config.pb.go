// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/pkg/config/config.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContextSource int32

const (
	ContextSource_SELF      ContextSource = 0
	ContextSource_CONFIG_V1 ContextSource = 1
)

var ContextSource_name = map[int32]string{
	0: "SELF",
	1: "CONFIG_V1",
}

var ContextSource_value = map[string]int32{
	"SELF":      0,
	"CONFIG_V1": 1,
}

func (x ContextSource) String() string {
	return proto.EnumName(ContextSource_name, int32(x))
}

func (ContextSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{0}
}

// Config specifies the pachyderm config that is read and interpreted by the
// pachctl command-line tool. Right now, this is stored at
// $HOME/.pachyderm/config.
//
// Different versions of the pachyderm config are specified as subfields of this
// message (this allows us to make significant changes to the config structure
// without breaking existing users by defining a new config version).
//
// DO NOT change or remove field numbers from this proto, otherwise ALL
// pachyderm user configs will fail to parse.
type Config struct {
	UserID string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Configuration options. Exactly one of these fields should be set
	// (depending on which version of the config is being used)
	V1                   *ConfigV1 `protobuf:"bytes,2,opt,name=v1,proto3" json:"v1,omitempty"`
	V2                   *ConfigV2 `protobuf:"bytes,3,opt,name=v2,proto3" json:"v2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Config) GetV1() *ConfigV1 {
	if m != nil {
		return m.V1
	}
	return nil
}

func (m *Config) GetV2() *ConfigV2 {
	if m != nil {
		return m.V2
	}
	return nil
}

// ConfigV1 specifies v1 of the pachyderm config (June 30 2017 - June 2019)
// DO NOT change or remove field numbers from this proto, as if you do, v1 user
// configs will become unparseable.
type ConfigV1 struct {
	// A host:port pointing pachd at a pachyderm cluster. Similar to the
	// PACHD_ADDRESS environment variable, though PACHD_ADDRESS overrides
	// this.
	PachdAddress string `protobuf:"bytes,2,opt,name=pachd_address,json=pachdAddress,proto3" json:"pachd_address,omitempty"`
	// Trusted root certificates (overrides installed certificates), formatted
	// as base64-encoded PEM
	ServerCAs string `protobuf:"bytes,3,opt,name=server_cas,json=serverCas,proto3" json:"server_cas,omitempty"`
	// A secret token identifying the current pachctl user within their
	// pachyderm cluster. This is included in all RPCs sent by pachctl, and used
	// to determine if pachctl actions are authorized.
	SessionToken string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	// The currently active transaction for batching together pachctl commands.
	// This can be set or cleared via many of the `pachctl * transaction` commands.
	// This is the ID of the transaction object stored in the pachyderm etcd.
	ActiveTransaction    string   `protobuf:"bytes,4,opt,name=active_transaction,json=activeTransaction,proto3" json:"active_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigV1) Reset()         { *m = ConfigV1{} }
func (m *ConfigV1) String() string { return proto.CompactTextString(m) }
func (*ConfigV1) ProtoMessage()    {}
func (*ConfigV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{1}
}
func (m *ConfigV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigV1.Merge(m, src)
}
func (m *ConfigV1) XXX_Size() int {
	return m.Size()
}
func (m *ConfigV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigV1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigV1 proto.InternalMessageInfo

func (m *ConfigV1) GetPachdAddress() string {
	if m != nil {
		return m.PachdAddress
	}
	return ""
}

func (m *ConfigV1) GetServerCAs() string {
	if m != nil {
		return m.ServerCAs
	}
	return ""
}

func (m *ConfigV1) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *ConfigV1) GetActiveTransaction() string {
	if m != nil {
		return m.ActiveTransaction
	}
	return ""
}

// ConfigV2 specifies v2 of the pachyderm config (June 2019 - present)
type ConfigV2 struct {
	ActiveContext        string              `protobuf:"bytes,1,opt,name=active_context,json=activeContext,proto3" json:"active_context,omitempty"`
	Contexts             map[string]*Context `protobuf:"bytes,2,rep,name=contexts,proto3" json:"contexts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfigV2) Reset()         { *m = ConfigV2{} }
func (m *ConfigV2) String() string { return proto.CompactTextString(m) }
func (*ConfigV2) ProtoMessage()    {}
func (*ConfigV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{2}
}
func (m *ConfigV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigV2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigV2.Merge(m, src)
}
func (m *ConfigV2) XXX_Size() int {
	return m.Size()
}
func (m *ConfigV2) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigV2.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigV2 proto.InternalMessageInfo

func (m *ConfigV2) GetActiveContext() string {
	if m != nil {
		return m.ActiveContext
	}
	return ""
}

func (m *ConfigV2) GetContexts() map[string]*Context {
	if m != nil {
		return m.Contexts
	}
	return nil
}

type Context struct {
	// Where this context came from
	Source ContextSource `protobuf:"varint,1,opt,name=source,proto3,enum=config.ContextSource" json:"source,omitempty"`
	// The hostname or IP address pointing pachd at a pachyderm cluster.
	PachdAddress string `protobuf:"bytes,2,opt,name=pachd_address,json=pachdAddress,proto3" json:"pachd_address,omitempty"`
	// Trusted root certificates (overrides installed certificates), formatted
	// as base64-encoded PEM
	ServerCAs string `protobuf:"bytes,3,opt,name=server_cas,json=serverCas,proto3" json:"server_cas,omitempty"`
	// A secret token identifying the current pachctl user within their
	// pachyderm cluster. This is included in all RPCs sent by pachctl, and used
	// to determine if pachctl actions are authorized.
	SessionToken string `protobuf:"bytes,4,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	// The currently active transaction for batching together pachctl commands.
	// This can be set or cleared via many of the `pachctl * transaction` commands.
	// This is the ID of the transaction object stored in the pachyderm etcd.
	ActiveTransaction    string   `protobuf:"bytes,5,opt,name=active_transaction,json=activeTransaction,proto3" json:"active_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{3}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Context.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return m.Size()
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetSource() ContextSource {
	if m != nil {
		return m.Source
	}
	return ContextSource_SELF
}

func (m *Context) GetPachdAddress() string {
	if m != nil {
		return m.PachdAddress
	}
	return ""
}

func (m *Context) GetServerCAs() string {
	if m != nil {
		return m.ServerCAs
	}
	return ""
}

func (m *Context) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *Context) GetActiveTransaction() string {
	if m != nil {
		return m.ActiveTransaction
	}
	return ""
}

func init() {
	proto.RegisterEnum("config.ContextSource", ContextSource_name, ContextSource_value)
	proto.RegisterType((*Config)(nil), "config.Config")
	proto.RegisterType((*ConfigV1)(nil), "config.ConfigV1")
	proto.RegisterType((*ConfigV2)(nil), "config.ConfigV2")
	proto.RegisterMapType((map[string]*Context)(nil), "config.ConfigV2.ContextsEntry")
	proto.RegisterType((*Context)(nil), "config.Context")
}

func init() { proto.RegisterFile("client/pkg/config/config.proto", fileDescriptor_60f651abce1dcdf3) }

var fileDescriptor_60f651abce1dcdf3 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x26, 0xa9, 0x1b, 0x4f, 0xeb, 0x12, 0x56, 0x20, 0x45, 0x3d, 0xb8, 0x51, 0xaa, 0x4a,
	0x11, 0xa2, 0xb1, 0x62, 0x38, 0xa0, 0xde, 0x1a, 0xd3, 0xa0, 0x48, 0x15, 0x48, 0x4e, 0xe9, 0x81,
	0x8b, 0xe5, 0xae, 0x97, 0xd4, 0x4a, 0xf1, 0x86, 0xdd, 0x8d, 0x45, 0xde, 0x84, 0xd7, 0xe0, 0xcc,
	0x0b, 0x70, 0xe4, 0x09, 0x0a, 0x32, 0x2f, 0x82, 0xf6, 0x27, 0x6d, 0x29, 0x20, 0x71, 0xe2, 0xe4,
	0xf1, 0xf7, 0x7d, 0xf3, 0xcd, 0xcc, 0x8e, 0x06, 0x7c, 0x72, 0x99, 0xd3, 0x42, 0x06, 0xf3, 0xd9,
	0x34, 0x20, 0xac, 0x78, 0x9b, 0xaf, 0x3e, 0xfd, 0x39, 0x67, 0x92, 0x61, 0xc7, 0xfc, 0xed, 0x3c,
	0x98, 0xb2, 0x29, 0xd3, 0x50, 0xa0, 0x22, 0xc3, 0x76, 0xdf, 0x83, 0x13, 0x69, 0x1e, 0xef, 0xc1,
	0xc6, 0x42, 0x50, 0x9e, 0xe4, 0x59, 0x1b, 0x75, 0x50, 0xcf, 0x1d, 0x42, 0x75, 0xb5, 0xeb, 0xbc,
	0x16, 0x94, 0x8f, 0x9f, 0xc7, 0x8e, 0xa2, 0xc6, 0x19, 0xee, 0x40, 0xad, 0x1c, 0xb4, 0x6b, 0x1d,
	0xd4, 0xdb, 0x0c, 0x5b, 0x7d, 0x5b, 0xc7, 0x18, 0x9c, 0x0d, 0xe2, 0x5a, 0x39, 0xd0, 0x8a, 0xb0,
	0x5d, 0xff, 0xa3, 0x22, 0x8c, 0x6b, 0x65, 0xd8, 0xfd, 0x84, 0xa0, 0xb9, 0x4a, 0xc1, 0x7b, 0xe0,
	0xcd, 0x53, 0x72, 0x91, 0x25, 0x69, 0x96, 0x71, 0x2a, 0x84, 0xf6, 0x76, 0xe3, 0x2d, 0x0d, 0x1e,
	0x19, 0x0c, 0x3f, 0x06, 0x10, 0x94, 0x97, 0x94, 0x27, 0x24, 0x15, 0xda, 0xdb, 0x1d, 0x7a, 0xd5,
	0xd5, 0xae, 0x3b, 0xd1, 0x68, 0x74, 0x24, 0x62, 0xd7, 0x08, 0xa2, 0x54, 0x28, 0x4b, 0x41, 0x85,
	0xc8, 0x59, 0x91, 0x48, 0x36, 0xa3, 0x85, 0x19, 0x27, 0xde, 0xb2, 0xe0, 0xa9, 0xc2, 0xf0, 0x01,
	0xe0, 0x94, 0xc8, 0xbc, 0xa4, 0x89, 0xe4, 0x69, 0x21, 0x54, 0xcc, 0x8a, 0x76, 0x43, 0x2b, 0xef,
	0x1b, 0xe6, 0xf4, 0x86, 0xe8, 0x7e, 0xbe, 0xe9, 0x39, 0xc4, 0xfb, 0xb0, 0x6d, 0x73, 0x09, 0x2b,
	0x24, 0xfd, 0x20, 0x6d, 0x05, 0xcf, 0xa0, 0x91, 0x01, 0xf1, 0x21, 0x34, 0x2d, 0xaf, 0xa6, 0xaa,
	0xf7, 0x36, 0x43, 0xff, 0xee, 0x7b, 0xf4, 0xad, 0x56, 0x1c, 0x17, 0x92, 0x2f, 0xe3, 0x6b, 0xfd,
	0xce, 0x09, 0x78, 0xbf, 0x50, 0xb8, 0x05, 0xf5, 0x19, 0x5d, 0xda, 0x42, 0x2a, 0xc4, 0xfb, 0xb0,
	0x5e, 0xa6, 0x97, 0x0b, 0x6a, 0xb7, 0x71, 0xef, 0x96, 0xb7, 0xca, 0x8b, 0x0d, 0x7b, 0x58, 0x7b,
	0x86, 0xba, 0xdf, 0x10, 0x6c, 0xac, 0xba, 0x3a, 0x00, 0x47, 0xb0, 0x05, 0x27, 0x54, 0x7b, 0x6d,
	0x87, 0x0f, 0xef, 0xe4, 0x4d, 0x34, 0x19, 0x5b, 0xd1, 0x7f, 0xd9, 0x4f, 0xe3, 0x9f, 0xf7, 0xb3,
	0xfe, 0x97, 0xfd, 0x3c, 0xea, 0x5d, 0xbf, 0x97, 0xe9, 0x1f, 0x37, 0xa1, 0x31, 0x39, 0x3e, 0x19,
	0xb5, 0xd6, 0xb0, 0x07, 0x6e, 0xf4, 0xea, 0xe5, 0x68, 0xfc, 0x22, 0x39, 0x1b, 0xb4, 0xd0, 0x70,
	0xf4, 0xa5, 0xf2, 0xd1, 0xd7, 0xca, 0x47, 0xdf, 0x2b, 0x1f, 0x7d, 0xfc, 0xe1, 0xaf, 0xbd, 0x79,
	0x3a, 0xcd, 0xe5, 0xc5, 0xe2, 0xbc, 0x4f, 0xd8, 0xbb, 0x40, 0x8d, 0xb5, 0xcc, 0x28, 0xbf, 0x1d,
	0x09, 0x4e, 0x82, 0xdf, 0x6e, 0xec, 0xdc, 0xd1, 0xf7, 0xf3, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x1b, 0xd7, 0x83, 0x7f, 0x03, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if m.V1 != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.V1.Size()))
		n1, err1 := m.V1.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.V2 != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.V2.Size()))
		n2, err2 := m.V2.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigV1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SessionToken) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SessionToken)))
		i += copy(dAtA[i:], m.SessionToken)
	}
	if len(m.PachdAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PachdAddress)))
		i += copy(dAtA[i:], m.PachdAddress)
	}
	if len(m.ServerCAs) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ServerCAs)))
		i += copy(dAtA[i:], m.ServerCAs)
	}
	if len(m.ActiveTransaction) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ActiveTransaction)))
		i += copy(dAtA[i:], m.ActiveTransaction)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigV2) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ActiveContext) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ActiveContext)))
		i += copy(dAtA[i:], m.ActiveContext)
	}
	if len(m.Contexts) > 0 {
		for k, _ := range m.Contexts {
			dAtA[i] = 0x12
			i++
			v := m.Contexts[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovConfig(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + msgSize
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintConfig(dAtA, i, uint64(v.Size()))
				n3, err3 := v.MarshalTo(dAtA[i:])
				if err3 != nil {
					return 0, err3
				}
				i += n3
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Context) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Context) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Source != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Source))
	}
	if len(m.PachdAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PachdAddress)))
		i += copy(dAtA[i:], m.PachdAddress)
	}
	if len(m.ServerCAs) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ServerCAs)))
		i += copy(dAtA[i:], m.ServerCAs)
	}
	if len(m.SessionToken) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SessionToken)))
		i += copy(dAtA[i:], m.SessionToken)
	}
	if len(m.ActiveTransaction) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ActiveTransaction)))
		i += copy(dAtA[i:], m.ActiveTransaction)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.V1 != nil {
		l = m.V1.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.V2 != nil {
		l = m.V2.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PachdAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ServerCAs)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ActiveTransaction)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ActiveContext)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Contexts) > 0 {
		for k, v := range m.Contexts {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovConfig(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Context) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Source != 0 {
		n += 1 + sovConfig(uint64(m.Source))
	}
	l = len(m.PachdAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ServerCAs)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.SessionToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ActiveTransaction)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V1 == nil {
				m.V1 = &ConfigV1{}
			}
			if err := m.V1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V2 == nil {
				m.V2 = &ConfigV2{}
			}
			if err := m.V2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PachdAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PachdAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerCAs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerCAs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveTransaction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveTransaction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contexts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Contexts == nil {
				m.Contexts = make(map[string]*Context)
			}
			var mapkey string
			var mapvalue *Context
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthConfig
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthConfig
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Context{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Contexts[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Context) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Context: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Context: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			m.Source = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Source |= ContextSource(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PachdAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PachdAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerCAs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerCAs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveTransaction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveTransaction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
