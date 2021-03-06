// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/lameduck/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Options struct {
	// How often to check for lame-ducked targets
	ReEvalSec *int32 `protobuf:"varint,1,opt,name=re_eval_sec,json=reEvalSec,def=10" json:"re_eval_sec,omitempty"`
	// Runtime config project. If running on GCE, this defaults to the project
	// containing the VM.
	RuntimeconfigProject *string `protobuf:"bytes,2,opt,name=runtimeconfig_project,json=runtimeconfigProject" json:"runtimeconfig_project,omitempty"`
	// Lame duck targets runtime config name. An operator will create a variable
	// here to mark a target as lame-ducked.
	// TODO(izzycecil): This name needs to be changed.
	RuntimeconfigName *string `protobuf:"bytes,3,opt,name=runtimeconfig_name,json=runtimeconfigName,def=lame-duck-targets" json:"runtimeconfig_name,omitempty"`
	// Lame duck expiration time. We ignore variables (targets) that have been
	// updated more than these many seconds ago. This is a safety mechanism for
	// failing to cleanup. Also, the idea is that if a target has actually
	// disappeared, automatic targets expansion will take care of that some time
	// during this expiration period.
	ExpirationSec        *int32   `protobuf:"varint,4,opt,name=expiration_sec,json=expirationSec,def=300" json:"expiration_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_17e782ab79ded7e1, []int{0}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (dst *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(dst, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

const Default_Options_ReEvalSec int32 = 10
const Default_Options_RuntimeconfigName string = "lame-duck-targets"
const Default_Options_ExpirationSec int32 = 300

func (m *Options) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_Options_ReEvalSec
}

func (m *Options) GetRuntimeconfigProject() string {
	if m != nil && m.RuntimeconfigProject != nil {
		return *m.RuntimeconfigProject
	}
	return ""
}

func (m *Options) GetRuntimeconfigName() string {
	if m != nil && m.RuntimeconfigName != nil {
		return *m.RuntimeconfigName
	}
	return Default_Options_RuntimeconfigName
}

func (m *Options) GetExpirationSec() int32 {
	if m != nil && m.ExpirationSec != nil {
		return *m.ExpirationSec
	}
	return Default_Options_ExpirationSec
}

func init() {
	proto.RegisterType((*Options)(nil), "cloudprober.targets.lameduck.Options")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/lameduck/proto/config.proto", fileDescriptor_config_17e782ab79ded7e1)
}

var fileDescriptor_config_17e782ab79ded7e1 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xce, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x71, 0xb6, 0x55, 0xa4, 0x11, 0x85, 0x06, 0x85, 0x3d, 0x78, 0x28, 0x3d, 0x15, 0xa1,
	0x9b, 0x95, 0xde, 0x7a, 0xd2, 0x83, 0x57, 0x95, 0xed, 0x03, 0x2c, 0xe9, 0x74, 0x8c, 0xd1, 0x64,
	0x27, 0xcc, 0x66, 0x8b, 0x8f, 0xe9, 0x23, 0x49, 0xb6, 0x15, 0xbb, 0xc7, 0xf0, 0xe5, 0xcf, 0xfc,
	0xc4, 0x93, 0xb1, 0xf1, 0xa3, 0xdb, 0x16, 0x40, 0x5e, 0x19, 0x22, 0xe3, 0x50, 0x81, 0xa3, 0x6e,
	0x17, 0x98, 0xb6, 0xc8, 0x2a, 0x6a, 0x36, 0x18, 0x5b, 0xe5, 0xb4, 0xc7, 0x5d, 0x07, 0x5f, 0x2a,
	0x30, 0x45, 0x52, 0x40, 0xcd, 0xbb, 0x35, 0x45, 0xff, 0x90, 0x77, 0x27, 0x41, 0x71, 0x0c, 0x8a,
	0xbf, 0x60, 0xfe, 0x93, 0x89, 0x8b, 0xd7, 0x10, 0x2d, 0x35, 0xad, 0x9c, 0x8b, 0x4b, 0xc6, 0x1a,
	0xf7, 0xda, 0xd5, 0x2d, 0x42, 0x9e, 0xcd, 0xb2, 0xc5, 0xf9, 0x7a, 0xf4, 0x50, 0x56, 0x13, 0xc6,
	0xe7, 0xbd, 0x76, 0x1b, 0x04, 0xb9, 0x12, 0xb7, 0xdc, 0x35, 0xd1, 0x7a, 0x3c, 0x1c, 0xa9, 0x03,
	0xd3, 0x27, 0x42, 0xcc, 0x47, 0xb3, 0x6c, 0x31, 0xa9, 0x6e, 0x06, 0xe3, 0xdb, 0x61, 0x93, 0x8f,
	0x42, 0x0e, 0xa3, 0x46, 0x7b, 0xcc, 0xc7, 0xa9, 0x58, 0x4f, 0x13, 0x65, 0x99, 0x2c, 0xcb, 0x23,
	0xae, 0x9a, 0x0e, 0x3e, 0xbf, 0x68, 0x8f, 0xf2, 0x5e, 0x5c, 0xe3, 0x77, 0xb0, 0xac, 0x93, 0xb4,
	0xd7, 0x9d, 0xf5, 0xba, 0xf1, 0xaa, 0x2c, 0xab, 0xab, 0xff, 0x69, 0x83, 0xf0, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x81, 0x30, 0x11, 0x9a, 0x34, 0x01, 0x00, 0x00,
}
