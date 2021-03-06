// Code generated by protoc-gen-go.
// source: myprobe/myprobe.proto
// DO NOT EDIT!

/*
Package myprobe is a generated protocol buffer package.

It is generated from these files:
	myprobe/myprobe.proto

It has these top-level messages:
	ProbeConf
*/
package myprobe

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_probes "github.com/google/cloudprober/probes/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Redis operation
type ProbeConf_Op int32

const (
	ProbeConf_GET    ProbeConf_Op = 0
	ProbeConf_SET    ProbeConf_Op = 1
	ProbeConf_DELETE ProbeConf_Op = 2
)

var ProbeConf_Op_name = map[int32]string{
	0: "GET",
	1: "SET",
	2: "DELETE",
}
var ProbeConf_Op_value = map[string]int32{
	"GET":    0,
	"SET":    1,
	"DELETE": 2,
}

func (x ProbeConf_Op) Enum() *ProbeConf_Op {
	p := new(ProbeConf_Op)
	*p = x
	return p
}
func (x ProbeConf_Op) String() string {
	return proto.EnumName(ProbeConf_Op_name, int32(x))
}
func (x *ProbeConf_Op) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_Op_value, data, "ProbeConf_Op")
	if err != nil {
		return err
	}
	*x = ProbeConf_Op(value)
	return nil
}
func (ProbeConf_Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ProbeConf struct {
	Op *ProbeConf_Op `protobuf:"varint,1,req,name=op,enum=cloudprober.ProbeConf_Op" json:"op,omitempty"`
	// Key and value for the redis operation
	Key              *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProbeConf) GetOp() ProbeConf_Op {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return ProbeConf_GET
}

func (m *ProbeConf) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ProbeConf) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

var E_RedisProbe = &proto.ExtensionDesc{
	ExtendedType:  (*cloudprober_probes.ProbeDef)(nil),
	ExtensionType: (*ProbeConf)(nil),
	Field:         200,
	Name:          "cloudprober.redis_probe",
	Tag:           "bytes,200,opt,name=redis_probe,json=redisProbe",
	Filename:      "myprobe/myprobe.proto",
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.ProbeConf")
	proto.RegisterEnum("cloudprober.ProbeConf_Op", ProbeConf_Op_name, ProbeConf_Op_value)
	proto.RegisterExtension(E_RedisProbe)
}

func init() { proto.RegisterFile("myprobe/myprobe.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x4f, 0x4a, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0xc9, 0x39,
	0xf9, 0xa5, 0x29, 0x60, 0x91, 0x22, 0x29, 0xf3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xfd, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x24, 0x69, 0x7d, 0x30, 0x55, 0xac,
	0x0f, 0xd6, 0xa8, 0x9f, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x0e, 0x31, 0x45, 0xa9, 0x89, 0x91, 0x8b,
	0x33, 0x00, 0x24, 0xeb, 0x9c, 0x9f, 0x97, 0x26, 0xa4, 0xc9, 0xc5, 0x94, 0x5f, 0x20, 0xc1, 0xa8,
	0xc0, 0xa4, 0xc1, 0x67, 0x24, 0xa9, 0x87, 0x64, 0x82, 0x1e, 0x5c, 0x8d, 0x9e, 0x7f, 0x41, 0x10,
	0x53, 0x7e, 0x81, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0x93, 0x02, 0x93, 0x06, 0x67,
	0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0xe1, 0x28, 0x29, 0x71, 0x31, 0xf9, 0x17, 0x08, 0xb1, 0x73, 0x31, 0xbb, 0xbb,
	0x86, 0x08, 0x30, 0x80, 0x18, 0xc1, 0xae, 0x21, 0x02, 0x8c, 0x42, 0x5c, 0x5c, 0x6c, 0x2e, 0xae,
	0x3e, 0xae, 0x21, 0xae, 0x02, 0x4c, 0x56, 0x61, 0x5c, 0xdc, 0x45, 0xa9, 0x29, 0x99, 0xc5, 0xf1,
	0x60, 0xcb, 0x84, 0x64, 0x50, 0x6c, 0x86, 0xb8, 0x1d, 0xe2, 0x00, 0x97, 0xd4, 0x34, 0x89, 0x13,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0xd8, 0x9d, 0x17, 0xc4, 0x05, 0x36, 0x09, 0xcc, 0x07,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x25, 0x17, 0x20, 0x3a, 0x01, 0x00, 0x00,
}
